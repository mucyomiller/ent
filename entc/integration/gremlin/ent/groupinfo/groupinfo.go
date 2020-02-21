// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package groupinfo

import "github.com/facebookincubator/ent/entc/integration/ent/schema"

const (
	// Label holds the string label denoting the groupinfo type in the database.
	Label = "group_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID       = "id"   // FieldDesc holds the string denoting the desc vertex property in the database.
	FieldDesc     = "desc" // FieldMaxUsers holds the string denoting the max_users vertex property in the database.
	FieldMaxUsers = "max_users"

	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"

	// GroupsInverseLabel holds the string label denoting the groups inverse edge type in the database.
	GroupsInverseLabel = "group_info"
)

var (
	fields = schema.GroupInfo{}.Fields()
	// descMaxUsers is the schema descriptor for max_users field.
	descMaxUsers = fields[1].Descriptor()
	// DefaultMaxUsers holds the default value on creation for the max_users field.
	DefaultMaxUsers = descMaxUsers.Default.(int)
)
