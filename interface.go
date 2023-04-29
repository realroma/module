package main

import "fmt"

type Checker struct {
        items []interface{}
}

func (Checker *Checker) Add(items Checkable) { //Присваиваем метод Add к Checke$
        Checker.items = append(Checker.items, items) //Добавляем в слайс структ$
}

func (Checker *Checker) String() string {//Данный метод отвечает за возвращение$
        var str = ""
//      for i, value := range Checker.items  {//Присваиваем каждый элемент в пе$
//              fmt.Println(i, value)
//      }
        return fmt.Sprintf(str)
}


type Measurable interface {
        GetMetrics()
}

func GetMertics() string {
        return "nil"
}

type Checkable interface {
        Measurable //Встравание интерфейса и его методов.
        GetID()
}


func GetID() string {
	fmt.Println("GetID")
        return "nil"
}

func main() {
        items := &Checker{} //Присваивание переменной указателя на определённую$
        var Checkable1 Checkable
        var Checkable2 Checkable
        var Checkable3 Checkable
        items.Add(Checkable1)
        items.Add(Checkable2)
        items.Add(Checkable3)
        fmt.Println(items)
        fmt.Println(Checkable1)
}

