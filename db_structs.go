package db

import (
	"time"
)

type Migrations struct{
	Id int64	`db:"id" schema:"id"`
	Name string	`db:"name" schema:"name"`
	RunOn time.Time	`db:"run_on" schema:"run_on"`
}

type Users struct{
	Id int64	`db:"id" schema:"id"`
	First string	`db:"first" schema:"first"`
	Last string	`db:"last" schema:"last"`
	Emain string	`db:"emain" schema:"emain"`
	Password string	`db:"password" schema:"password"`
}