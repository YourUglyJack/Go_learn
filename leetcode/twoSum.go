package main

/*
解答成功:

	执行耗时:4 ms,击败了95.17% 的Go用户
	内存消耗:4.1 MB,击败了42.97% 的Go用户
*/
func twoSum(nums []int, target int) []int {
	// hash, k is num and value is index
	//m := map[int]int{}
	m := make(map[int]int)
	for i, num := range nums {
		// 如果作差在m里，就说明找到了
		if j, flag := m[target-num]; flag {
			return []int{i, j}
		} else {
			// 如果没有就放进m里
			m[num] = i
		}
	}
	return nil
}

func main() {

}
