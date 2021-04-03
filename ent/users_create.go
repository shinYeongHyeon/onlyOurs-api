// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shinYeongHyeon/onlyOurs-api/ent/users"
)

// UsersCreate is the builder for creating a Users entity.
type UsersCreate struct {
	config
	mutation *UsersMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (uc *UsersCreate) SetUserID(s string) *UsersCreate {
	uc.mutation.SetUserID(s)
	return uc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uc *UsersCreate) SetNillableUserID(s *string) *UsersCreate {
	if s != nil {
		uc.SetUserID(*s)
	}
	return uc
}

// SetName sets the "name" field.
func (uc *UsersCreate) SetName(s string) *UsersCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uc *UsersCreate) SetNillableName(s *string) *UsersCreate {
	if s != nil {
		uc.SetName(*s)
	}
	return uc
}

// SetPassword sets the "password" field.
func (uc *UsersCreate) SetPassword(s string) *UsersCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uc *UsersCreate) SetNillablePassword(s *string) *UsersCreate {
	if s != nil {
		uc.SetPassword(*s)
	}
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UsersCreate) SetCreatedAt(t time.Time) *UsersCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UsersCreate) SetNillableCreatedAt(t *time.Time) *UsersCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UsersCreate) SetID(s string) *UsersCreate {
	uc.mutation.SetID(s)
	return uc
}

// Mutation returns the UsersMutation object of the builder.
func (uc *UsersCreate) Mutation() *UsersMutation {
	return uc.mutation
}

// Save creates the Users in the database.
func (uc *UsersCreate) Save(ctx context.Context) (*Users, error) {
	var (
		err  error
		node *Users
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UsersMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UsersCreate) SaveX(ctx context.Context) *Users {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (uc *UsersCreate) defaults() {
	if _, ok := uc.mutation.UserID(); !ok {
		v := users.DefaultUserID
		uc.mutation.SetUserID(v)
	}
	if _, ok := uc.mutation.Name(); !ok {
		v := users.DefaultName
		uc.mutation.SetName(v)
	}
	if _, ok := uc.mutation.Password(); !ok {
		v := users.DefaultPassword
		uc.mutation.SetPassword(v)
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := users.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UsersCreate) check() error {
	if _, ok := uc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New("ent: missing required field \"user_id\"")}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New("ent: missing required field \"password\"")}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	return nil
}

func (uc *UsersCreate) sqlSave(ctx context.Context) (*Users, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (uc *UsersCreate) createSpec() (*Users, *sqlgraph.CreateSpec) {
	var (
		_node = &Users{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: users.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: users.FieldID,
			},
		}
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldName,
		})
		_node.Name = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: users.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	return _node, _spec
}

// UsersCreateBulk is the builder for creating many Users entities in bulk.
type UsersCreateBulk struct {
	config
	builders []*UsersCreate
}

// Save creates the Users entities in the database.
func (ucb *UsersCreateBulk) Save(ctx context.Context) ([]*Users, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*Users, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UsersMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UsersCreateBulk) SaveX(ctx context.Context) []*Users {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
