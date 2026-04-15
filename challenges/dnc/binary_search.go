package dnc

func BinarySearch(list []int, search int) int {
	return RecursiveBinarySearch(list, 0, len(list)-1, search)
}

func RecursiveBinarySearch(list []int, start, end, search int) int {
	if start > end {
		return -1
	}

	mid := start + (end-start)/2
	value := list[mid]

	if value == search {
		return mid
	} else if value > search {
		return RecursiveBinarySearch(list, start, mid-1, search)
	} else {
		return RecursiveBinarySearch(list, mid+1, end, search)
	}
}
