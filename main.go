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
	model.Push(member1)
	model.Push(member2)
	allMember := model.ViewAllMember()
	fmt.Println(allMember)
}
