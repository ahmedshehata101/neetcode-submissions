func searchMatrix(matrix [][]int, target int) bool {
m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])

	// Step 1: Find candidate row
	top, bottom := 0, m-1
	rowIndex := -1

	for top <= bottom {
		mid := top + (bottom-top)/2
		if target < matrix[mid][0] {
			bottom = mid - 1
		} else if target > matrix[mid][n-1] {
			top = mid + 1
		} else {
			rowIndex = mid
			break
		}
	}

	if rowIndex == -1 {
		return false
	}

	// Step 2: Binary search in the row
	row := matrix[rowIndex]
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if row[mid] == target {
			return true
		} else if row[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
