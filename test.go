package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	//test_map()
	//test_func()
	//test_slice()
	//start := time.Now()
	//n := 40
	//hanoi(n, 1, 3, 2)
	//fmt.Println("start at:", start, "\n end at:", time.Now())
	//testEqual()
	//test := [][]int{
	//	{0, 0, 1},
	//	//{2, 2, 1},
	//	//{3, 3, 3, 3},
	//	//{3, 3, 7, 2, 5, 8, 4, 6, 8, 1},
	//	//{6, 1, 1, 9, 8, 11},
	//	//{9, 1, 4, 7, 3, 21, 88, 5, 8, 11, 6},
	//}

	test := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	fmt.Println(numIslands(test))
	//for _, v := range test {
	//	fmt.Println(v, numIslands(test))
	//}
	//
	//scanDir()
	//a := []int{1}
	//fmt.Println(a[2])

}

func hanoi(n, from, to, aux int) {
	if n == 1 {
		//fmt.Printf("Move disk 1 from rod %d to rod %d\n", from, to)
		return
	}
	hanoi(n-1, from, aux, to)
	//fmt.Printf("Move disk %d from rod %d to rod %d\n", n, from, to)
	hanoi(n-1, aux, to, from)
}

func test_map() {
	var s map[string]int // 这样声明的map是 nil 不能直接赋值
	s = map[string]int{}
	s["xxx"] = 1
	fmt.Println(s)
}

func test_func() {
	s := [3]int{1, 2, 3}
	func(array [3]int) {
		array[0] = 7
		fmt.Println(array)
	}(s)
	fmt.Println(s)
}

func test_slice() {
	s := []int{1, 2, 3}
	func(s []int) {
		s = append(s, s[:]...)
		s[0] = 7
		fmt.Println(s)
	}(s)

	fmt.Println(s)
}

func appendSlice(slice []int) []int {
	// 向切片添加一个元素，注意这会改变切片的长度
	slice = append(slice, 4)
	return slice
}

func modifySliceElement(slice []int) {
	// 修改切片中的一个元素
	if len(slice) > 0 {
		slice[0] = 10
	}
}

func testEqual() bool {
	a := "cba"
	s := []byte(a)
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	fmt.Println(string(s))
	return true
}

func testDoublePtr(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	nums = filterRepeat(nums)
	if len(nums) == 0 {
		return 0
	}

	flag := 1
	max := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+1 == nums[i+1] {
			flag++
		} else {
			if flag > max {
				max = flag
			}
			flag = 1
		}
	}

	if flag > max {
		max = flag
	}

	return max
}

func filterRepeat(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}
	tmp := []int{a[0]}
	for i := 1; i < len(a); i++ {
		if a[i] != a[i-1] {
			tmp = append(tmp, a[i])
		}
	}
	return tmp
}

func DoneWithMap(nums []int) int {
	// 1. 先给她装到map里面
	mp := make(map[int]bool)
	for _, v := range nums {
		mp[v] = true
	}

	// 向上找
	current := 0
	for k := range mp {
		if !mp[k-1] {
			key := k
			count := 1
			for mp[key+1] {
				key++
				count++
			}

			if count > current {
				current = count
			}
		}

	}

	return current
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func singleNumber(nums []int) int {
	//mp := make(map[int]int)
	//for _, v := range nums {
	//	if _, ok := mp[v]; !ok {
	//		mp[v] = 1
	//	} else {
	//		delete(mp, v)
	//	}
	//}
	//fmt.Printf("%#v", mp)
	//os.Exit(1)
	//for k, v := range mp {
	//	if v == 1 {
	//		return k
	//	}
	//}
	//return 0
	// 神奇的位运算
	a := 0
	for _, num := range nums {
		a ^= num
	}

	return a
}

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func scanDir() {
	path := "F:\\cxxzoom\\everydayNormalGo2\\src"
	folder, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	count := 0
	for _, entry := range folder {
		if !entry.IsDir() {
			count++
		}
	}
	fmt.Println(folder)
	fmt.Println(count)
}

func testSlice(nums []int) []int {
	// 1. 首先想到的是双指针，分别指向开头和结尾
	// 2. 当头部指针有为0时，截取前面的部分，并把0加入最后
	// 3. 这里需要注意的是，两根指针的移动
	// 4. nums切片是 [:) 左闭右开
	//if len(nums) == 1 {
	//	return nums
	//}
	//
	//left, right := 0, len(nums)-1
	//// 两个指针相遇则结束了
	//for left < right {
	//	if nums[left] != 0 {
	//		left++
	//		continue
	//	}
	//
	//	nums = append(nums[:left], append(nums[left+1:], nums[left:left+1]...)...)
	//	if nums[right] != 0 {
	//		left++
	//	}
	//	right--
	//}

	slow, fast := 0, 0
	// 1. 快指针一直移动到第一个非0的地方
	// 2. 慢指针一直移动到以一个为0的地方
	// 3. 然后交换位置
	// 4. fast再移动，slow再移动

	// slow 一定会移动到最后一步吗
	for slow < len(nums) {
		if nums[fast] == 0 {
			slow = fast
			fast++
		}

		// 避免越界
		if fast >= len(nums) {
			return nums
		}

		// 同时满足时才进行交换
		if nums[fast] != 0 && nums[slow] == 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
	}

	return nums
}

func numIslands(grid [][]byte) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])

	gridNum := 0
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == '1' {
				gridNum++
				isGrid(&grid, r, c)
			}
		}
	}
	return gridNum
}

// 说的是把遍历了的为 1 的全部设置为 0
func isGrid(grid *[][]byte, r, c int) {
	row := len(*grid)      // 行
	col := len((*grid)[0]) // 列
	//log.Panic(row, col)
	(*grid)[r][c] = 0
	if r-1 >= 0 && (*grid)[r-1][c] == '1' {
		isGrid(grid, r-1, c)
	}

	if r+1 < row && (*grid)[r+1][c] == '1' {
		isGrid(grid, r+1, c)
	}

	if c-1 >= 0 && (*grid)[r][c-1] == '1' {
		isGrid(grid, r, c-1)
	}

	if c+1 < col && (*grid)[r][c+1] == '1' {
		isGrid(grid, r, c+1)
	}
}
