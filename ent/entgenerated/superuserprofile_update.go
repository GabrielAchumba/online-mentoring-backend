// Code generated by ent, DO NOT EDIT.

package entgenerated

import (
	"api/ent/entgenerated/predicate"
	"api/ent/entgenerated/superuserprofile"
	"api/ent/entgenerated/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SuperuserProfileUpdate is the builder for updating SuperuserProfile entities.
type SuperuserProfileUpdate struct {
	config
	hooks    []Hook
	mutation *SuperuserProfileMutation
}

// Where appends a list predicates to the SuperuserProfileUpdate builder.
func (spu *SuperuserProfileUpdate) Where(ps ...predicate.SuperuserProfile) *SuperuserProfileUpdate {
	spu.mutation.Where(ps...)
	return spu
}

// SetOwnerID sets the "owner_id" field.
func (spu *SuperuserProfileUpdate) SetOwnerID(i int) *SuperuserProfileUpdate {
	spu.mutation.SetOwnerID(i)
	return spu
}

// SetOwner sets the "owner" edge to the User entity.
func (spu *SuperuserProfileUpdate) SetOwner(u *User) *SuperuserProfileUpdate {
	return spu.SetOwnerID(u.ID)
}

// Mutation returns the SuperuserProfileMutation object of the builder.
func (spu *SuperuserProfileUpdate) Mutation() *SuperuserProfileMutation {
	return spu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (spu *SuperuserProfileUpdate) ClearOwner() *SuperuserProfileUpdate {
	spu.mutation.ClearOwner()
	return spu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (spu *SuperuserProfileUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, spu.sqlSave, spu.mutation, spu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (spu *SuperuserProfileUpdate) SaveX(ctx context.Context) int {
	affected, err := spu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (spu *SuperuserProfileUpdate) Exec(ctx context.Context) error {
	_, err := spu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spu *SuperuserProfileUpdate) ExecX(ctx context.Context) {
	if err := spu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (spu *SuperuserProfileUpdate) check() error {
	if _, ok := spu.mutation.OwnerID(); spu.mutation.OwnerCleared() && !ok {
		return errors.New(`entgenerated: clearing a required unique edge "SuperuserProfile.owner"`)
	}
	return nil
}

func (spu *SuperuserProfileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := spu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(superuserprofile.Table, superuserprofile.Columns, sqlgraph.NewFieldSpec(superuserprofile.FieldID, field.TypeInt))
	if ps := spu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if spu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   superuserprofile.OwnerTable,
			Columns: []string{superuserprofile.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   superuserprofile.OwnerTable,
			Columns: []string{superuserprofile.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, spu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{superuserprofile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	spu.mutation.done = true
	return n, nil
}

// SuperuserProfileUpdateOne is the builder for updating a single SuperuserProfile entity.
type SuperuserProfileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SuperuserProfileMutation
}

// SetOwnerID sets the "owner_id" field.
func (spuo *SuperuserProfileUpdateOne) SetOwnerID(i int) *SuperuserProfileUpdateOne {
	spuo.mutation.SetOwnerID(i)
	return spuo
}

// SetOwner sets the "owner" edge to the User entity.
func (spuo *SuperuserProfileUpdateOne) SetOwner(u *User) *SuperuserProfileUpdateOne {
	return spuo.SetOwnerID(u.ID)
}

// Mutation returns the SuperuserProfileMutation object of the builder.
func (spuo *SuperuserProfileUpdateOne) Mutation() *SuperuserProfileMutation {
	return spuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (spuo *SuperuserProfileUpdateOne) ClearOwner() *SuperuserProfileUpdateOne {
	spuo.mutation.ClearOwner()
	return spuo
}

// Where appends a list predicates to the SuperuserProfileUpdate builder.
func (spuo *SuperuserProfileUpdateOne) Where(ps ...predicate.SuperuserProfile) *SuperuserProfileUpdateOne {
	spuo.mutation.Where(ps...)
	return spuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (spuo *SuperuserProfileUpdateOne) Select(field string, fields ...string) *SuperuserProfileUpdateOne {
	spuo.fields = append([]string{field}, fields...)
	return spuo
}

// Save executes the query and returns the updated SuperuserProfile entity.
func (spuo *SuperuserProfileUpdateOne) Save(ctx context.Context) (*SuperuserProfile, error) {
	return withHooks(ctx, spuo.sqlSave, spuo.mutation, spuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (spuo *SuperuserProfileUpdateOne) SaveX(ctx context.Context) *SuperuserProfile {
	node, err := spuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (spuo *SuperuserProfileUpdateOne) Exec(ctx context.Context) error {
	_, err := spuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spuo *SuperuserProfileUpdateOne) ExecX(ctx context.Context) {
	if err := spuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (spuo *SuperuserProfileUpdateOne) check() error {
	if _, ok := spuo.mutation.OwnerID(); spuo.mutation.OwnerCleared() && !ok {
		return errors.New(`entgenerated: clearing a required unique edge "SuperuserProfile.owner"`)
	}
	return nil
}

func (spuo *SuperuserProfileUpdateOne) sqlSave(ctx context.Context) (_node *SuperuserProfile, err error) {
	if err := spuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(superuserprofile.Table, superuserprofile.Columns, sqlgraph.NewFieldSpec(superuserprofile.FieldID, field.TypeInt))
	id, ok := spuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entgenerated: missing "SuperuserProfile.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := spuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, superuserprofile.FieldID)
		for _, f := range fields {
			if !superuserprofile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entgenerated: invalid field %q for query", f)}
			}
			if f != superuserprofile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := spuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if spuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   superuserprofile.OwnerTable,
			Columns: []string{superuserprofile.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   superuserprofile.OwnerTable,
			Columns: []string{superuserprofile.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SuperuserProfile{config: spuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, spuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{superuserprofile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	spuo.mutation.done = true
	return _node, nil
}
