// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/examples/m2m2types/ent/migrate"

	"github.com/facebookincubator/ent/examples/m2m2types/ent/group"
	"github.com/facebookincubator/ent/examples/m2m2types/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	c := config{log: log.Println}
	c.options(opts...)
	return &Client{
		config: c,
		Schema: migrate.NewSchema(c.driver),
		Group:  NewGroupClient(c),
		User:   NewUserClient(c),
	}
}

// Open opens a connection to the database specified by the driver name and a
// driver-specific data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug}
	return &Tx{
		config: cfg,
		Group:  NewGroupClient(cfg),
		User:   NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Group.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true}
	return &Client{
		config: cfg,
		Schema: migrate.NewSchema(cfg.driver),
		Group:  NewGroupClient(cfg),
		User:   NewUserClient(cfg),
	}
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...ent.Hook) {
	c.Group.Use(hooks...)
	c.User.Use(hooks...)
}

// GroupClient is a client for the Group schema.
type GroupClient struct {
	config
}

// NewGroupClient returns a client for the Group from the given config.
func NewGroupClient(c config) *GroupClient {
	return &GroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `group.Hooks(f(g(h())))`.
func (c *GroupClient) Use(hooks ...ent.Hook) {
	c.hooks = append(c.hooks[:len(c.hooks):len(c.hooks)], hooks...)
}

// Create returns a create builder for Group.
func (c *GroupClient) Create() *GroupCreate {
	hooks := c.hooks
	mutation := newGroupMutation(OpCreate)
	return &GroupCreate{config: c.config, hooks: hooks, mutation: mutation}
}

// Update returns an update builder for Group.
func (c *GroupClient) Update() *GroupUpdate {
	hooks := c.hooks
	mutation := newGroupMutation(OpUpdate)
	return &GroupUpdate{config: c.config, hooks: hooks, mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupClient) UpdateOne(gr *Group) *GroupUpdateOne {
	return c.UpdateOneID(gr.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupClient) UpdateOneID(id int) *GroupUpdateOne {
	hooks := c.hooks
	mutation := newGroupMutation(OpUpdateOne)
	mutation.id = &id
	return &GroupUpdateOne{config: c.config, hooks: hooks, mutation: mutation}
}

// Delete returns a delete builder for Group.
func (c *GroupClient) Delete() *GroupDelete {
	hooks := c.hooks
	mutation := newGroupMutation(OpDelete)
	return &GroupDelete{config: c.config, hooks: hooks, mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GroupClient) DeleteOne(gr *Group) *GroupDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GroupClient) DeleteOneID(id int) *GroupDeleteOne {
	builder := c.Delete().Where(group.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupDeleteOne{builder}
}

// Create returns a query builder for Group.
func (c *GroupClient) Query() *GroupQuery {
	return &GroupQuery{config: c.config}
}

// Get returns a Group entity by its id.
func (c *GroupClient) Get(ctx context.Context, id int) (*Group, error) {
	return c.Query().Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupClient) GetX(ctx context.Context, id int) *Group {
	gr, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return gr
}

// QueryUsers queries the users edge of a Group.
func (c *GroupClient) QueryUsers(gr *Group) *UserQuery {
	query := &UserQuery{config: c.config}
	id := gr.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(group.Table, group.FieldID, id),
		sqlgraph.To(user.Table, user.FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, group.UsersTable, group.UsersPrimaryKey...),
	)
	query.sql = sqlgraph.Neighbors(gr.driver.Dialect(), step)

	return query
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...ent.Hook) {
	c.hooks = append(c.hooks[:len(c.hooks):len(c.hooks)], hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	hooks := c.hooks
	mutation := newUserMutation(OpCreate)
	return &UserCreate{config: c.config, hooks: hooks, mutation: mutation}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	hooks := c.hooks
	mutation := newUserMutation(OpUpdate)
	return &UserUpdate{config: c.config, hooks: hooks, mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	return c.UpdateOneID(u.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	hooks := c.hooks
	mutation := newUserMutation(OpUpdateOne)
	mutation.id = &id
	return &UserUpdateOne{config: c.config, hooks: hooks, mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	hooks := c.hooks
	mutation := newUserMutation(OpDelete)
	return &UserDelete{config: c.config, hooks: hooks, mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryGroups queries the groups edge of a User.
func (c *UserClient) QueryGroups(u *User) *GroupQuery {
	query := &GroupQuery{config: c.config}
	id := u.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(user.Table, user.FieldID, id),
		sqlgraph.To(group.Table, group.FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, user.GroupsTable, user.GroupsPrimaryKey...),
	)
	query.sql = sqlgraph.Neighbors(u.driver.Dialect(), step)

	return query
}
