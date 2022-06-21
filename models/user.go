package models

type UserPos struct {
	Name string
	CurrPos int
}

func NewUser(userName string) *UserPos{
	return &UserPos{
		Name: userName,
		CurrPos: 0,
	}
}