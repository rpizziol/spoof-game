package util

/**
 * Return absolute value of an integer.
 * @param x	The input integer.
 * @return	The absolute value of x.
 */
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/**
 * Check whether or not an element is in a given list.
 * @param list	The list to check.
 * @param elem	The element to find.
 * @return		The outcome of the check.
 */
func IsInList(list []int, elem int) bool {
	for _, a := range list {
		if a == elem {
			return true
		}
	}
	return false
}
