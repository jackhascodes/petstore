package pet

// Inventory is an alias of map to enable encapsulation.
type Inventory map[string]int32

func InitInventory() Inventory {
	return map[string]int32{}
}

func (i Inventory) AddItemCount(name string, count int32) {
	i[name] = count
}
