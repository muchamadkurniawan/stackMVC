package model

import (
	"stackMVC/database"
	"stackMVC/entity"
)

func Push(member entity.Member) {
	newGerbong := entity.ListMember{}
	newGerbong.Data = member
	temp := &database.DBMember
	if temp.Next == nil {
		temp.Next = &newGerbong
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &newGerbong
	}
}

func Pop() {

}

func ViewAllMember() []entity.Member {
	temp := database.DBMember.Next
	members := []entity.Member{}
	for temp != nil {
		members = append(members, temp.Data)
		temp = temp.Next
	}
	return members
}
