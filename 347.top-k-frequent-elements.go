/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	frequencyMap := make(map[int]int, len(nums))
	for _, v := range nums {
		frequencyMap[v]++
	}

	h:= &minHeap{}
	heap.Init(h)

	for key,value := range frequencyMap {
		heap.Push(h,pair{key:key,value:value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	results :=make([]int, 0,k)
	for h.Len() > 0 {
		results = append(results, heap.Pop(h).(pair).key)
	}

	return results
}

type pair struct {
	key   int
	value int
}

type minHeap []pair

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].value < h[j].value
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x any) {
	*h =append(*h,x.(pair))
	
}

func (h *minHeap) Pop() any {
	old := *h
	n:=len(old)
	p := old[n-1]
	old[n-1] = pair{}
	*h = old[:n-1]
	return p
}

// @lc c
// ode=end

