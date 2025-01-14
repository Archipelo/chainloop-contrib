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
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/integration"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/integrationattachment"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/predicate"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/workflow"
	"github.com/google/uuid"
)

// IntegrationAttachmentUpdate is the builder for updating IntegrationAttachment entities.
type IntegrationAttachmentUpdate struct {
	config
	hooks    []Hook
	mutation *IntegrationAttachmentMutation
}

// Where appends a list predicates to the IntegrationAttachmentUpdate builder.
func (iau *IntegrationAttachmentUpdate) Where(ps ...predicate.IntegrationAttachment) *IntegrationAttachmentUpdate {
	iau.mutation.Where(ps...)
	return iau
}

// SetConfiguration sets the "configuration" field.
func (iau *IntegrationAttachmentUpdate) SetConfiguration(b []byte) *IntegrationAttachmentUpdate {
	iau.mutation.SetConfiguration(b)
	return iau
}

// ClearConfiguration clears the value of the "configuration" field.
func (iau *IntegrationAttachmentUpdate) ClearConfiguration() *IntegrationAttachmentUpdate {
	iau.mutation.ClearConfiguration()
	return iau
}

// SetDeletedAt sets the "deleted_at" field.
func (iau *IntegrationAttachmentUpdate) SetDeletedAt(t time.Time) *IntegrationAttachmentUpdate {
	iau.mutation.SetDeletedAt(t)
	return iau
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (iau *IntegrationAttachmentUpdate) SetNillableDeletedAt(t *time.Time) *IntegrationAttachmentUpdate {
	if t != nil {
		iau.SetDeletedAt(*t)
	}
	return iau
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (iau *IntegrationAttachmentUpdate) ClearDeletedAt() *IntegrationAttachmentUpdate {
	iau.mutation.ClearDeletedAt()
	return iau
}

// SetIntegrationID sets the "integration" edge to the Integration entity by ID.
func (iau *IntegrationAttachmentUpdate) SetIntegrationID(id uuid.UUID) *IntegrationAttachmentUpdate {
	iau.mutation.SetIntegrationID(id)
	return iau
}

// SetIntegration sets the "integration" edge to the Integration entity.
func (iau *IntegrationAttachmentUpdate) SetIntegration(i *Integration) *IntegrationAttachmentUpdate {
	return iau.SetIntegrationID(i.ID)
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (iau *IntegrationAttachmentUpdate) SetWorkflowID(id uuid.UUID) *IntegrationAttachmentUpdate {
	iau.mutation.SetWorkflowID(id)
	return iau
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (iau *IntegrationAttachmentUpdate) SetWorkflow(w *Workflow) *IntegrationAttachmentUpdate {
	return iau.SetWorkflowID(w.ID)
}

// Mutation returns the IntegrationAttachmentMutation object of the builder.
func (iau *IntegrationAttachmentUpdate) Mutation() *IntegrationAttachmentMutation {
	return iau.mutation
}

// ClearIntegration clears the "integration" edge to the Integration entity.
func (iau *IntegrationAttachmentUpdate) ClearIntegration() *IntegrationAttachmentUpdate {
	iau.mutation.ClearIntegration()
	return iau
}

// ClearWorkflow clears the "workflow" edge to the Workflow entity.
func (iau *IntegrationAttachmentUpdate) ClearWorkflow() *IntegrationAttachmentUpdate {
	iau.mutation.ClearWorkflow()
	return iau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iau *IntegrationAttachmentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, iau.sqlSave, iau.mutation, iau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iau *IntegrationAttachmentUpdate) SaveX(ctx context.Context) int {
	affected, err := iau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iau *IntegrationAttachmentUpdate) Exec(ctx context.Context) error {
	_, err := iau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iau *IntegrationAttachmentUpdate) ExecX(ctx context.Context) {
	if err := iau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iau *IntegrationAttachmentUpdate) check() error {
	if _, ok := iau.mutation.IntegrationID(); iau.mutation.IntegrationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "IntegrationAttachment.integration"`)
	}
	if _, ok := iau.mutation.WorkflowID(); iau.mutation.WorkflowCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "IntegrationAttachment.workflow"`)
	}
	return nil
}

func (iau *IntegrationAttachmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(integrationattachment.Table, integrationattachment.Columns, sqlgraph.NewFieldSpec(integrationattachment.FieldID, field.TypeUUID))
	if ps := iau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iau.mutation.Configuration(); ok {
		_spec.SetField(integrationattachment.FieldConfiguration, field.TypeBytes, value)
	}
	if iau.mutation.ConfigurationCleared() {
		_spec.ClearField(integrationattachment.FieldConfiguration, field.TypeBytes)
	}
	if value, ok := iau.mutation.DeletedAt(); ok {
		_spec.SetField(integrationattachment.FieldDeletedAt, field.TypeTime, value)
	}
	if iau.mutation.DeletedAtCleared() {
		_spec.ClearField(integrationattachment.FieldDeletedAt, field.TypeTime)
	}
	if iau.mutation.IntegrationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   integrationattachment.IntegrationTable,
			Columns: []string{integrationattachment.IntegrationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integration.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iau.mutation.IntegrationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   integrationattachment.IntegrationTable,
			Columns: []string{integrationattachment.IntegrationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integration.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iau.mutation.WorkflowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   integrationattachment.WorkflowTable,
			Columns: []string{integrationattachment.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iau.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   integrationattachment.WorkflowTable,
			Columns: []string{integrationattachment.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{integrationattachment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iau.mutation.done = true
	return n, nil
}

// IntegrationAttachmentUpdateOne is the builder for updating a single IntegrationAttachment entity.
type IntegrationAttachmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IntegrationAttachmentMutation
}

// SetConfiguration sets the "configuration" field.
func (iauo *IntegrationAttachmentUpdateOne) SetConfiguration(b []byte) *IntegrationAttachmentUpdateOne {
	iauo.mutation.SetConfiguration(b)
	return iauo
}

// ClearConfiguration clears the value of the "configuration" field.
func (iauo *IntegrationAttachmentUpdateOne) ClearConfiguration() *IntegrationAttachmentUpdateOne {
	iauo.mutation.ClearConfiguration()
	return iauo
}

// SetDeletedAt sets the "deleted_at" field.
func (iauo *IntegrationAttachmentUpdateOne) SetDeletedAt(t time.Time) *IntegrationAttachmentUpdateOne {
	iauo.mutation.SetDeletedAt(t)
	return iauo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (iauo *IntegrationAttachmentUpdateOne) SetNillableDeletedAt(t *time.Time) *IntegrationAttachmentUpdateOne {
	if t != nil {
		iauo.SetDeletedAt(*t)
	}
	return iauo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (iauo *IntegrationAttachmentUpdateOne) ClearDeletedAt() *IntegrationAttachmentUpdateOne {
	iauo.mutation.ClearDeletedAt()
	return iauo
}

// SetIntegrationID sets the "integration" edge to the Integration entity by ID.
func (iauo *IntegrationAttachmentUpdateOne) SetIntegrationID(id uuid.UUID) *IntegrationAttachmentUpdateOne {
	iauo.mutation.SetIntegrationID(id)
	return iauo
}

// SetIntegration sets the "integration" edge to the Integration entity.
func (iauo *IntegrationAttachmentUpdateOne) SetIntegration(i *Integration) *IntegrationAttachmentUpdateOne {
	return iauo.SetIntegrationID(i.ID)
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (iauo *IntegrationAttachmentUpdateOne) SetWorkflowID(id uuid.UUID) *IntegrationAttachmentUpdateOne {
	iauo.mutation.SetWorkflowID(id)
	return iauo
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (iauo *IntegrationAttachmentUpdateOne) SetWorkflow(w *Workflow) *IntegrationAttachmentUpdateOne {
	return iauo.SetWorkflowID(w.ID)
}

// Mutation returns the IntegrationAttachmentMutation object of the builder.
func (iauo *IntegrationAttachmentUpdateOne) Mutation() *IntegrationAttachmentMutation {
	return iauo.mutation
}

// ClearIntegration clears the "integration" edge to the Integration entity.
func (iauo *IntegrationAttachmentUpdateOne) ClearIntegration() *IntegrationAttachmentUpdateOne {
	iauo.mutation.ClearIntegration()
	return iauo
}

// ClearWorkflow clears the "workflow" edge to the Workflow entity.
func (iauo *IntegrationAttachmentUpdateOne) ClearWorkflow() *IntegrationAttachmentUpdateOne {
	iauo.mutation.ClearWorkflow()
	return iauo
}

// Where appends a list predicates to the IntegrationAttachmentUpdate builder.
func (iauo *IntegrationAttachmentUpdateOne) Where(ps ...predicate.IntegrationAttachment) *IntegrationAttachmentUpdateOne {
	iauo.mutation.Where(ps...)
	return iauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iauo *IntegrationAttachmentUpdateOne) Select(field string, fields ...string) *IntegrationAttachmentUpdateOne {
	iauo.fields = append([]string{field}, fields...)
	return iauo
}

// Save executes the query and returns the updated IntegrationAttachment entity.
func (iauo *IntegrationAttachmentUpdateOne) Save(ctx context.Context) (*IntegrationAttachment, error) {
	return withHooks(ctx, iauo.sqlSave, iauo.mutation, iauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iauo *IntegrationAttachmentUpdateOne) SaveX(ctx context.Context) *IntegrationAttachment {
	node, err := iauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iauo *IntegrationAttachmentUpdateOne) Exec(ctx context.Context) error {
	_, err := iauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iauo *IntegrationAttachmentUpdateOne) ExecX(ctx context.Context) {
	if err := iauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iauo *IntegrationAttachmentUpdateOne) check() error {
	if _, ok := iauo.mutation.IntegrationID(); iauo.mutation.IntegrationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "IntegrationAttachment.integration"`)
	}
	if _, ok := iauo.mutation.WorkflowID(); iauo.mutation.WorkflowCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "IntegrationAttachment.workflow"`)
	}
	return nil
}

func (iauo *IntegrationAttachmentUpdateOne) sqlSave(ctx context.Context) (_node *IntegrationAttachment, err error) {
	if err := iauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(integrationattachment.Table, integrationattachment.Columns, sqlgraph.NewFieldSpec(integrationattachment.FieldID, field.TypeUUID))
	id, ok := iauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "IntegrationAttachment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, integrationattachment.FieldID)
		for _, f := range fields {
			if !integrationattachment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != integrationattachment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iauo.mutation.Configuration(); ok {
		_spec.SetField(integrationattachment.FieldConfiguration, field.TypeBytes, value)
	}
	if iauo.mutation.ConfigurationCleared() {
		_spec.ClearField(integrationattachment.FieldConfiguration, field.TypeBytes)
	}
	if value, ok := iauo.mutation.DeletedAt(); ok {
		_spec.SetField(integrationattachment.FieldDeletedAt, field.TypeTime, value)
	}
	if iauo.mutation.DeletedAtCleared() {
		_spec.ClearField(integrationattachment.FieldDeletedAt, field.TypeTime)
	}
	if iauo.mutation.IntegrationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   integrationattachment.IntegrationTable,
			Columns: []string{integrationattachment.IntegrationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integration.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iauo.mutation.IntegrationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   integrationattachment.IntegrationTable,
			Columns: []string{integrationattachment.IntegrationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integration.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iauo.mutation.WorkflowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   integrationattachment.WorkflowTable,
			Columns: []string{integrationattachment.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iauo.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   integrationattachment.WorkflowTable,
			Columns: []string{integrationattachment.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &IntegrationAttachment{config: iauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{integrationattachment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iauo.mutation.done = true
	return _node, nil
}
