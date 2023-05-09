package main

import (
	"fmt"
	"stackMVC/entity"
	"stackMVC/model"
)

func main() {
	member1 := entity.Member{
		Id:    1,
		Nama:  "A",
		Point: 31,
	}
	member2 := entity.Member{
		Id:    2,
		Nama:  "B",
		Point: 32,
	}

	member3 := entity.Member{
		Id:    3,
		Nama:  "C",
		Point: 33,
	}
	model.Push(member1)
	model.Push(member2)
	model.Push(member3)
	allMember := model.ViewAllMember()
	fmt.Println(allMember)
	model.Pop()
	model.Pop()
	model.Pop()
	fmt.Println("setelah POP()")
	fmt.Println(model.ViewAllMember())
}
