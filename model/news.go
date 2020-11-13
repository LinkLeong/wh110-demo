package model

import (
	"time"
)

type News struct {
	Id         string    `json:"id" bson:"_id"`
	Type       int       `json:"type"`
	Title      string    `json:"title"`
	Summary    string    `json:"summary"`
	Image      string    `json:"image"`
	URL        string    `json:"url"`
	Content    string    `json:"content"`
	Showtime   time.Time `json:"showtime"`
	Read       int       `json:"read"`
	Recommend  bool      `json:"recommend"`
	Top        bool      `json:"top"`
	State      int       `json:"state"`
	Order      string    `json:"order"`
	Migrate    int       `json:"migrate"`
	RealRead   int       `json:"real_read"`
	Creatime   time.Time `json:"creatime"`
	Updatetime time.Time `json:"updatetime"`
}

//func (user *User) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("IdStr",com.ToStr( uuid.NewV4()))
//	return nil
//}
