package model

type User struct {
	Id      int
	Name    string
	Surname string
}

func GetAllUsers() (users []User, err error) {
	users = []User{
		{1, "Jon", "Doa"},
		{2, "Goward", "Dowson"},
		{3, "Jack", "Rorck"},
		{4, "Diesel", "Meminger"},
		{5, "Jayn", "Air"},
		{6, "Martin", "Iden"},
		{7, "John", "Golt"},
		{8, "Samuel", "Tarly"},
		{9, "Germiona", "Greynger"},
	}
	return
}
