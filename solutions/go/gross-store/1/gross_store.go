package gross
// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    units := map[string]int{
        "quarter_of_a_dozen": 3,
		"half_of_a_dozen": 6,
		"dozen": 12,
		"small_gross": 120,
		"gross": 144,
		"great_gross": 1728,
    }
    return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    bill := map[string]int{}
    return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    quantity, exists := units[unit]
    if !exists {
        return false
    }
    bill[item] += quantity
    return true

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// 1. Check if item exists in bill
	currentQty, ok := bill[item]
	if !ok {
		return false
	}
	// 2. Check if unit exists
	unitQty, ok := units[unit]
	if !ok {
		return false
	}
	// 3. Compute new quantity
	newQty := currentQty - unitQty
	// 4. If negative → invalid
	if newQty < 0 {
		return false
	}
	// 5. If zero → remove item
	if newQty == 0 {
		delete(bill, item)
		return true
	}
	// 6. Otherwise update quantity
	bill[item] = newQty
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    qty, ok := bill[item] 
    if !ok {
        return 0, false
    }
    return qty, true
}
