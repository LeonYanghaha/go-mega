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

	model.AddUser("bonfy", "abc123", "i@bonfy.im")
	model.AddUser("rene", "abc123", "rene@test.com")
	model.AddUser("yanghaha", "yanghaha", "yanghaha@test.com")

	u1, _ := model.GetUserByUsername("bonfy")
	u1.CreatePost("Beautiful day in Portland!")
	model.UpdateAboutMe(u1.Username, `I'm the author of Go-Mega Tutorial you are reading now!`)

	u2, _ := model.GetUserByUsername("rene")
	u2.CreatePost("The Avengers movie was so cool!")
	u2.CreatePost("Sun shine is beautiful")

	u3, _ := model.GetUserByUsername("yanghaha")
	u3.CreatePost("yanghahaaaa")
	u3.CreatePost(".............")

	u1.Follow(u2.Username)
	u2.Follow(u3.Username)
	u3.Follow(u2.Username)
}
