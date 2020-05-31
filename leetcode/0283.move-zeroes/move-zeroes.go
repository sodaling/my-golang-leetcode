package problem0283

func moveZeroes2(nums []int) {
	// 非零的先抽出来
	NotZeroAlice := make([]int, 0, len(nums))
	for i := range nums {
		if nums[i] != 0 {
			NotZeroAlice = append(NotZeroAlice, nums[i])
		}
	}
	for i := range NotZeroAlice {
		nums[i] = NotZeroAlice[i]
	}
	for i := len(NotZeroAlice); i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeroes3(nums []int) {
	var pre int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[pre] = nums[i]
			pre++
		}
	}
	for ; pre < len(nums); pre++ {
		nums[pre] = 0
	}

}

func moveZeroes(nums []int) {
	var pre int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// 下面判断是优化,不等的时候才进行交换操作
			if i != pre {
				nums[pre], nums[i] = nums[i], nums[pre]
			}
			pre++
		}
	}
}
