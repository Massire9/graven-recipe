package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"main/models"
	"strconv"
)

// RecipeController operations for RecipeController
type RecipeController struct {
	beego.Controller
}

// URLMapping ...
func (c *RecipeController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.Get)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create RecipeController
// @Param	body		body 	models.RecipeController	true		"body for RecipeController content"
// @Success 201 {object} models.RecipeController
// @Failure 403 body is empty
// @router /recipes [post]
func (c *RecipeController) Post() {
	var recipe = models.Recipe{}
	requestBody := c.Ctx.Input.RequestBody
	err := json.Unmarshal(requestBody, &recipe)
	err = c.ParseForm(&recipe)
	if err != nil {
		return
	}
	o := orm.NewOrm()
	_, err = o.Insert(&recipe)
	if err != nil {
		return
	}
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = recipe
	err = c.ServeJSON()
	if err != nil {
		return
	}
}

// GetOne GetOne ...
// @Title GetOne
// @Description get RecipeController by id
// @Param id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.RecipeController
// @Failure 403 :id is empty
// @router /recipes/:id [get]
func (c *RecipeController) GetOne() {
	strId := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}
	recipe := models.Recipe{Id: id}
	o := orm.NewOrm()
	err = o.QueryTable("recipes").RelatedSel("receipt_steps").One(&recipe)
	if err != nil {
		panic(err)
	}
	c.Data["json"] = recipe
	err = c.ServeJSON()
	if err != nil {
		return
	}
}

// Get ...
// @Title Get
// @Description get RecipeController
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.RecipeController
// @Failure 403
// @router /recipes [get]
func (c *RecipeController) Get() {
	var recipes []*models.Recipe
	o := orm.NewOrm()
	_, err := o.QueryTable("recipes").RelatedSel().All(&recipes)
	if err != nil {
		panic(err)
	}
	for _, recipe := range recipes {
		_, err := o.LoadRelated(recipe, "ReceiptSteps")
		if err != nil {
			panic(err)
		}
	}
	c.Data["json"] = recipes
	err = c.ServeJSON()
	if err != nil {
		return
	}
}

// Put ...
// @Title Put
// @Description update the RecipeController
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.RecipeController	true		"body for RecipeController content"
// @Success 200 {object} models.RecipeController
// @Failure 403 :id is not int
// @router /recipes/:id [put]
func (c *RecipeController) Put() {
	o := orm.NewOrm()
	recipe := models.Recipe{}
	recipe.Id, _ = strconv.Atoi(c.Ctx.Input.Param(":id"))

	requestBody := c.Ctx.Input.RequestBody
	err := json.Unmarshal(requestBody, &recipe)
	if err != nil {
		return
	}
	_, err = o.Update(&recipe, "content", "display_order")
	if err != nil {
		return
	}
	c.Data["json"] = recipe
	err = c.ServeJSON()
	if err != nil {
		return
	}
}

// Delete ...
// @Title Delete
// @Description delete the RecipeController
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /recipes/:id [delete]
func (c *RecipeController) Delete() {
	o := orm.NewOrm()
	_, err := o.QueryTable("recipes").Filter("id", c.Ctx.Input.Param(":id")).Delete()
	if err != nil {
		return
	}
	err = c.ServeJSON()
	if err != nil {
		return
	}
}
