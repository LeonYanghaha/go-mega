package controller

import (
	"github.com/gorilla/mux"
	"go-mega/vm"
	"html/template"
	"log"
	"net/http"
)

type home struct{}

func (h home) registerRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/logout", middleAuth(logoutHandler))
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/register", registerHandler)
	r.HandleFunc("/user/{username}", middleAuth(profileHandler))
	r.HandleFunc("/profile_edit", middleAuth(profileEditHandler))
	r.HandleFunc("/follow/{username}", middleAuth(followHandler))
	r.HandleFunc("/unfollow/{username}", middleAuth(unFollowHandler))
	r.HandleFunc("/", middleAuth(indexHandler))
	r.HandleFunc("/explore", middleAuth(exploreHandler))
	r.HandleFunc("/reset_password_request", resetPasswordRequestHandler)
	r.HandleFunc("/reset_password/{token}", resetPasswordHandler)
	r.HandleFunc("/user/{username}/popup", popupHandler)

	r.HandleFunc("/post-delete", middleAuth(postHandler)).Methods("POST")
	r.HandleFunc("/post-edit", middleAuth(postEditHandler)).Methods("POST")
	r.HandleFunc("/post-update", middleAuth(postUpdateHandler)).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(notfoundHandler)
	r.HandleFunc("/404", notfoundHandler)

	http.Handle("/", r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "index.html"
	vop := vm.IndexViewModelOp{}
	page := getPage(r)
	username, _ := getSessionUser(r)
	if r.Method == http.MethodGet {
		flash := getFlash(w, r)
		v := vop.GetVM(username, flash, page, pageLimit)
		_ = templates[tpName].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		_ = r.ParseForm()
		body := r.Form.Get("body")
		errMessage := checkLen("Post", body, 1, 180)
		if errMessage != "" {
			setFlash(w, r, errMessage)
		} else {
			err := vm.CreatePost(username, body)
			if err != nil {
				log.Println("add Post error:", err)
				_, _ = w.Write([]byte("Error insert Post in database"))
				return
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	flash := getFlash(w, r)
	message := vm.NotFoundMessage{Flash: flash}
	tpl, _ := template.ParseFiles("templates/content/404.html")
	_ = tpl.Execute(w, &message)
}
