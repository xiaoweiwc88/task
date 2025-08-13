package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	// if (len(intervals) <= 1) {
	// 	return intervals
	// }
	// // 1)按起点排序
	// sort.Slice(intervals, func(i, j int) bool {
	// 	if intervals[i][0] == intervals[j][0] {
	// 		return intervals[i][1] < intervals[j][1]
	// 	}
	// 	return intervals[i][0] < intervals[j][0]
	// })
	//
	// //2) 扫描合并
	// res := make([][]int, 0, len(intervals))
	// for _, cur := range intervals {
	// 	//res为空或不重叠，直接放入
	// 	if len(res) == 0 || cur[0] > res[len(res) - 1][1] {
	// 		//复制一份，避免后续修改影响排序切片(可选)
	// 		res = append(res, []int{cur[0], cur[1]})
	// 	} else {
	// 		//重叠，扩展终点
	// 		if cur[1] > res[len(res) - 1][1] {
	// 			res[len(res) - 1][1] = cur[1]
	// 		}
	// 	}
	// }
	// return res
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 1. 按照区间的起始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果切片，放入第一个区间
	merged := [][]int{intervals[0]}

	// 2. 遍历排序后的区间数组
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		last := merged[len(merged)-1] // 取出结果切片中最后一个区间

		// 3. 比较当前区间与最后一个区间
		if current[0] <= last[1] { // 有重叠
			// 合并区间，结束位置取两者中较大的
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else { // 无重叠
			merged = append(merged, current)
		}
	}
	return merged
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // [[1 6] [8 10] [15 18]]
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))                    // [[1 5]]
}
