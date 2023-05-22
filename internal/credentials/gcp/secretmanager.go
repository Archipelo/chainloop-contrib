package gcp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"google.golang.org/api/option"

	"github.com/chainloop-dev/chainloop/internal/credentials"
	"github.com/chainloop-dev/chainloop/internal/servicelogger"
	"github.com/docker/distribution/uuid"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/googleapis/gax-go/v2"
)

type SecretsManagerInterface interface {
	CreateSecret(ctx context.Context, req *secretmanagerpb.CreateSecretRequest, opts ...gax.CallOption) (*secretmanagerpb.Secret, error)
	AddSecretVersion(ctx context.Context, req *secretmanagerpb.AddSecretVersionRequest, opts ...gax.CallOption) (*secretmanagerpb.SecretVersion, error)
	AccessSecretVersion(ctx context.Context, req *secretmanagerpb.AccessSecretVersionRequest, opts ...gax.CallOption) (*secretmanagerpb.AccessSecretVersionResponse, error)
	DeleteSecret(ctx context.Context, req *secretmanagerpb.DeleteSecretRequest, opts ...gax.CallOption) error
	GetSecret(ctx context.Context, req *secretmanagerpb.GetSecretRequest, opts ...gax.CallOption) (*secretmanagerpb.Secret, error)
	Close() error
}

type Manager struct {
	projectID string
	client    SecretsManagerInterface
	logger    *log.Helper
}

type NewManagerOpts struct {
	ProjectID, AuthKey string
	Logger             log.Logger
}

func NewManager(opts *NewManagerOpts) (*Manager, error) {
	if opts.ProjectID == "" || opts.AuthKey == "" {
		return nil, errors.New("projectID and authKey are required")
	}

	l := opts.Logger
	if l == nil {
		l = log.NewStdLogger(io.Discard)
	}

	logger := servicelogger.ScopedHelper(l, "credentials/gcp-secrets-manager")
	logger.Infow("msg", "configuring gcp secrets-manager", "projectID", opts.ProjectID)

	cli, err := secretmanager.NewRESTClient(context.TODO(), option.WithCredentialsJSON([]byte(opts.AuthKey)))
	if err != nil {
		return nil, fmt.Errorf("error while connecting to the project: %v", err)
	}

	return &Manager{
		projectID: opts.ProjectID,
		client:    cli,
		logger:    logger,
	}, nil
}

// SaveCredentials saves credentials
func (m *Manager) SaveCredentials(ctx context.Context, orgID string, creds any) (string, error) {
	secretID := strings.Join([]string{orgID, uuid.Generate().String()}, "-")

	// first create the secret itself
	createSecretReq := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", m.projectID),
		SecretId: secretID,
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
		},
	}
	secret, err := m.client.CreateSecret(ctx, createSecretReq)
	if err != nil {
		return "", fmt.Errorf("creating secret in GCP: %w", err)
	}

	// store creds in key-value pair
	c, err := json.Marshal(creds)
	if err != nil {
		return "", fmt.Errorf("marshaling credentials to be stored: %w", err)
	}

	// once the secret is created store it as the newest version
	addSecretVersionReq := &secretmanagerpb.AddSecretVersionRequest{
		Parent: secret.Name,
		Payload: &secretmanagerpb.SecretPayload{
			Data: c,
		},
	}
	_, err = m.client.AddSecretVersion(ctx, addSecretVersionReq)
	if err != nil {
		return "", fmt.Errorf("creating secret version in GCP: %w", err)
	}

	return secretID, nil
}

// ReadCredentials reads the latest version of the credentials
func (m *Manager) ReadCredentials(ctx context.Context, secretID string, creds any) error {
	getSecretRequest := secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v/versions/latest", m.projectID, secretID),
	}

	result, err := m.client.AccessSecretVersion(ctx, &getSecretRequest)
	if err != nil {
		return fmt.Errorf("%w: path=%s", credentials.ErrNotFound, secretID)
	}

	return json.Unmarshal(result.Payload.Data, creds)
}

// DeleteCredentials deletes credentials and versions
func (m *Manager) DeleteCredentials(ctx context.Context, secretID string) error {
	if m.secretExists(ctx, secretID) {
		deleteRequest := secretmanagerpb.DeleteSecretRequest{
			Name: fmt.Sprintf("projects/%v/secrets/%v", m.projectID, secretID),
		}

		return m.client.DeleteSecret(ctx, &deleteRequest)
	}

	return fmt.Errorf("%w: path=%s", credentials.ErrNotFound, secretID)
}

func (c *Manager) secretExists(ctx context.Context, secretID string) bool {
	accessRequest := secretmanagerpb.GetSecretRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v", c.projectID, secretID),
	}

	_, err := c.client.GetSecret(ctx, &accessRequest)
	return err == nil
}

// Close closes the manager
func (m *Manager) Close() error {
	if m.client == nil {
		return nil
	}
	return m.client.Close()
}