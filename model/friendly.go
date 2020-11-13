package model

import "time"

type Friend struct {
	Id         string    `json:"id" bson:"_id"`
	Name       string    `json:"name"`
	IsEnable   bool      `json:"is_enable"`
	Href       string    `json:"href"`
	Type       int       `json:"type"`
	Creatime   time.Time `json:"creatime"`
	Updatetime time.Time `json:"updatetime"`
}
