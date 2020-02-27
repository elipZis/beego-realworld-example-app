package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Setup_20200227_115017 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Setup_20200227_115017{}
	m.Created = "20200227_115017"

	migration.Register("Setup_20200227_115017", m)
}

// Run the migrations
func (m *Setup_20200227_115017) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Setup_20200227_115017) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
