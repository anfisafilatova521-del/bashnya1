package main
import(
"errors"
"fmt"
)

type Stack struct{
element []int
}

func(a *Stack) Push(new int){
a.element=append(a.element,new)
}

func (a*Stack) Pop() (int,error){
if  len(a.element) == 0{
return 0, errors.New("стек пуст")
}


i :=len(a.element)-1
del:=a.element[i]
a.element=a.element[:i]
return del,nil
}


func (a *Stack) IsEmpty() bool {
return len(a.element)==0
}

func (a *Stack) Size() int{
return len(a.element)
}

func (a *Stack) Clear(){
a.element=nil
}
func main() {
    var a Stack
    fmt.Println("Стек пуст?", a.IsEmpty()) 
    fmt.Println("Размер стека:", a.Size()) 
    a.Push(1)
    a.Push(2)
    fmt.Println("Размер стека после добавления двух элементов:", a.Size()) 

    a.Pop()
    fmt.Println("Размер стека после удаления последнего элемента:", a.Size()) 

    a.Clear()
    fmt.Println("Стек пуст после очищения?", a.IsEmpty()) 
}