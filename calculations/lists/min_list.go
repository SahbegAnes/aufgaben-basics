package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
func MinList(nums []int) int {

	if len(nums) == 0 {
		return 0
	} else {
		low := nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] < low {
				low = nums[i]
			}
		}
		return low
	}

}
