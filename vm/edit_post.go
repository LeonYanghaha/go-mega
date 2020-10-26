package vm

import (
	"go-mega/model"
)

type EditPostModel struct {
	BaseViewModel
	Username string
	Token    string
	model.Post
	BasePageViewModel
}

type EditPostModelOp struct{}

func (EditPostModelOp) GetEditPostVM(post model.Post) EditPostModel {
	v := EditPostModel{}
	v.Post = post
	return v
}
