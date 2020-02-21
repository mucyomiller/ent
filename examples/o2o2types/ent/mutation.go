// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/examples/o2o2types/ent/card"
	"github.com/facebookincubator/ent/examples/o2o2types/ent/user"
)

const (
	// Operation types.
	OpCreate    = "Create"
	OpDelete    = "Delete"
	OpDeleteOne = "DeleteOne"
	OpUpdate    = "Update"
	OpUpdateOne = "UpdateOne"

	// Node types.
	TypeCard = "Card"
	TypeUser = "User"
)

// CardMutation represents an operation that mutate the Cards
// nodes in the graph.
type CardMutation struct {
	op, typ       string
	id            *int
	expired       *time.Time
	number        *string
	clearedFields map[string]bool
	owner         map[int]struct{}
	clearedowner  bool
}

var _ ent.Mutation = (*CardMutation)(nil)

// newCardMutation creates new mutation for $n.Name.
func newCardMutation(op string) *CardMutation {
	return &CardMutation{
		op:            op,
		typ:           TypeCard,
		clearedFields: make(map[string]bool),
	}
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *CardMutation) ID() (id int, _ bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetExpired sets the expired field.
func (m *CardMutation) SetExpired(t time.Time) {
	m.expired = &t
}

// Expired returns the expired value in the mutation.
func (m *CardMutation) Expired() (r time.Time, _ bool) {
	v := m.expired
	if v == nil {
		return
	}
	return *v, true
}

// ResetExpired reset all changes of the expired field.
func (m *CardMutation) ResetExpired() {
	m.expired = nil
}

// SetNumber sets the number field.
func (m *CardMutation) SetNumber(s string) {
	m.number = &s
}

// Number returns the number value in the mutation.
func (m *CardMutation) Number() (r string, _ bool) {
	v := m.number
	if v == nil {
		return
	}
	return *v, true
}

// ResetNumber reset all changes of the number field.
func (m *CardMutation) ResetNumber() {
	m.number = nil
}

// SetOwnerID sets the owner edge to User by id.
func (m *CardMutation) SetOwnerID(id int) {
	if m.owner == nil {
		m.owner = make(map[int]struct{})
	}
	m.owner[id] = struct{}{}
}

// ClearOwner clears the owner edge to User.
func (m *CardMutation) ClearOwner() {
	m.clearedowner = true
}

// OwnerCleared returns if the edge owner was cleared.
func (m *CardMutation) OwnerCleared() bool {
	return m.clearedowner
}

// OwnerIDs returns the owner ids in the mutation.
func (m *CardMutation) OwnerIDs() (ids []int) {
	for id := range m.owner {
		ids = append(ids, id)
	}
	return
}

// ResetOwner reset all changes of the owner edge.
func (m *CardMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// Op returns the operation name.
func (m *CardMutation) Op() string {
	return m.op
}

// Type returns the node type of this mutation (Card).
func (m *CardMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *CardMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.expired != nil {
		fields = append(fields, card.FieldExpired)
	}
	if m.number != nil {
		fields = append(fields, card.FieldNumber)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *CardMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case card.FieldExpired:
		return m.Expired()
	case card.FieldNumber:
		return m.Number()
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *CardMutation) SetField(name string, value ent.Value) error {
	switch name {
	case card.FieldExpired:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetExpired(v)
		return nil
	case card.FieldNumber:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNumber(v)
		return nil
	}
	return fmt.Errorf("unknown Card field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *CardMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *CardMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *CardMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Card numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *CardMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *CardMutation) FieldCleared(name string) bool {
	return m.clearedFields[name]
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *CardMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Card nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *CardMutation) ResetField(name string) error {
	switch name {
	case card.FieldExpired:
		m.ResetExpired()
		return nil
	case card.FieldNumber:
		m.ResetNumber()
		return nil
	}
	return fmt.Errorf("unknown Card field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *CardMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.owner != nil {
		edges = append(edges, card.EdgeOwner)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *CardMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case card.EdgeOwner:
		ids := make([]int, 0, len(m.owner))
		for id := range m.owner {
			ids = append(ids, id)
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *CardMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *CardMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *CardMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedowner {
		edges = append(edges, card.EdgeOwner)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *CardMutation) EdgeCleared(name string) bool {
	switch name {
	case card.EdgeOwner:
		return m.clearedowner
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *CardMutation) ClearEdge(name string) error {
	switch name {
	case card.EdgeOwner:
		m.ClearOwner()
		return nil
	}
	return fmt.Errorf("unknown Card unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *CardMutation) ResetEdge(name string) error {
	switch name {
	case card.EdgeOwner:
		m.ResetOwner()
		return nil
	}
	return fmt.Errorf("unknown Card edge %s", name)
}

// UserMutation represents an operation that mutate the Users
// nodes in the graph.
type UserMutation struct {
	op, typ       string
	id            *int
	age           *int
	addage        *int
	name          *string
	clearedFields map[string]bool
	card          map[int]struct{}
	clearedcard   bool
}

var _ ent.Mutation = (*UserMutation)(nil)

// newUserMutation creates new mutation for $n.Name.
func newUserMutation(op string) *UserMutation {
	return &UserMutation{
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]bool),
	}
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *UserMutation) ID() (id int, _ bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetAge sets the age field.
func (m *UserMutation) SetAge(i int) {
	m.age = &i
	m.addage = nil
}

// Age returns the age value in the mutation.
func (m *UserMutation) Age() (r int, _ bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// AddAge adds i to age.
func (m *UserMutation) AddAge(i int) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the age field in this mutation.
func (m *UserMutation) AddedAge() (r int, _ bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge reset all changes of the age field.
func (m *UserMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// SetName sets the name field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *UserMutation) Name() (r string, _ bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// ResetName reset all changes of the name field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetCardID sets the card edge to Card by id.
func (m *UserMutation) SetCardID(id int) {
	if m.card == nil {
		m.card = make(map[int]struct{})
	}
	m.card[id] = struct{}{}
}

// ClearCard clears the card edge to Card.
func (m *UserMutation) ClearCard() {
	m.clearedcard = true
}

// CardCleared returns if the edge card was cleared.
func (m *UserMutation) CardCleared() bool {
	return m.clearedcard
}

// CardIDs returns the card ids in the mutation.
func (m *UserMutation) CardIDs() (ids []int) {
	for id := range m.card {
		ids = append(ids, id)
	}
	return
}

// ResetCard reset all changes of the card edge.
func (m *UserMutation) ResetCard() {
	m.card = nil
	m.clearedcard = false
}

// Op returns the operation name.
func (m *UserMutation) Op() string {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.age != nil {
		fields = append(fields, user.FieldAge)
	}
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.Age()
	case user.FieldName:
		return m.Name()
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, user.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	return m.clearedFields[name]
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldAge:
		m.ResetAge()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.card != nil {
		edges = append(edges, user.EdgeCard)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeCard:
		ids := make([]int, 0, len(m.card))
		for id := range m.card {
			ids = append(ids, id)
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedcard {
		edges = append(edges, user.EdgeCard)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeCard:
		return m.clearedcard
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	case user.EdgeCard:
		m.ClearCard()
		return nil
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeCard:
		m.ResetCard()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
