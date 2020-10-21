package vm

import (
	"go-mega/model"
)

// EmailViewModel struct
type EditPostModel struct {
	Username string
	Token    string
	model.Post
}

// EmailViewModelOp struct
type EditPostModelOp struct{}

// GetVM func
func (EditPostModelOp) GetEditPostVM(post model.Post) EditPostModel {
	v := EditPostModel{}
	v.Post = post
	return v
}
