package db

import (
	"testing"
)

var (
	TestSetWhere = []Set{
		{
			Query: SelectQuery("TABLE").Fields("*").Where(Equal(Field("TABLE1.TEST"), "test")),
			Want:  Query("SELECT * FROM TABLE WHERE TABLE1.TEST = ? ;"),
		},
	}
)

func TestWhere(t *testing.T) {
	for _, ts := range TestSetWhere {
		got, _ := ts.Query.Build()

		if got != ts.Want {
			t.Errorf("Got: %s\nWant: %s", got, ts.Want)
		}
	}
}
