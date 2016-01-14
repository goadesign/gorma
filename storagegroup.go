package gorma

import "github.com/raphael/goa/design"

// NewStorageGroup creates a StorageGroup structure by parsing
// the APIDefinition and creating all the necessary Stores and Models
func NewStorageGroup(a *design.APIDefinition) *StorageGroup {
	sg := &StorageGroup{}

	return sg
}
