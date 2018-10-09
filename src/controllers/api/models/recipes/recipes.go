package recipes

import (
	"Projects/TruoraTest/pkg/db"
)

type Recipes []Recipe

type Recipe db.Recipes

func (u *Recipe) TableName() string {
	return "recipes"
}
