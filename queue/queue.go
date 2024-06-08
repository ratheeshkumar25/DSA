package main

import "fmt"

type Items struct {
  val interface {}
  Next *Items
}
type Queue struct{
  Front *Items
  Rear *Items
  length int
}

func (q*Queue)len()int{
  return q.length
}
func (q*Queue)isEmpty()bool{
  if q.length == 0{
    return true
  }
  return false
}

func (q *Queue)Enqueue(val interface{}){
  fmt.Println(val)
  newNode := &Items{val:val}
  if q.isEmpty(){
    q.Front = newNode
    q.Rear = newNode
  }else{
    q.Rear.Next = newNode
    q.Rear = newNode
  }
  q.length++
}

func (q *Queue)Dequeue()interface{}{
  if q.isEmpty(){
    return nil
  }
  val := q.Front.val
  q.Front = q.Front.Next
  q.length--
  return val
}

func (q *Queue)Display(){
   current := q.Front
    for current != nil {
        fmt.Printf("%d ", current.val)
        current = current.Next
    }
    fmt.Println()

}
func main(){
  queue := &Queue{}
  queue.Enqueue(20)
  queue.Enqueue(30)
  queue.Enqueue(40)
  fmt.Println("Queue is ")
  queue.Display()
  queue.Dequeue()
  fmt.Println("After removing the first element")
  queue.Display()
}
