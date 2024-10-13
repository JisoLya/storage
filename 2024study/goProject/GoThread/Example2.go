package main

func main() {

	a := []int{1, 2, 3, 4, 5, 7, 9}
	println(binary(10, a))
}
func binary(tar int, array []int) int {
	begin := 0
	end := len(array) - 1
	var mid int
	for begin <= end {
		mid = (end-begin)/2 + begin
		if array[mid] > tar {
			end = mid - 1
		} else if array[mid] < tar {
			begin = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
