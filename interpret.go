package main

import "fmt"

func solution(val int, cmds []string, nums []int) int {
	if len(cmds) != len(nums) {
		return -1
	}

	ans := val
	for i, c := range cmds {
		switch c {
		case "*":
			ans *= nums[i]
		case "+":
			ans += nums[i]
		case "-":
			ans -= nums[i]
		}
	}

	return ans
}

func main() {
	ans := solution(4, []string{"+", "-", "*"}, []int{1, 2, 3})
	if ans != 9 {
		fmt.Printf("FAIL: expected 9, got %d\n", ans)
	} else {
		fmt.Printf("PASS: got %d\n", ans)
	}
}
