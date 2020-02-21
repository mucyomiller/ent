// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/card"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/predicate"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/spec"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/user"
)

// CardUpdate is the builder for updating Card entities.
type CardUpdate struct {
	config
	hooks      []ent.Hook
	mutation   *CardMutation
	predicates []predicate.Card
}

// Where adds a new predicate for the builder.
func (cu *CardUpdate) Where(ps ...predicate.Card) *CardUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetName sets the name field.
func (cu *CardUpdate) SetName(s string) *CardUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the name field if the given value is not nil.
func (cu *CardUpdate) SetNillableName(s *string) *CardUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// ClearName clears the value of name.
func (cu *CardUpdate) ClearName() *CardUpdate {
	cu.mutation.ClearName()
	return cu
}

// SetOwnerID sets the owner edge to User by id.
func (cu *CardUpdate) SetOwnerID(id string) *CardUpdate {
	cu.mutation.SetOwnerID(id)
	return cu
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (cu *CardUpdate) SetNillableOwnerID(id *string) *CardUpdate {
	if id != nil {
		cu = cu.SetOwnerID(*id)
	}
	return cu
}

// SetOwner sets the owner edge to User.
func (cu *CardUpdate) SetOwner(u *User) *CardUpdate {
	return cu.SetOwnerID(u.ID)
}

// AddSpecIDs adds the spec edge to Spec by ids.
func (cu *CardUpdate) AddSpecIDs(ids ...string) *CardUpdate {
	cu.mutation.AddSpecIDs(ids...)
	return cu
}

// AddSpec adds the spec edges to Spec.
func (cu *CardUpdate) AddSpec(s ...*Spec) *CardUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.AddSpecIDs(ids...)
}

// ClearOwner clears the owner edge to User.
func (cu *CardUpdate) ClearOwner() *CardUpdate {
	cu.mutation.ClearOwner()
	return cu
}

// RemoveSpecIDs removes the spec edge to Spec by ids.
func (cu *CardUpdate) RemoveSpecIDs(ids ...string) *CardUpdate {
	cu.mutation.RemoveSpecIDs(ids...)
	return cu
}

