package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-mega/model"
	"log"
)

func main() {
	log.Println("DB Init ...")
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	db.DropTableIfExists(model.User{}, model.Post{}, "follower")
	db.CreateTable(model.User{}, model.Post{})

	_ = model.AddUser("bonfy", "abc123", "i@bonfy.im")
	_ = model.AddUser("rene", "abc123", "rene@test.com")
	_ = model.AddUser("yanghaha", "yanghaha", "yanghaha@test.com")

	u1, _ := model.GetUserByUsername("bonfy")
	_ = u1.CreatePost("Beautiful day in Portland!")
	_ = model.UpdateAboutMe(u1.Username, `I'm the author of Go-Mega Tutorial you are reading now!`)

	u2, _ := model.GetUserByUsername("rene")
	_ = u2.CreatePost("The Avengers movie was so cool!")
	_ = u2.CreatePost("Sun shine is beautiful")

	u3, _ := model.GetUserByUsername("yanghaha")
	_ = u3.CreatePost("yanghahaaaa")
	_ = u3.CreatePost(".............")

	_ = u1.Follow(u2.Username)
	_ = u2.Follow(u3.Username)
	_ = u3.Follow(u2.Username)
}
