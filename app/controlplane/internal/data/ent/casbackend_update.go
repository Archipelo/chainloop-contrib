// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/biz"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/casbackend"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/organization"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/predicate"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/workflowrun"
	"github.com/google/uuid"
)

// CASBackendUpdate is the builder for updating CASBackend entities.
type CASBackendUpdate struct {
	config
	hooks    []Hook
	mutation *CASBackendMutation
}

// Where appends a list predicates to the CASBackendUpdate builder.
func (cbu *CASBackendUpdate) Where(ps ...predicate.CASBackend) *CASBackendUpdate {
	cbu.mutation.Where(ps...)
	return cbu
}

// SetDescription sets the "description" field.
func (cbu *CASBackendUpdate) SetDescription(s string) *CASBackendUpdate {
	cbu.mutation.SetDescription(s)
	return cbu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cbu *CASBackendUpdate) SetNillableDescription(s *string) *CASBackendUpdate {
	if s != nil {
		cbu.SetDescription(*s)
	}
	return cbu
}

// ClearDescription clears the value of the "description" field.
func (cbu *CASBackendUpdate) ClearDescription() *CASBackendUpdate {
	cbu.mutation.ClearDescription()
	return cbu
}

// SetSecretName sets the "secret_name" field.
func (cbu *CASBackendUpdate) SetSecretName(s string) *CASBackendUpdate {
	cbu.mutation.SetSecretName(s)
	return cbu
}

// SetValidationStatus sets the "validation_status" field.
func (cbu *CASBackendUpdate) SetValidationStatus(bbvs biz.CASBackendValidationStatus) *CASBackendUpdate {
	cbu.mutation.SetValidationStatus(bbvs)
	return cbu
}

// SetNillableValidationStatus sets the "validation_status" field if the given value is not nil.
func (cbu *CASBackendUpdate) SetNillableValidationStatus(bbvs *biz.CASBackendValidationStatus) *CASBackendUpdate {
	if bbvs != nil {
		cbu.SetValidationStatus(*bbvs)
	}
	return cbu
}

// SetValidatedAt sets the "validated_at" field.
func (cbu *CASBackendUpdate) SetValidatedAt(t time.Time) *CASBackendUpdate {
	cbu.mutation.SetValidatedAt(t)
	return cbu
}

// SetNillableValidatedAt sets the "validated_at" field if the given value is not nil.
func (cbu *CASBackendUpdate) SetNillableValidatedAt(t *time.Time) *CASBackendUpdate {
	if t != nil {
		cbu.SetValidatedAt(*t)
	}
	return cbu
}

// SetDefault sets the "default" field.
func (cbu *CASBackendUpdate) SetDefault(b bool) *CASBackendUpdate {
	cbu.mutation.SetDefault(b)
	return cbu
}

// SetNillableDefault sets the "default" field if the given value is not nil.
func (cbu *CASBackendUpdate) SetNillableDefault(b *bool) *CASBackendUpdate {
	if b != nil {
		cbu.SetDefault(*b)
	}
	return cbu
}

