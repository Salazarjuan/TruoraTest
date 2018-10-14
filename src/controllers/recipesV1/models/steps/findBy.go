package steps

import (
	"github.com/go-xorm/xorm"
)

func FindBy(db *xorm.Engine, limit int, offset int) (steps Steps, err error) {
	err = db.
		Limit(limit, offset).
		Find(&steps)

	return
}
