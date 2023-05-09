package entity

type Member struct {
	Id    int
	Nama  string
	Point float32
}

type ListMember struct {
	Data Member
	Next *ListMember
}
