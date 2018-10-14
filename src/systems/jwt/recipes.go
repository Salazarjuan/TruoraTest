package jwt

import (
	Recipes "Projects/TruoraTest-server/pkg/types/recipes"
	ORM "Projects/TruoraTest-server/src/systems/db"

	"errors"
	"github.com/go-xorm/xorm"
)

func GetRecipeFromToken(db *xorm.Engine, tokenVal string) (recipe Recipes.Recipe, err error) {
	if tokenVal == "" {
		err = errors.New("No token present.")
		return
	}

	userId, err := IsTokenValid(tokenVal)
	if err != nil {
		err = errors.New("Token is invalid.")
		return
	}

	if userId < 1 {
		err = errors.New("Token missing required data.")
		return
	}

	recipe = Recipes.Recipe{Id: userId}
	err = ORM.FindBy(db, &recipe)
	if err != nil || recipe.Id < 1 {
		err = errors.New("User in token does not exist in system.")
		return
	}

	return
}
