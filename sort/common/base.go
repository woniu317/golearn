package common

import (
	"fmt"
	"math/rand"
)

type Value []int

func NewValue(number int) Value {
	rand.Seed(1)
	v := make(Value, number)
	maxNumber := number
	for i := 0; i < number; i++ {
		v[i] = rand.Intn(maxNumber)
	}
	return v
}

func NewSameValue(number int) Value {
	rand.Seed(1)
	v := make(Value, number)
	maxNumber := number
	va := rand.Intn(maxNumber)
	for i := 0; i < number; i++ {
		v[i] = va
	}
	return v
}

func NewDifferentValue(number int) Value {
	rand.Seed(1)
	tmp := make(Value, number)
	for i := 0; i < number; i++ {
		tmp[i] = i
	}
	va := make(Value, number)
	for i := 0; i < number; i++ {
		index := rand.Intn(number - i)
		va[i] = tmp[index]
		tmp[index] = tmp[len(tmp)-i-1]
	}
	return va
}

func NewValueNil(number int) Value {
	return make(Value, number)[:0]
}

func (v Value) Print() {
	fmt.Println(v)
}

func (v Value) Len() int {
	return len(v)
}

func (v Value) Less(i, j int) bool {
	return v[i] < v[j]
}

func (v Value) Greater(i, j int) bool {
	return v[i] > v[j]
}

func (v Value) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func NewBucket(bucketSize int) map[int][]int {
	bucket := make(map[int][]int, bucketSize)
	return bucket
}

//notSafe
//return index
func (v Value) Pivot() int {
	middle := v.Len() / 2
	midIndex := mid(v[0], v[middle], v[v.Len()-1])
	if v[0] == midIndex {
		return 0
	}
	if v[middle] == midIndex {
		return middle
	}
	return v.Len() - 1
}

func mid(a, b, c int) int {
	if (a >= b && b >= c) || (c >= b && b >= a) {
		return b
	}
	if (b >= a && a >= c) || (c >= a && a >= b) {
		return a
	}
	return c
}