// RemoveSpec removes spec edges to Spec.
func (cu *CardUpdate) RemoveSpec(s ...*Spec) *CardUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.RemoveSpecIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CardUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := cu.mutation.UpdateTime(); !ok {
		v := card.UpdateDefaultUpdateTime()
		cu.mutation.SetUpdateTime(v)
	}
	if v, ok := cu.mutation.Name(); ok {
		if err := card.NameValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if len(cu.mutation.OwnerIDs()) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.gremlinSave(ctx)
	} else {
		var mut ent.Mutator = ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mutation, ok := m.(*CardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.gremlinSave(ctx)
			return affected, err
		})
		for _, hook := range cu.hooks {
			mut = hook(mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CardUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CardUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CardUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CardUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := cu.gremlin().Query()
	if err := cu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	return res.ReadInt()
}

func (cu *CardUpdate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.V().HasLabel(card.Label)
	for _, p := range cu.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := cu.mutation.UpdateTime(); ok {
		v.Property(dsl.Single, card.FieldUpdateTime, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		v.Property(dsl.Single, card.FieldName, value)
	}
	var properties []interface{}
	if cu.mutation.NameCleared() {
		properties = append(properties, card.FieldName)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if cu.mutation.OwnerCleared() {
		tr := rv.Clone().InE(user.CardLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range cu.mutation.OwnerIDs() {
		v.AddE(user.CardLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.CardLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(card.Label, user.CardLabel, id)),
		})
	}
	for _, id := range cu.mutation.RemovedSpecIDs() {
		tr := rv.Clone().InE(spec.CardLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range cu.mutation.SpecIDs() {
		v.AddE(spec.CardLabel).From(g.V(id)).InV()
	}
	v.Count()
	if len(constraints) > 0 {
		constraints = append(constraints, &constraint{
			pred: rv.Count(),
			test: __.Is(p.GT(1)).Constant(&ConstraintError{msg: "update traversal contains more than one vertex"}),
		})
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// CardUpdateOne is the builder for updating a single Card entity.
type CardUpdateOne struct {
	config
	hooks    []ent.Hook
	mutation *CardMutation
}

// SetName sets the name field.
func (cuo *CardUpdateOne) SetName(s string) *CardUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the name field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableName(s *string) *CardUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// ClearName clears the value of name.
func (cuo *CardUpdateOne) ClearName() *CardUpdateOne {
	cuo.mutation.ClearName()
	return cuo
}

// SetOwnerID sets the owner edge to User by id.
func (cuo *CardUpdateOne) SetOwnerID(id string) *CardUpdateOne {
	cuo.mutation.SetOwnerID(id)
	return cuo
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableOwnerID(id *string) *CardUpdateOne {
	if id != nil {
		cuo = cuo.SetOwnerID(*id)
	}
	return cuo
}

// SetOwner sets the owner edge to User.
func (cuo *CardUpdateOne) SetOwner(u *User) *CardUpdateOne {
	return cuo.SetOwnerID(u.ID)
}

// AddSpecIDs adds the spec edge to Spec by ids.
func (cuo *CardUpdateOne) AddSpecIDs(ids ...string) *CardUpdateOne {
	cuo.mutation.AddSpecIDs(ids...)
	return cuo
}

// AddSpec adds the spec edges to Spec.
func (cuo *CardUpdateOne) AddSpec(s ...*Spec) *CardUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.AddSpecIDs(ids...)
}

// ClearOwner clears the owner edge to User.
func (cuo *CardUpdateOne) ClearOwner() *CardUpdateOne {
	cuo.mutation.ClearOwner()
	return cuo
}

// RemoveSpecIDs removes the spec edge to Spec by ids.
func (cuo *CardUpdateOne) RemoveSpecIDs(ids ...string) *CardUpdateOne {
	cuo.mutation.RemoveSpecIDs(ids...)
	return cuo
}

// RemoveSpec removes spec edges to Spec.
func (cuo *CardUpdateOne) RemoveSpec(s ...*Spec) *CardUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.RemoveSpecIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (cuo *CardUpdateOne) Save(ctx context.Context) (*Card, error) {
	if _, ok := cuo.mutation.UpdateTime(); !ok {
		v := card.UpdateDefaultUpdateTime()
		cuo.mutation.SetUpdateTime(v)
	}
	if v, ok := cuo.mutation.Name(); ok {
		if err := card.NameValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if len(cuo.mutation.OwnerIDs()) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	var (
		err  error
		node *Card
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.gremlinSave(ctx)
	} else {
		var mut ent.Mutator = ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mutation, ok := m.(*CardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.gremlinSave(ctx)
			return node, err
		})
		for _, hook := range cuo.hooks {
			mut = hook(mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CardUpdateOne) SaveX(ctx context.Context) *Card {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *CardUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CardUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CardUpdateOne) gremlinSave(ctx context.Context) (*Card, error) {
	res := &gremlin.Response{}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Card.ID for update")
	}
	query, bindings := cuo.gremlin(id).Query()
	if err := cuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	c := &Card{config: cuo.config}
	if err := c.FromResponse(res); err != nil {
		return nil, err
	}
	return c, nil
}

func (cuo *CardUpdateOne) gremlin(id string) *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := cuo.mutation.UpdateTime(); ok {
		v.Property(dsl.Single, card.FieldUpdateTime, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		v.Property(dsl.Single, card.FieldName, value)
	}
	var properties []interface{}
	if cuo.mutation.NameCleared() {
		properties = append(properties, card.FieldName)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if cuo.mutation.OwnerCleared() {
		tr := rv.Clone().InE(user.CardLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range cuo.mutation.OwnerIDs() {
		v.AddE(user.CardLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.CardLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(card.Label, user.CardLabel, id)),
		})
	}
	for _, id := range cuo.mutation.RemovedSpecIDs() {
		tr := rv.Clone().InE(spec.CardLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range cuo.mutation.SpecIDs() {
		v.AddE(spec.CardLabel).From(g.V(id)).InV()
	}
	v.ValueMap(true)
	if len(constraints) > 0 {
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
