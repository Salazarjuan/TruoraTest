package recipes

import "Projects/TruoraTest-server/pkg/db"

type Recipes []Recipe

type Recipe db.Recipe

func (u *Recipe) TableName() string {
	return "recipe"
}