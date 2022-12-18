package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type PickUpTable_20221218_080855 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &PickUpTable_20221218_080855{}
	m.Created = "20221218_080855"

	migration.Register("PickUpTable_20221218_080855", m)
}

// Run the migrations
func (m *PickUpTable_20221218_080855) Up() {
	m.SQL(`CREATE TABLE pick_up_history (
		id serial PRIMARY KEY,
		cover_id int,
		pick_up_at timestamp DEFAULT NULL,
		return_at timestamp  DEFAULT NULL,
		created_at timestamp NOT NULL DEFAULT current_timestamp,
		updated_at timestamp NOT NULL DEFAULT current_timestamp);
	`)
}

// Reverse the migrations
func (m *PickUpTable_20221218_080855) Down() {
	m.SQL("DROP TABLE pick_up_history;")
}
