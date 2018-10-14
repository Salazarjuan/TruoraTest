package recipes

import (
	"github.com/go-xorm/xorm"
)

func FindBy(db *xorm.Engine, limit int, offset int) (recipes Recipes, err error) {
	err = db.
		Limit(limit, offset).
		Find(&recipes)

	return
}
