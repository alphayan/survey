package main

import (
	"github.com/alphayan/survey"
	"github.com/alphayan/survey/tests/util"
)

var value = ""

var table = []TestUtil.TestTableEntry{
	{
		"standard", &survey.Password{Message: "Please type your password:"}, &value,
	},
	{
		"please make sure paste works", &survey.Password{Message: "Please paste your password:"}, &value,
	},
	{
		"no help, send '?'", &survey.Password{Message: "Please type your password:"}, &value,
	},
}

func main() {
	TestUtil.RunTable(table)
}
