package listops

// ListMinus return the items from list1,except of the items in list2
func ListMinus(list1, list2 []interface{}) []interface{} {
	list := make([]interface{}, len(list1))
	copy(list, list1)

	for _, item := range list2 {
		if IsInList(list, item) {
			list = RmListItem(list, item)
		}
	}

	return list
}

// UniqueList return a list that only keep the unique obj from list1
func UniqueList(list []interface{}) []interface{} {
	newList := make([]interface{}, 0)

	for _, item := range list {
		if !IsInList(newList, item) {
			newList = append(newList, item)
		}
	}

	return newList
}

// ListAnd do and operation for 2 lists
func ListAnd(list1, list2 []interface{}) []interface{} {
	newList := make([]interface{}, 0)
	for _, item := range list1 {
		if IsInList(list2, item) && !IsInList(newList, item) {
			newList = append(newList, item)
		}
	}
	return newList
}

// IsInList Judge whether an item in list or not
func IsInList(list []interface{}, item interface{}) bool {
	for _, vItem := range list {
		if vItem == item {
			return true
		}
	}
	return false
}

// RmListItem remove item from list
func RmListItem(list []interface{}, item interface{}) []interface{} {
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] == item {
			if i == len(list)-1 {
				list = list[:i]
			} else {
				list = append(list[:i], list[i+1:]...)
			}
		}
	}
	return list
}
