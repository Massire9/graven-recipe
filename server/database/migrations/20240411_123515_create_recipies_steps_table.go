package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateRecipiesStepsTable_20240411_123515 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateRecipiesStepsTable_20240411_123515{}
	m.Created = "20240411_123515"

	migration.Register("CreateRecipiesStepsTable_20240411_123515", m)
}

// Run the migrations
func (m *CreateRecipiesStepsTable_20240411_123515) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE recipe_steps (id serial primary key, content text, display_order int, recipe_id integer references recipes(id))")

}

// Reverse the migrations
func (m *CreateRecipiesStepsTable_20240411_123515) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE recipe_steps")
}
