package main

import (
	"fmt"
	"roobo.com/golearn/sort/common"
	"sort"
	"time"
)

var (
	DataNumber         = 100
	BucketSize         = 10000
	UseSameNumber      = false
	UseRandNumber      = false
	UseDifferentNumber = true
	PrintResult        = true
	//PrintResult = false
)

func init() {
	if UseSameNumber {
		UseRandNumber = false
		UseDifferentNumber = false
		fmt.Printf("use %d same number\n", DataNumber)
	} else if UseRandNumber {
		UseDifferentNumber = false
		fmt.Printf("use %d rand number\n", DataNumber)
	} else {
		fmt.Printf("use %d different number\n", DataNumber)
	}
}

func main() {
	testSortProxy(bubbleSortASC, "bubbleSortASC")
	testSortProxy(bubbleSortAscImpl1, "bubbleSortAscImpl1")
	testSortProxy(bubbleSortAscImpl2, "bubbleSortAscImpl2")
	testSortProxy(insertSortAsc, "insertSortAsc")
	testSortProxy(selectSortASC, "selectSortASC")
	testSortProxy(mergeSortASC, "mergeSortASC")
	testSortProxy(qSortASC, "qSortASC")
	testSortProxy(bucketSortASC, "bucketSortASC")
	testSortProxy(goSortASC, "goSortASC")
	testSortProxy(goSortDESC, "goSortDESC")
}

func testSortProxy(sort func(value common.Value), sortFuncName string) {
	var data common.Value
	if UseSameNumber {
		data = common.NewSameValue(DataNumber)
	}
	if UseRandNumber {
		data = common.NewValue(DataNumber)
	}
	if UseDifferentNumber {
		data = common.NewDifferentValue(DataNumber)
	}
	//if PrintResult {
	//	data.Print()
	//}
	printRunTime(sort, data, sortFuncName)
	if PrintResult {
		data.Print()
	}
}

func printRunTime(sort func(common.Value), value common.Value, sortFuncName string) {
	startTime := time.Now().UnixNano()
	sort(value)
	endTime := time.Now().UnixNano()
	fmt.Println(sortFuncName, " need ", (endTime-startTime)/1e6, "ms")
}

//
func bubbleSortASC(value common.Value) {
	len := value.Len()
	for i := 0; i < len; i++ {
		for j := 0; j < len-i-1; j++ {
			if value.Less(j+1, j) {
				value.Swap(j, j+1)
			}
		}
	}
}

//若本轮无交换则排序结果
func bubbleSortAscImpl1(value common.Value) {
	len := value.Len()
	for i := 0; i < len; i++ {
		hasSwap := false
		for j := 0; j < len-i-1; j++ {
			if value.Less(j+1, j) {
				value.Swap(j, j+1)
				hasSwap = true
			}
		}
		if !hasSwap {
			return
		}
	}
}

//获取最后一次交换的位置，该位置即为下次停止的终点
func bubbleSortAscImpl2(value common.Value) {
	len := value.Len()
	lastSwapPosition := len
	for i := 0; i < len; i++ {
		end := min(len-i-1, lastSwapPosition)
		lastSwapPosition = 0
		for j := 0; j < end; j++ {
			if value.Less(j+1, j) {
				value.Swap(j, j+1)
				lastSwapPosition = j
			}
		}
		if lastSwapPosition == 0 {
			return
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//插入排序
func insertSortAsc(value common.Value) {
	len := value.Len()
	for i := 1; i < len; i++ {
		for j := i - 1; j >= 0; j-- {
			if value.Less(j+1, j) {
				value.Swap(j+1, j)
			} else {
				break
			}
		}
	}
}

//选择排序
func selectSortASC(value common.Value) {
	len := value.Len()
	for i := 0; i < len; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if value.Less(j, minIndex) {
				minIndex = j
			}
		}
		value.Swap(minIndex, i)
	}
}

//归并
func mergeSortASC(value common.Value) {
	len := value.Len()
	if len < 2 {
		return
	}
	mid := len / 2
	lValue := value[:mid]
	rValue := value[mid:]
	mergeSortASC(lValue)
	mergeSortASC(rValue)
	tempValue := common.NewValueNil(len)
	i, j := 0, 0
	for i < lValue.Len() && j < rValue.Len() {
		if lValue[i] <= rValue[j] {
			tempValue = append(tempValue, lValue[i])
			i++
		} else {
			tempValue = append(tempValue, rValue[j])
			j++
		}
	}
	if i < lValue.Len() {
		tempValue = append(tempValue, lValue[i:]...)
	}
	if j < rValue.Len() {
		tempValue = append(tempValue, rValue[j:]...)
	}
	for i = 0; i < value.Len(); i++ {
		value[i] = tempValue[i]
	}
}

//快排
func qSortASC(value common.Value) {
	len := value.Len()
	if len < 2 {
		return
	}
	pivotIndex := partition(value)
	qSortASC(value[:pivotIndex])
	qSortASC(value[pivotIndex+1:])
}

func partition(value common.Value) int {
	pivotIndex := value.Pivot()
	value.Swap(0, pivotIndex)
	i, j := 0, 1
	len := value.Len()
	for j < len {
		if value.Less(j, 0) {
			i++
			value.Swap(i, j)
		}
		j++
	}
	value.Swap(0, i)
	return i
}

//桶排序+快排
func bucketSortASC(value common.Value) {
	bucket := common.NewBucket(BucketSize)
	bucketSortASCCore(bucket, value, func(va int) int {
		return va / BucketSize
	})
	for _, v := range bucket {
		qSortASC(v)
	}
	value = value[:0]
	for i := 0; i < len(bucket); i++ {
		if va, ok := bucket[i]; ok {
			value = append(value, va...)
		}
	}
}

func bucketSortASCCore(bucket map[int][]int, value common.Value, hashFunc func(int2 int) int) {
	len := value.Len()
	for i := 0; i < len; i++ {
		code := hashFunc(value[i])
		bucket[code] = append(bucket[code], value[i])
	}
}

//利用go进行排序
func goSortASC(value common.Value) {
	sort.Sort(value)
}

//利用go倒排序
func goSortDESC(value common.Value) {
	sort.Sort(sort.Reverse(sort.IntSlice(value)))
}
