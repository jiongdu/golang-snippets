package data_structure

//golang对于生成多维数组的支持不太好
//不能根据变量直接指定（只能使用常量）
//以3维数组为例
func generator(i,j,k int) [][][]int {
	nums := make([][][]int, i)
	for m := 0; m < i; m ++ {
		nums[m] = make([][]int, j)
		for n := 0; n < j; n ++ {
			nums[m][n] = make([]int, k)
		}
	}
	return nums
}
