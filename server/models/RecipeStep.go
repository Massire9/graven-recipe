package models

func (u *RecipeStep) TableName() string {
	return "recipe_steps"
}

type RecipeStep struct {
	Id           int
	Content      string
	DisplayOrder int
	Recipe       *Recipe `orm:"rel(fk)"`
}
