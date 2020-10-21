package controller

import (
	"go-mega/model"
	"go-mega/vm"
	"net/http"
	"strconv"
)

func postEditHandler(w http.ResponseWriter, r *http.Request) {

	tpName := "edit_post.html"
	vop := vm.EditPostModelOp{}
	println("......")
	//_ = r.ParseForm()
	//username, _ := getSessionUser(r)
	//postId, _ := strconv.Atoi(r.Form.Get("post_id"))
	//currentUser, err := model.GetUserByUsername(username)
	//if err != nil {
	//	fmt.Print(err)
	//	http.Redirect(w, r, "/", http.StatusSeeOther)
	//}
	//post, errs := model.GetPostByPostIdAndUserId(postId, currentUser.ID)
	//if err != nil {
	//	fmt.Print(errs)
	//	http.Redirect(w, r, "/", http.StatusSeeOther)
	//}
	//fmt.Println(post)
	v := vop.GetEditPostVM(model.Post{
		ID:        1,
		UserID:    1,
		User:      model.User{ID:2,Username:"sdfsd"},
		Body:      "vsdkjfhsdfhsiod",
		Timestamp: nil,
	})

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
