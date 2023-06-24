// Code generated by ent, DO NOT EDIT.

package database

import (
	"ocfcore/internal/database/node"
	"ocfcore/internal/database/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	nodeFields := schema.Node{}.Fields()
	_ = nodeFields
	// nodeDescPeerId is the schema descriptor for peerId field.
	nodeDescPeerId := nodeFields[0].Descriptor()
	// node.DefaultPeerId holds the default value on creation for the peerId field.
	node.DefaultPeerId = nodeDescPeerId.Default.(string)
	// nodeDescStatus is the schema descriptor for status field.
	nodeDescStatus := nodeFields[1].Descriptor()
	// node.DefaultStatus holds the default value on creation for the status field.
	node.DefaultStatus = nodeDescStatus.Default.(string)
}