// SetDeletedAt sets the "deleted_at" field.
func (cbu *CASBackendUpdate) SetDeletedAt(t time.Time) *CASBackendUpdate {
	cbu.mutation.SetDeletedAt(t)
	return cbu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cbu *CASBackendUpdate) SetNillableDeletedAt(t *time.Time) *CASBackendUpdate {
	if t != nil {
		cbu.SetDeletedAt(*t)
	}
	return cbu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cbu *CASBackendUpdate) ClearDeletedAt() *CASBackendUpdate {
	cbu.mutation.ClearDeletedAt()
	return cbu
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (cbu *CASBackendUpdate) SetOrganizationID(id uuid.UUID) *CASBackendUpdate {
	cbu.mutation.SetOrganizationID(id)
	return cbu
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (cbu *CASBackendUpdate) SetOrganization(o *Organization) *CASBackendUpdate {
	return cbu.SetOrganizationID(o.ID)
}

// AddWorkflowRunIDs adds the "workflow_run" edge to the WorkflowRun entity by IDs.
func (cbu *CASBackendUpdate) AddWorkflowRunIDs(ids ...uuid.UUID) *CASBackendUpdate {
	cbu.mutation.AddWorkflowRunIDs(ids...)
	return cbu
}

// AddWorkflowRun adds the "workflow_run" edges to the WorkflowRun entity.
func (cbu *CASBackendUpdate) AddWorkflowRun(w ...*WorkflowRun) *CASBackendUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return cbu.AddWorkflowRunIDs(ids...)
}

// Mutation returns the CASBackendMutation object of the builder.
func (cbu *CASBackendUpdate) Mutation() *CASBackendMutation {
	return cbu.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (cbu *CASBackendUpdate) ClearOrganization() *CASBackendUpdate {
	cbu.mutation.ClearOrganization()
	return cbu
}

// ClearWorkflowRun clears all "workflow_run" edges to the WorkflowRun entity.
func (cbu *CASBackendUpdate) ClearWorkflowRun() *CASBackendUpdate {
	cbu.mutation.ClearWorkflowRun()
	return cbu
}

// RemoveWorkflowRunIDs removes the "workflow_run" edge to WorkflowRun entities by IDs.
func (cbu *CASBackendUpdate) RemoveWorkflowRunIDs(ids ...uuid.UUID) *CASBackendUpdate {
	cbu.mutation.RemoveWorkflowRunIDs(ids...)
	return cbu
}

// RemoveWorkflowRun removes "workflow_run" edges to WorkflowRun entities.
func (cbu *CASBackendUpdate) RemoveWorkflowRun(w ...*WorkflowRun) *CASBackendUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return cbu.RemoveWorkflowRunIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cbu *CASBackendUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cbu.sqlSave, cbu.mutation, cbu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cbu *CASBackendUpdate) SaveX(ctx context.Context) int {
	affected, err := cbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cbu *CASBackendUpdate) Exec(ctx context.Context) error {
	_, err := cbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbu *CASBackendUpdate) ExecX(ctx context.Context) {
	if err := cbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cbu *CASBackendUpdate) check() error {
	if v, ok := cbu.mutation.ValidationStatus(); ok {
		if err := casbackend.ValidationStatusValidator(v); err != nil {
			return &ValidationError{Name: "validation_status", err: fmt.Errorf(`ent: validator failed for field "CASBackend.validation_status": %w`, err)}
		}
	}
	if _, ok := cbu.mutation.OrganizationID(); cbu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CASBackend.organization"`)
	}
	return nil
}

