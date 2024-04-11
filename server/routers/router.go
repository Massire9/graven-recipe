package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"main/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//Recipe
	beego.CtrlGet("/recipes", (*controllers.RecipeController).Get)
	beego.CtrlGet("/recipes/:id", (*controllers.RecipeController).GetOne)
	beego.CtrlPut("/recipes/:id", (*controllers.RecipeController).Put)
	beego.CtrlDelete("/recipes/:id", (*controllers.RecipeController).Delete)
	beego.CtrlPost("/recipes", (*controllers.RecipeController).Post)

	//RecipeStep
	beego.CtrlPost("/recipes/:id/recipe-steps", (*controllers.RecipestepController).Post)
	beego.CtrlPut("/recipe-steps/:id", (*controllers.RecipestepController).Put)
	beego.CtrlDelete("/recipes/recipe-steps/:id", (*controllers.RecipestepController).Delete)
}
