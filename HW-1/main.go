package main

import (
	"fmt"
	"sort"
)

func removeDuplicates(nums []int) []int {
	var size int = len(nums)
	var Map map[int]int = map[int]int{}

	for i := 0; i < size; i++ {
		Map[nums[i]]++

		if Map[nums[i]] > 1 {
			Map[nums[i]]--
			nums[i] = nums[size-1]
			size--
			i--
		}
	}

	// N - размер массива, K - количество уникальных элементов в массиве
	// Это in-place алгоритм за O(K) памяти и O(N) времени
	return nums[:size]
}

func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}

func fibonacci(N int) []int {
	var arr []int = make([]int, 0, N)
	arr = append(arr, 0)
	arr = append(arr, 1)

	for i := 2; i < N; i++ {
		arr = append(arr, arr[i-2]+arr[i-1])
	}

	return arr[:N]
}

func elementOccurrencesCount(nums []int, element int) int {
	var count int = 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == element {
			count++
		}
	}

	return count
}

func arraysIntersection(nums1, nums2 []int) []int {
	sort.Sort(sort.IntSlice(nums1))
	sort.Sort(sort.IntSlice(nums2))

	var result []int
	var index1, index2 int = 0, 0

	if len(nums1) < len(nums2) {
		result = make([]int, 0, len(nums1))

		for index1 == len(nums1) || index2 == len(nums2) {
			if nums1[index1] > nums2[index2] {
				index2++
			} else if nums1[index1] < nums2[index2] {
				index1++
			} else {
				result = append(result, nums1[index1])

				index1++
				index2++
			}
		}
	} else {
		result = make([]int, 0, len(nums2))

		for index1 == len(nums1) || index2 == len(nums2) {
			if nums2[index2] > nums1[index1] {
				index1++
			} else if nums2[index2] < nums1[index1] {
				index2++
			} else {
				result = append(result, nums2[index2])

				index1++
				index2++
			}
		}
	}

	return result
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var Map map[byte]int = map[byte]int{}

	for i := 0; i < len(s1); i++ {
		Map[s1[i]]++
	}

	for i := 0; i < len(s2); i++ {
		Map[s2[i]]--

		if Map[s2[i]] == 0 {
			delete(Map, s2[i])
		} else if Map[s2[i]] < 0 {
			return false
		}
	}

	return len(Map) == 0
}

func merge(nums1, nums2 []int) []int {
	var result []int = make([]int, 0, len(nums1)+len(nums2))
	var index1, index2 int = 0, 0

	for index1 != len(nums1) && index2 != len(nums2) {
		if nums1[index1] <= nums2[index2] {
			result = append(result, nums1[index1])
			index1++
		} else {
			result = append(result, nums2[index2])
			index2++
		}
	}

	for i := index1; i < len(nums1); i++ {
		result = append(result, nums1[i])
	}

	for i := index2; i < len(nums2); i++ {
		result = append(result, nums2[i])
	}

	return result
}

type Node struct {
	key   string
	value interface{}
	next  *Node
}

type HashTable struct {
	buckets []*Node
	size    int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		buckets: make([]*Node, size),
		size:    size,
	}
}

func (ht *HashTable) hash(key string) int {
	hash := 0

	for i := 0; i < len(key); i++ {
		hash = (19 * hash) + int(key[i])
	}

	return hash % ht.size
}

func (ht *HashTable) put(key string, value interface{}) {
	index := ht.hash(key)
	node := ht.buckets[index]

	if node == nil {
		ht.buckets[index] = &Node{key: key, value: value}
		return
	}

	for node != nil {
		if node.key == key {
			node.value = value
			return
		}

		if node.next == nil {
			node.next = &Node{key: key, value: value}
			return
		}

		node = node.next
	}
}

func (ht *HashTable) get(key string) (interface{}, bool) {
	index := ht.hash(key)
	node := ht.buckets[index]

	for node != nil {
		if node.key == key {
			return node.value, true
		}
		node = node.next
	}

	return "", false
}