func (cbu *CASBackendUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cbu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(casbackend.Table, casbackend.Columns, sqlgraph.NewFieldSpec(casbackend.FieldID, field.TypeUUID))
	if ps := cbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cbu.mutation.Description(); ok {
		_spec.SetField(casbackend.FieldDescription, field.TypeString, value)
	}
	if cbu.mutation.DescriptionCleared() {
		_spec.ClearField(casbackend.FieldDescription, field.TypeString)
	}
	if value, ok := cbu.mutation.SecretName(); ok {
		_spec.SetField(casbackend.FieldSecretName, field.TypeString, value)
	}
	if value, ok := cbu.mutation.ValidationStatus(); ok {
		_spec.SetField(casbackend.FieldValidationStatus, field.TypeEnum, value)
	}
	if value, ok := cbu.mutation.ValidatedAt(); ok {
		_spec.SetField(casbackend.FieldValidatedAt, field.TypeTime, value)
	}
	if value, ok := cbu.mutation.Default(); ok {
		_spec.SetField(casbackend.FieldDefault, field.TypeBool, value)
	}
	if value, ok := cbu.mutation.DeletedAt(); ok {
		_spec.SetField(casbackend.FieldDeletedAt, field.TypeTime, value)
	}
	if cbu.mutation.DeletedAtCleared() {
		_spec.ClearField(casbackend.FieldDeletedAt, field.TypeTime)
	}
	if cbu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   casbackend.OrganizationTable,
			Columns: []string{casbackend.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cbu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   casbackend.OrganizationTable,
			Columns: []string{casbackend.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cbu.mutation.WorkflowRunCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   casbackend.WorkflowRunTable,
			Columns: casbackend.WorkflowRunPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowrun.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cbu.mutation.RemovedWorkflowRunIDs(); len(nodes) > 0 && !cbu.mutation.WorkflowRunCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   casbackend.WorkflowRunTable,
			Columns: casbackend.WorkflowRunPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowrun.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cbu.mutation.WorkflowRunIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   casbackend.WorkflowRunTable,
			Columns: casbackend.WorkflowRunPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowrun.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{casbackend.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cbu.mutation.done = true
	return n, nil
}

// CASBackendUpdateOne is the builder for updating a single CASBackend entity.
type CASBackendUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CASBackendMutation
}

// SetDescription sets the "description" field.
func (cbuo *CASBackendUpdateOne) SetDescription(s string) *CASBackendUpdateOne {
	cbuo.mutation.SetDescription(s)
	return cbuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cbuo *CASBackendUpdateOne) SetNillableDescription(s *string) *CASBackendUpdateOne {
	if s != nil {
		cbuo.SetDescription(*s)
	}
	return cbuo
}

// ClearDescription clears the value of the "description" field.
func (cbuo *CASBackendUpdateOne) ClearDescription() *CASBackendUpdateOne {
	cbuo.mutation.ClearDescription()
	return cbuo
}

// SetSecretName sets the "secret_name" field.
func (cbuo *CASBackendUpdateOne) SetSecretName(s string) *CASBackendUpdateOne {
	cbuo.mutation.SetSecretName(s)
	return cbuo
}

// SetValidationStatus sets the "validation_status" field.
func (cbuo *CASBackendUpdateOne) SetValidationStatus(bbvs biz.CASBackendValidationStatus) *CASBackendUpdateOne {
	cbuo.mutation.SetValidationStatus(bbvs)
	return cbuo
}

// SetNillableValidationStatus sets the "validation_status" field if the given value is not nil.
func (cbuo *CASBackendUpdateOne) SetNillableValidationStatus(bbvs *biz.CASBackendValidationStatus) *CASBackendUpdateOne {
	if bbvs != nil {
		cbuo.SetValidationStatus(*bbvs)
	}
	return cbuo
}

// SetValidatedAt sets the "validated_at" field.
func (cbuo *CASBackendUpdateOne) SetValidatedAt(t time.Time) *CASBackendUpdateOne {
	cbuo.mutation.SetValidatedAt(t)
	return cbuo
}

// SetNillableValidatedAt sets the "validated_at" field if the given value is not nil.
func (cbuo *CASBackendUpdateOne) SetNillableValidatedAt(t *time.Time) *CASBackendUpdateOne {
	if t != nil {
		cbuo.SetValidatedAt(*t)
	}
	return cbuo
}

// SetDefault sets the "default" field.
func (cbuo *CASBackendUpdateOne) SetDefault(b bool) *CASBackendUpdateOne {
	cbuo.mutation.SetDefault(b)
	return cbuo
}

// SetNillableDefault sets the "default" field if the given value is not nil.
func (cbuo *CASBackendUpdateOne) SetNillableDefault(b *bool) *CASBackendUpdateOne {
	if b != nil {
		cbuo.SetDefault(*b)
	}
	return cbuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cbuo *CASBackendUpdateOne) SetDeletedAt(t time.Time) *CASBackendUpdateOne {
	cbuo.mutation.SetDeletedAt(t)
	return cbuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cbuo *CASBackendUpdateOne) SetNillableDeletedAt(t *time.Time) *CASBackendUpdateOne {
	if t != nil {
		cbuo.SetDeletedAt(*t)
	}
	return cbuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cbuo *CASBackendUpdateOne) ClearDeletedAt() *CASBackendUpdateOne {
	cbuo.mutation.ClearDeletedAt()
	return cbuo
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (cbuo *CASBackendUpdateOne) SetOrganizationID(id uuid.UUID) *CASBackendUpdateOne {
	cbuo.mutation.SetOrganizationID(id)
	return cbuo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (cbuo *CASBackendUpdateOne) SetOrganization(o *Organization) *CASBackendUpdateOne {
	return cbuo.SetOrganizationID(o.ID)
}

// AddWorkflowRunIDs adds the "workflow_run" edge to the WorkflowRun entity by IDs.
func (cbuo *CASBackendUpdateOne) AddWorkflowRunIDs(ids ...uuid.UUID) *CASBackendUpdateOne {
	cbuo.mutation.AddWorkflowRunIDs(ids...)
	return cbuo
}

// AddWorkflowRun adds the "workflow_run" edges to the WorkflowRun entity.
func (cbuo *CASBackendUpdateOne) AddWorkflowRun(w ...*WorkflowRun) *CASBackendUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return cbuo.AddWorkflowRunIDs(ids...)
}

// Mutation returns the CASBackendMutation object of the builder.
func (cbuo *CASBackendUpdateOne) Mutation() *CASBackendMutation {
	return cbuo.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (cbuo *CASBackendUpdateOne) ClearOrganization() *CASBackendUpdateOne {
	cbuo.mutation.ClearOrganization()
	return cbuo
}

// ClearWorkflowRun clears all "workflow_run" edges to the WorkflowRun entity.
func (cbuo *CASBackendUpdateOne) ClearWorkflowRun() *CASBackendUpdateOne {
	cbuo.mutation.ClearWorkflowRun()
	return cbuo
}

// RemoveWorkflowRunIDs removes the "workflow_run" edge to WorkflowRun entities by IDs.
func (cbuo *CASBackendUpdateOne) RemoveWorkflowRunIDs(ids ...uuid.UUID) *CASBackendUpdateOne {
	cbuo.mutation.RemoveWorkflowRunIDs(ids...)
	return cbuo
}

// RemoveWorkflowRun removes "workflow_run" edges to WorkflowRun entities.
func (cbuo *CASBackendUpdateOne) RemoveWorkflowRun(w ...*WorkflowRun) *CASBackendUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return cbuo.RemoveWorkflowRunIDs(ids...)
}

// Where appends a list predicates to the CASBackendUpdate builder.
func (cbuo *CASBackendUpdateOne) Where(ps ...predicate.CASBackend) *CASBackendUpdateOne {
	cbuo.mutation.Where(ps...)
	return cbuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cbuo *CASBackendUpdateOne) Select(field string, fields ...string) *CASBackendUpdateOne {
	cbuo.fields = append([]string{field}, fields...)
	return cbuo
}

// Save executes the query and returns the updated CASBackend entity.
func (cbuo *CASBackendUpdateOne) Save(ctx context.Context) (*CASBackend, error) {
	return withHooks(ctx, cbuo.sqlSave, cbuo.mutation, cbuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cbuo *CASBackendUpdateOne) SaveX(ctx context.Context) *CASBackend {
	node, err := cbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cbuo *CASBackendUpdateOne) Exec(ctx context.Context) error {
	_, err := cbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbuo *CASBackendUpdateOne) ExecX(ctx context.Context) {
	if err := cbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cbuo *CASBackendUpdateOne) check() error {
	if v, ok := cbuo.mutation.ValidationStatus(); ok {
		if err := casbackend.ValidationStatusValidator(v); err != nil {
			return &ValidationError{Name: "validation_status", err: fmt.Errorf(`ent: validator failed for field "CASBackend.validation_status": %w`, err)}
		}
	}
	if _, ok := cbuo.mutation.OrganizationID(); cbuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CASBackend.organization"`)
	}
	return nil
}

func (cbuo *CASBackendUpdateOne) sqlSave(ctx context.Context) (_node *CASBackend, err error) {
	if err := cbuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(casbackend.Table, casbackend.Columns, sqlgraph.NewFieldSpec(casbackend.FieldID, field.TypeUUID))
	id, ok := cbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CASBackend.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, casbackend.FieldID)
		for _, f := range fields {
			if !casbackend.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != casbackend.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cbuo.mutation.Description(); ok {
		_spec.SetField(casbackend.FieldDescription, field.TypeString, value)
	}
	if cbuo.mutation.DescriptionCleared() {
		_spec.ClearField(casbackend.FieldDescription, field.TypeString)
	}
	if value, ok := cbuo.mutation.SecretName(); ok {
		_spec.SetField(casbackend.FieldSecretName, field.TypeString, value)
	}
	if value, ok := cbuo.mutation.ValidationStatus(); ok {
		_spec.SetField(casbackend.FieldValidationStatus, field.TypeEnum, value)
	}
	if value, ok := cbuo.mutation.ValidatedAt(); ok {
		_spec.SetField(casbackend.FieldValidatedAt, field.TypeTime, value)
	}
	if value, ok := cbuo.mutation.Default(); ok {
		_spec.SetField(casbackend.FieldDefault, field.TypeBool, value)
	}
	if value, ok := cbuo.mutation.DeletedAt(); ok {
		_spec.SetField(casbackend.FieldDeletedAt, field.TypeTime, value)
	}
	if cbuo.mutation.DeletedAtCleared() {
		_spec.ClearField(casbackend.FieldDeletedAt, field.TypeTime)
	}
	if cbuo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   casbackend.OrganizationTable,
			Columns: []string{casbackend.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cbuo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   casbackend.OrganizationTable,
			Columns: []string{casbackend.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cbuo.mutation.WorkflowRunCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   casbackend.WorkflowRunTable,
			Columns: casbackend.WorkflowRunPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowrun.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cbuo.mutation.RemovedWorkflowRunIDs(); len(nodes) > 0 && !cbuo.mutation.WorkflowRunCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   casbackend.WorkflowRunTable,
			Columns: casbackend.WorkflowRunPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowrun.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cbuo.mutation.WorkflowRunIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   casbackend.WorkflowRunTable,
			Columns: casbackend.WorkflowRunPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowrun.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CASBackend{config: cbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{casbackend.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cbuo.mutation.done = true
	return _node, nil
}
