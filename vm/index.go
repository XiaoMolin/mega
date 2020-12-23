package vm

import (
	"mega/model"
)

type IndexViewModel struct {
	BaseViewModel
	Posts []model.Post
	Flash string

	BasePageViewModel
}

// IndexViewModelOp struct
type IndexViewModelOp struct{}

// GetVM func
func (IndexViewModelOp) GetVM(username string,flash string,page,limit int) IndexViewModel {
	u, _ := model.GetUserByUsername(username)
	posts, total, _ := u.FollowingPostsByPageAndLimit(page, limit)
	v := IndexViewModel{}
	v.SetTitle("Homepage")
	v.Posts = *posts
	v.Flash = flash
	v.SetBasePageViewModel(total, page, limit)
	v.SetCurrentUser(username)
	return v
}

func CreatePost(username,post string)error{
	u,_:=model.GetUserByUsername(username)
	return u.CreatePost(post)
}