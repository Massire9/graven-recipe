package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"main/models"
	"strconv"
)

// RecipestepController operations for Recipestep
type RecipestepController struct {
	beego.Controller
}

// URLMapping ...
func (c *RecipestepController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Recipestep
// @Param	body		body 	models.Recipestep	true		"body for Recipestep content"
// @Success 201 {int} models.Recipestep
// @Failure 403 body is empty
// @router / [post]
func (c *RecipestepController) Post() {
	o := orm.NewOrm()
	step := models.RecipeStep{}
	recipe := models.Recipe{}

	requestBody := c.Ctx.Input.RequestBody
	err := json.Unmarshal(requestBody, &step)

	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	err = o.QueryTable("recipes").Filter("Id", id).One(&recipe)
	if err != nil {
		return
	}
	step.Recipe = &recipe
	insert, err := o.Insert(&step)
	if err != nil {
		return
	}
	c.Data["json"] = insert
	err = c.ServeJSON()
	if err != nil {
		return
	}
}

// GetOne ...
// @Title Get One
// @Description get Recipestep by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Recipestep
// @Failure 403 :id is empty
// @router /:id [get]

// Put ...
// @Title Put
// @Description update the Recipestep
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Recipestep	true		"body for Recipestep content"
// @Success 200 {object} models.Recipestep
// @Failure 403 :id is not int
// @router /:id [put]
func (c *RecipestepController) Put() {

	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Recipestep
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *RecipestepController) Delete() {

	c.ServeJSON()
}
