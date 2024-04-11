package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateRecipeTable_20240411_121429 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateRecipeTable_20240411_121429{}
	m.Created = "20240411_121429"

	migration.Register("CreateRecipeTable_20240411_121429", m)
}

// Run the migrations
func (m *CreateRecipeTable_20240411_121429) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE recipes (id serial PRIMARY KEY, title varchar(255), description text, cooking_time integer)")
}

// Reverse the migrations
func (m *CreateRecipeTable_20240411_121429) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE recipes")
}
