package db

import (
	"testing"

	"go.dutchsec.com/beagle/db"
)

var (
	TestSetWhere = []Set{
		{
			Query: db.SelectQuery("TABLE").Fields("*").Where(db.Equal(db.Field("TABLE1.TEST"), "test")),
			Want:  db.Query("SELECT * FROM TABLE WHERE TABLE1.TEST = ? ;"),
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
