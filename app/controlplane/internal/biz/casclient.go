//
// Copyright 2023 The Chainloop Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package biz

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/chainloop-dev/chainloop/app/controlplane/internal/conf"
	"github.com/chainloop-dev/chainloop/internal/casclient"
	"github.com/chainloop-dev/chainloop/internal/grpcconn"
	casJWT "github.com/chainloop-dev/chainloop/internal/robotaccount/cas"
	"github.com/chainloop-dev/chainloop/internal/servicelogger"
	"github.com/go-kratos/kratos/v2/log"
)

type CASClientUseCase struct {
	// to generate temporary credentials
	credsProvider *CASCredentialsUseCase
	// configuration to generate the client
	casServerConf *conf.Bootstrap_CASServer
	// factory to generate the client
	casClientFactory CASClientFactory
	logger           *log.Helper
}

type CASUploader interface {
	Upload(ctx context.Context, secretID string, content io.Reader, filename, digest string) error
}

type CASDownloader interface {
	Download(ctx context.Context, secretID string, w io.Writer, digest string) error
}

type CASClient interface {
	CASUploader
	CASDownloader
}

type CASClientFactory func(conf *conf.Bootstrap_CASServer, token string) (casclient.DownloaderUploader, error)
type CASClientOpts func(u *CASClientUseCase)

func WithClientFactory(f CASClientFactory) CASClientOpts {
	return func(c *CASClientUseCase) {
		c.casClientFactory = f
	}
}

func NewCASClientUseCase(credsProvider *CASCredentialsUseCase, config *conf.Bootstrap_CASServer, l log.Logger, opts ...CASClientOpts) *CASClientUseCase {
	// generate a client from the given configuration
	defaultCasClientFactory := func(conf *conf.Bootstrap_CASServer, token string) (casclient.DownloaderUploader, error) {
		conn, err := grpcconn.New(conf.GetGrpc().GetAddr(), token, conf.GetInsecure())
		if err != nil {
			return nil, fmt.Errorf("failed to create grpc connection: %w", err)
		}

		return casclient.New(conn), nil
	}

	uc := &CASClientUseCase{
		credsProvider:    credsProvider,
		casServerConf:    config,
		logger:           servicelogger.ScopedHelper(l, "biz/cas-client"),
		casClientFactory: defaultCasClientFactory,
	}

	for _, opt := range opts {
		opt(uc)
	}

	return uc
}

// The secretID is embedded in the JWT token and is used to identify the secret by the CAS server
func (uc *CASClientUseCase) Upload(ctx context.Context, secretID string, content io.Reader, filename, digest string) error {
	uc.logger.Infow("msg", "upload initialized", "filename", filename, "digest", digest)

	// client with temporary set of credentials
	client, err := uc.casAPIClient(secretID, casJWT.Uploader)
	if err != nil {
		return fmt.Errorf("failed to create cas client: %w", err)
	}

	status, err := client.Upload(ctx, content, filename, digest)
	if err != nil {
		return fmt.Errorf("failed to upload content: %w", err)
	}

	uc.logger.Infow("msg", "upload finished", "status", status)

	return nil
}

func (uc *CASClientUseCase) Download(ctx context.Context, secretID string, w io.Writer, digest string) error {
	uc.logger.Infow("msg", "download initialized", "digest", digest)

	client, err := uc.casAPIClient(secretID, casJWT.Downloader)
	if err != nil {
		return fmt.Errorf("failed to create cas client: %w", err)
	}

	if err := client.Download(ctx, w, digest); err != nil {
		return fmt.Errorf("failed to download content: %w", err)
	}

	uc.logger.Infow("msg", "download finalized", "digest", digest)

	return nil
}

// create a client with a temporary set of credentials for a specific operation
func (uc *CASClientUseCase) casAPIClient(secretID string, role casJWT.Role) (casclient.DownloaderUploader, error) {
	token, err := uc.credsProvider.GenerateTemporaryCredentials(secretID, role)
	if err != nil {
		return nil, fmt.Errorf("failed to generate temporary credentials: %w", err)
	}

	// Initialize connection to CAS server
	return uc.casClientFactory(uc.casServerConf, token)
}

// If the CAS server can be reached and reports readiness
func (uc *CASClientUseCase) IsReady(ctx context.Context) (bool, error) {
	if uc.casServerConf == nil {
		return false, errors.New("missing CAS server configuration")
	}

	err := uc.casServerConf.ValidateAll()
	if err != nil {
		return false, fmt.Errorf("invalid CAS client configuration: %w", err)
	}

	c, err := uc.casClientFactory(uc.casServerConf, "")
	if err != nil {
		return false, fmt.Errorf("failed to create CAS client: %w", err)
	}

	return c.IsReady(ctx)
}