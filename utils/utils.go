package utils

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func contains[T comparable](input []T, value T) bool {
	for _, v := range input {
		if v == value {
			return true
		}
	}
	return false
}

func move_element[T any](input []T, fromIndex int, toIndex int) []T {

	element := input[fromIndex]
	input = append(input[:fromIndex], input[fromIndex+1:]...)
	input = append(input[:toIndex], append([]T{element}, input[toIndex:]...)...)

	return input
}