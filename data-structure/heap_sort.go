//堆排序
//基本思想：将待排序序列构造成一个大顶堆，此时，整个序列的最大值就是堆顶的根节点。
//将其与末尾元素进行交换，此时末尾就为最大值。然后将剩余n-1个元素重新构造成一个堆，
//这样会得到n个元素的次小值。如此反复执行，便能得到一个有序序列了
func sortArray(nums []int) []int {
	length := len(nums)
	//先构造一个大顶堆
	for i := length/2 - 1; i >= 0; i-- {
		adjustHeap(nums, i, length)
	}
	//然后在交换root和数组最后一个位置，并调整结构
	for i := length - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		adjustHeap(nums, 0, i)
	}
	return nums
}

func adjustHeap(nums []int, start, length int) {
	current := nums[start]
	child := 2*start + 1
	for child < length {
		if child+1 < length && nums[child] < nums[child+1] {
			child++
		}
		if current < nums[child] {
			nums[start] = nums[child]
			start = child
			child = start*2 + 1
		} else {
			break
		}
	}
	nums[start] = current
}