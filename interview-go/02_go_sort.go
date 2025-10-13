package main

import (
	"fmt"
)

// 冒泡排序
// 原理：每次遍历数组，把相邻的元素两两比较，把大的“冒泡”到末尾
func bubbleSort(arr []int) {
	m := 0
	n := len(arr)
	end := n - 1 // 当前轮遍历的终点
	for end > 0 {
		lastSwap := 0 // 记录本轮最后一次交换的位置
		for j := 0; j < end; j++ {

			if arr[j] > arr[j+1] {
				m = m + 1
				arr[j], arr[j+1] = arr[j+1], arr[j]
				lastSwap = j // 更新最后交换位置
			}
		}
		end = lastSwap // 下一轮只遍历到最后交换位置
	}
	println("排序次数", m)
}

// 选择排序
// 原理：每次从未排序部分选择最小（或最大）的元素，放到已排序部分末尾
func selectionSort(arr []int) {
	n := len(arr)
	m := 0
	for i := 0; i < n-1; i++ {
		minIndex := i
		// 找到剩余部分的最小值
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
				m = m + 1
			}
		}
		// 交换当前元素和最小值
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	println("排序次数", m)
}

func main() {
	arr1 := []int{5, 2, 9, 1, 5, 6, 44, 4, 21, 49}
	arr2 := make([]int, len(arr1))
	copy(arr2, arr1)

	fmt.Println("原始数组:", arr1)

	bubbleSort(arr1)
	fmt.Println("冒泡排序结果:", arr1)

	selectionSort(arr2)
	fmt.Println("选择排序结果:", arr2)
}
