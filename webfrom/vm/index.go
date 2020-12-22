package vm

import (
	"mega/database/model"
)

type IndexViewModelOp struct {
}

type IndexViewModel struct {
	BaseViewModel
	model.User
	Posts []model.Post
}

func (IndexViewModelOp) GetVM() IndexViewModel {
	u1 := model.User{Username: "bonfy"}
	u2 := model.User{Username: "rene"}

	posts := []model.Post{
		model.Post{User: u1, Body: "Beautiful day in Portland!"},
		model.Post{User: u2, Body: "The Avengers movie was so cool!"},
	}

	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, u1, posts}
	return v
}
