// Code generated by ent, DO NOT EDIT.

package ent

import (
	"golang-boilerplate/ent/request"
	"golang-boilerplate/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	requestFields := schema.Request{}.Fields()
	_ = requestFields
	// requestDescCreatedAt is the schema descriptor for created_at field.
	requestDescCreatedAt := requestFields[6].Descriptor()
	// request.DefaultCreatedAt holds the default value on creation for the created_at field.
	request.DefaultCreatedAt = requestDescCreatedAt.Default.(func() time.Time)
	// requestDescUpdatedAt is the schema descriptor for updated_at field.
	requestDescUpdatedAt := requestFields[7].Descriptor()
	// request.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	request.DefaultUpdatedAt = requestDescUpdatedAt.Default.(func() time.Time)
	// request.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	request.UpdateDefaultUpdatedAt = requestDescUpdatedAt.UpdateDefault.(func() time.Time)
}
