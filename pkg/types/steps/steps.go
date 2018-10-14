package steps

import (
	"Projects/TruoraTest-server/pkg/db"
)

type Steps []Step

type Step db.Step

func (u *Step) TableName() string {
	return "Step"
}
