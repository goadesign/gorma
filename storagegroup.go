package gorma

import (
	"fmt"
	"sort"
)

// IterateStores runs an iterator function once per Relational Store in the StorageGroup's Store list
func (sd *StorageGroupDefinition) IterateStores(it StoreIterator) error {
	names := make([]string, len(sd.RelationalStores))
	i := 0
	for n := range sd.RelationalStores {
		names[i] = n
		i++
	}
	sort.Strings(names)
	for _, n := range names {
		if err := it(sd.RelationalStores[n]); err != nil {
			return err
		}
	}
	return nil
}

// Context returns the generic definition name used in error messages.
func (a *StorageGroupDefinition) Context() string {
	if a.Name != "" {
		return fmt.Sprintf("StorageGroup %#v", a.Name)
	}
	return "unnamed Storage Group"
}
