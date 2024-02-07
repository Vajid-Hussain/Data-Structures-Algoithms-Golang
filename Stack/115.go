type MinStack struct {
    arr []int
}


func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(val int)  {
    this.arr=append(this.arr, val)
}


func (this *MinStack) Pop()  {
   this.arr=this.arr[:len(this.arr)-1]
}


func (this *MinStack) Top() int {
    return this.arr[len(this.arr)-1]
}


func (this *MinStack) GetMin() int {
    min:=this.arr[0]
    for _,val:=range this.arr{
        if val<min{
            min=val
        }
    }
    return min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
