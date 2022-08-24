package helpers

import "calling-bill/constants"

//CalculateBlockCount calculates block count from call's duration
func CalculateBlockCount(duration int) int {
	if duration%constants.CallBlockSize == 0 {
		return duration / constants.CallBlockSize
	}
	return duration/constants.CallBlockSize + 1
}
