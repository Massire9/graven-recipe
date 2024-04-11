package models

import _ "github.com/beego/beego/v2/client/orm"

func (u *Recipe) TableName() string {
	return "recipes"
}

type Recipe struct {
	Id           int
	Title        string
	Description  string
	CookingTime  int
	ReceiptSteps []*RecipeStep `orm:"reverse(many)"`
}

type APIRecipe struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CookingTime int    `json:"cooking_time"`
}
