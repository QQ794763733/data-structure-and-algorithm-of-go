/*
希尔排序：希尔排序属于插入类排序的一种，先把元素分组，然后分组比较按排序规则排序，最终使其有序。
主要思想：
	1. 先把需要排序的元素分组。
	2. 按排序规则把分组内的元素进行比较，交换。
	3. 每一次都会排序好所有分组。
	4. 每一趟都会细分分组。
*/
package insertion

// 希尔排序
func ShellSort(array []int) []int {
	var temp int
	for half := len(array)/2; half > 0 ; half/=2 {
		for i := half; i < len(array); i++ {
			for j := i - half; j >= 0 ; j -= half {
				if array[j] > array[j+half] {
					temp = array[j]
					array[j] = array[j+half]
					array[j+half] = temp
				}
			}
		}
	}
	return array
}