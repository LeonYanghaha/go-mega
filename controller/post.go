package controller

import (
	"fmt"
	"go-mega/model"
	"go-mega/vm"
	"net/http"
	"strconv"
)

func postUpdateHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	username, _ := getSessionUser(r)
	currentUser, err := model.GetUserByUsername(username)
	if err != nil {
		fmt.Print(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	postId, _ := strconv.Atoi(r.Form.Get("post_id"))
	postBody := r.Form.Get("post_body")
	_ = model.UpdatePostById(postId, currentUser.ID, postBody)
	http.Redirect(w, r, "/user/"+username, http.StatusSeeOther)

}

func postEditHandler(w http.ResponseWriter, r *http.Request) {

	tpName := "edit_post.html"
	vop := vm.EditPostModelOp{}
	_ = r.ParseForm()
	username, _ := getSessionUser(r)
	postId, _ := strconv.Atoi(r.Form.Get("post_id"))
	currentUser, err := model.GetUserByUsername(username)
	if err != nil {
		fmt.Print(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	post, errs := model.GetPostByPostIdAndUserId(postId, currentUser.ID)
	if err != nil {
		fmt.Print(errs)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	v := vop.GetEditPostVM(*post)
	_ = templates[tpName].Execute(w, &v)

}

func postHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		_ = r.ParseForm()
		username, _ := getSessionUser(r)
		postId, _ := strconv.Atoi(r.Form.Get("post_id"))
		_ = vm.DeletePost(username, postId)
		http.Redirect(w, r, "/user/"+username, http.StatusSeeOther)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func exploreHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "explore.html"
	vop := vm.ExploreViewModelOp{}
	username, _ := getSessionUser(r)
	page := getPage(r)
	v := vop.GetVM(username, page, pageLimit)
	_ = templates[tpName].Execute(w, &v)
}