func (ht *HashTable) delete(key string) bool {
	index := ht.hash(key)
	node := ht.buckets[index]

	if node == nil {
		return false
	}

	if node.key == key {
		ht.buckets[index] = node.next
		return true
	}

	prev := node
	for node != nil {
		if node.key == key {
			prev.next = node.next
			return true
		}

		prev = node
		node = node.next
	}
	return false
}

func binarySearch(nums []int, element int) int {
	var l, r int = 0, len(nums)
	var m int

	for l < r {
		m = (l + r) / 2

		if nums[m] == element {
			return m
		} else if nums[m] < element {
			l = m + 1
		} else {
			r = m
		}
	}

	if nums[m] == element {
		return m
	}

	return -1
}

type Stack struct {
	array []int
}

func (s *Stack) push(element int) {
	s.array = append(s.array, element)
}

func (s *Stack) pop() int {
	if len(s.array) == 0 {
		panic("bruh...")
	}

	popElement := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]

	return popElement
}

func (s *Stack) size() int {
	return len(s.array)
}

type Queue struct {
	leftStack  Stack
	rightStack Stack
}

func (q *Queue) push(element int) {
	q.rightStack.push(element)
}

func (q *Queue) pop() int {
	var popElement int

	if q.leftStack.size() == 0 {
		if q.rightStack.size() == 0 {
			panic("bruh...")
		}

		tempStack := Stack{}

		halfRightStackSize := q.rightStack.size() / 2

		for halfRightStackSize > 0 {
			tempStack.push(q.rightStack.pop())
			halfRightStackSize--
		}

		for q.rightStack.size() > 0 {
			q.leftStack.push(q.rightStack.pop())
		}

		popElement = q.leftStack.pop()

		for tempStack.size() > 0 {
			q.rightStack.push(tempStack.pop())
		}
	} else {
		popElement = q.leftStack.pop()
	}

	return popElement
}

func main() {

	fmt.Println(removeDuplicates([]int{1, 1, 2, 56, 1, 2, 3, 574, 536002, 1, 1, 1, 2, 2, 3, 4}))

	fmt.Println(bubbleSort([]int{1, 1, 2, 56, 1, 2, 3, 574, 536002, 1, 1, 1, 2, 2, 3, 4}))

	fmt.Println(fibonacci(10))

	fmt.Println(elementOccurrencesCount([]int{1, 1, 2, 56, 1, 2, 3, 574, 536002, 1, 1, 1, 2, 2, 3, 4}, 11))

	fmt.Println(arraysIntersection([]int{1, 2, 3, 4, 5}, []int{3, 4, 13, 546, 4564, 2}))

	fmt.Println(isAnagram("fatta", "afftt"))

	fmt.Println(merge([]int{1, 2, 3, 3, 3, 4, 4, 5}, []int{1, 2, 2, 2, 2, 4, 5, 6}), "\n")

	ht := NewHashTable(10)
	ht.put("name1", 1)
	ht.put("name2", "amogus")
	ht.put("name3", 3)

	if value, found := ht.get("name1"); found {
		fmt.Println("name1:", value)
	} else {
		fmt.Println("bruh...")
	}

	if status := ht.delete("name4"); status {
		fmt.Println("name4 is deleted")
	} else {
		fmt.Println("bruh...")
	}

	if status := ht.delete("name1"); status {
		fmt.Println("name1 is deleted")
	} else {
		fmt.Println("bruh...")
	}

	if value, found := ht.get("name2"); found {
		fmt.Println("name2:", value)
	} else {
		fmt.Println("bruh...")
	}

	if value, found := ht.get("name1"); found {
		fmt.Println("name1:", value)
	} else {
		fmt.Println("bruh...")
	}

	fmt.Println("")

	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 3), "\n")

	q := Queue{}

	q.push(1)
	q.push(2)
	q.push(3)
	q.push(3)
	q.push(3)

	fmt.Println(q.pop())
	fmt.Println(q.pop())
	fmt.Println(q.pop())
	fmt.Println(q.pop())

	q.push(1)
	q.push(2)

	fmt.Println(q.pop())
	fmt.Println(q.pop())
	fmt.Println(q.pop())

}
