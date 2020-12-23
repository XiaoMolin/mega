package vm

import (
	"mega/model"
)

type IndexViewModel struct {
	BaseViewModel
	Posts []model.Post
	Flash string
}

// IndexViewModelOp struct
type IndexViewModelOp struct{}

// GetVM func
func (IndexViewModelOp) GetVM(username string,flash string) IndexViewModel {
	u1, _ := model.GetUserByUsername(username)
	posts, _ := model.GetPostsByUserID(u1.ID)
	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *posts,flash}
	v.SetCurrentUser(username)
	return v
}

func CreatePost(username,post string)error{
	u,_:=model.GetUserByUsername(username)
	return u.CreatePost(post)
}