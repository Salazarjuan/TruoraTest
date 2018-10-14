package db

import (
	"time"
)

type Migrations struct{
	Id int64	`xorm:"'id' pk autoincr" json:"id" schema:"id"`
	Name string	`xorm:"name" json:"name" schema:"name"`
	RunOn time.Time	`xorm:"run_on" json:"run_on" schema:"run_on"`
}

func (t Migrations) TableName() string {
	 return "migrations"
}

func (t Migrations) SetId(id int64) {
	t.Id = id
}

func (t Migrations) GetId() int64 {
	return t.Id
}

type Recipe struct{
	Id int64	`xorm:"'id' pk autoincr" json:"id" schema:"id"`
	Name string	`xorm:"name" json:"name" schema:"name"`
	Description string	`xorm:"description" json:"description" schema:"description"`
	Oven string	`xorm:"oven" json:"oven" schema:"oven"`
	Time int64	`xorm:"time" json:"time" schema:"time"`
	NoPersons int64	`xorm:"noPersons" json:"noPersons" schema:"noPersons"`
	CreatorID int64	`xorm:"creatorID" json:"creatorID" schema:"creatorID"`
}

func (t Recipe) TableName() string {
	 return "recipe"
}

func (t Recipe) SetId(id int64) {
	t.Id = id
}

func (t Recipe) GetId() int64 {
	return t.Id
}

type Step struct{
	Id int64	`xorm:"'id' pk autoincr" json:"id" schema:"id"`
	Step string	`xorm:"step" json:"step" schema:"step"`
	RecipeID int64	`xorm:"recipeID" json:"recipeID" schema:"recipeID"`
}

func (t Step) TableName() string {
	 return "steps"
}

func (t Step) SetId(id int64) {
	t.Id = id
}

func (t Step) GetId() int64 {
	return t.Id
}

type Users struct{
	Id int64	`xorm:"'id' pk autoincr" json:"id" schema:"id"`
	First string	`xorm:"first" json:"first" schema:"first"`
	Last string	`xorm:"last" json:"last" schema:"last"`
	Email string	`xorm:"email" json:"email" schema:"email"`
	Password string	`xorm:"password" json:"password" schema:"password"`
}