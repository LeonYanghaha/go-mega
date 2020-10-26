package model

import (
	"time"
)

// Post struct
type Post struct {
	ID        int `gorm:"primary_key"`
	UserID    int
	User      User
	Body      string     `gorm:"type:varchar(180)"`
	Timestamp *time.Time `sql:"DEFAULT:current_timestamp"`
}

func GetPostsByUserID(id int) (*[]Post, error) {
	var posts []Post
	if err := db.Preload("User").Where("user_id=?", id).Find(&posts).Error; err != nil {
		return nil, err
	}
	return &posts, nil
}

// GetPostsByUserIDPageAndLimit func
func GetPostsByUserIDPageAndLimit(id, page, limit int) (*[]Post, int, error) {
	var total int
	var posts []Post
	offset := (page - 1) * limit
	if err := db.Preload("User").Order("timestamp desc").Where("user_id=?", id).Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
		return nil, total, err
	}
	db.Model(&Post{}).Where("user_id=?", id).Count(&total)
	return &posts, total, nil
}

// GetPostsByPageAndLimit func
func GetPostsByPageAndLimit(page, limit int) (*[]Post, int, error) {
	var total int
	var posts []Post

	offset := (page - 1) * limit
	if err := db.Preload("User").Offset(offset).Limit(limit).Order("timestamp desc").Find(&posts).Error; err != nil {
		return nil, total, err
	}
	db.Model(&Post{}).Count(&total)

	return &posts, total, nil
}

// FormattedTimeAgo func
func (p *Post) FormattedTimeAgo() string {
	return FromTime(*p.Timestamp)
}

func GetPostByPostIdAndUserId(pid, uid int) (*Post, error) {
	var post Post
	if err := db.Where("id=? and user_id=?", pid, uid).Find(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func GetPostByPostId(pid int) (*Post, error) {
	var post Post
	if err := db.Where("id=? ", pid).Find(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func UpdatePostById(pId, uid int, pBody string) error {
	item, err := GetPostByPostIdAndUserId(pId, uid)
	if err != nil {
		return err
	}
	return db.Model(item).Update(map[string]interface{}{"body": pBody}).Error
}
