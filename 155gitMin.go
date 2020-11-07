// 155. 最小栈
// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

// push(x) —— 将元素 x 推入栈中。
// pop() —— 删除栈顶的元素。
// top() —— 获取栈顶元素。
// getMin() —— 检索栈中的最小元素。

// 示例:

// 输入：
// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]

// 输出：
// [null,null,null,null,-3,null,0,-2]

// 解释：
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin();   --> 返回 -3.
// minStack.pop();
// minStack.top();      --> 返回 0.
// minStack.getMin();   --> 返回 -2.

// 提示：

// pop、top 和 getMin 操作总是在 非空栈 上调用。
package main

import "fmt"

func main() {
	// MinStack minStack := new MinStack();
	// minStack.push(-2);
	// minStack.push(0);
	// minStack.push(-3);
	// minStack.getMin();
	// minStack.pop();
	// minStack.top();
	// minStack.getMin();
	obj := Constructor()
	obj.Push(2147483646)
	obj.Push(2147483646)
	obj.Push(2147483647)
	obj.Top()
	obj.Pop()
	param_4 := obj.GetMin()
	fmt.Println(param_4)
	obj.Pop()
	param_4 = obj.GetMin()
	fmt.Println(param_4)
	obj.Pop()
	obj.Push(2147483647)
	obj.Top()
	param_4 = obj.GetMin()
	fmt.Println(param_4)

	// param_3 := obj.Top()
	// param_4 = obj.GetMin()

	// fmt.Println(param_3)
	// fmt.Println(param_4)
}

type MinStack struct {
	data    []int
	minData []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:    make([]int, 0),
		minData: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.minData) == 0 {
		this.minData = append(this.minData, x)
	} else if x < this.minData[len(this.minData)-1] {
		this.minData = append(this.minData, x)
	} else {
		this.minData = append(this.minData, this.minData[len(this.minData)-1])
	}
}

func (this *MinStack) Pop() {
	if len(this.data) == 0 {
		return
	}
	length := len(this.data)
	this.data = this.data[:length-1]
	this.minData = this.minData[:length-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.minData[len(this.minData)-1]
}

// ============================================
// // 执行用时：60 ms, 在所有 Go 提交中击败了8.18%的用户
// // 内存消耗：10 MB, 在所有 Go 提交中击败了6.07%的用户
// // MinStack int
// type MinStack struct {
// 	n []int
// }

// /** initialize your data structure here. */
// func Constructor() MinStack {
// 	return MinStack{}
// }

// func (this *MinStack) Push(x int) {
// 	this.n = append(this.n, x)
// }

// func (this *MinStack) Pop() {
// 	if len(this.n) == 1 {
// 		this.n = nil
// 	}
// 	if len(this.n) > 1 {

// 		this.n = this.n[:len(this.n)-1]
// 	}

// }

// func (this *MinStack) Top() int {

// 	res := this.n[len(this.n)-1]
// 	// this.Pop()
// 	return res
// }

// func (this *MinStack) GetMin() int {

// 	if this.n == nil {
// 		return 0
// 	}
// 	res := this.n[0]
// 	for _, v := range this.n {
// 		if res > v {
// 			res = v
// 		}
// 	}
// 	return res
// }
//====================================
/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
