package main

import (
	"container/heap"
	"fmt"
)

func main() {
	// ["StockPrice","update","update","current","maximum","update","maximum","update","minimum"]
	// [[],[1,10],[2,5],[],[],[1,3],[],[4,2],[]]
	// s := StockPrice{hp{}, hp{}, make(map[int]int), 0}
	s := Constructor()
	s.Update(1, 10)
	s.Update(2, 5)
	fmt.Println(s.Current())
	fmt.Println(s.Maximum())
	s.Update(1, 3)
	fmt.Println(s.Maximum())
	s.Update(4, 2)
	// fmt.Println(s.Minimum())
	s.Update(5, 11)
	fmt.Println(s.Minimum())

}

//other

// type StockPrice struct {
//     prices       *redblacktree.Tree
//     timePriceMap map[int]int
//     maxTimestamp int
// }

// func Constructor() StockPrice {
//     return StockPrice{redblacktree.NewWithIntComparator(), map[int]int{}, 0}
// }

// func (sp *StockPrice) Update(timestamp, price int) {
//     if prevPrice := sp.timePriceMap[timestamp]; prevPrice > 0 {
//         if times, _ := sp.prices.Get(prevPrice); times.(int) > 1 {
//             sp.prices.Put(prevPrice, times.(int)-1)
//         } else {
//             sp.prices.Remove(prevPrice)
//         }
//     }
//     times := 0
//     if val, ok := sp.prices.Get(price); ok {
//         times = val.(int)
//     }
//     sp.prices.Put(price, times+1)
//     sp.timePriceMap[timestamp] = price
//     if timestamp >= sp.maxTimestamp {
//         sp.maxTimestamp = timestamp
//     }
// }

// func (sp *StockPrice) Current() int { return sp.timePriceMap[sp.maxTimestamp] }
// func (sp *StockPrice) Maximum() int { return sp.prices.Right().Key.(int) }
// func (sp *StockPrice) Minimum() int { return sp.prices.Left().Key.(int) }

//=========================================
//other
type StockPrice struct {
	maxPrice, minPrice hp
	timePriceMap       map[int]int
	maxTimestamp       int
}

func Constructor() StockPrice {
	return StockPrice{timePriceMap: map[int]int{}}
}

func (sp *StockPrice) Update(timestamp, price int) {
	heap.Push(&sp.maxPrice, pair{-price, timestamp})
	heap.Push(&sp.minPrice, pair{price, timestamp})
	sp.timePriceMap[timestamp] = price
	if timestamp > sp.maxTimestamp {
		sp.maxTimestamp = timestamp
	}
}

func (sp *StockPrice) Current() int {
	return sp.timePriceMap[sp.maxTimestamp]
}

func (sp *StockPrice) Maximum() int {
	for {
		if p := sp.maxPrice[0]; -p.price == sp.timePriceMap[p.timestamp] {
			return -p.price
		}
		heap.Pop(&sp.maxPrice)
	}
}

func (sp *StockPrice) Minimum() int {
	for {
		if p := sp.minPrice[0]; p.price == sp.timePriceMap[p.timestamp] {
			return p.price
		}
		heap.Pop(&sp.minPrice)
	}
}

type pair struct{ price, timestamp int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].price < h[j].price }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
