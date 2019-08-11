package db

import (
	"testing"

	"go.dutchsec.com/beagle/db"
)

type ExprFunc func() (db.Query, []interface{})

type Set struct {
	Query db.Queryx
	Want  db.Query
}

var (
	TestSet = []Set{{
		Query: db.SelectQuery("TABLE").Fields("*").OrderBy(db.Field("TEST")),
		Want:  db.Query("SELECT * FROM TABLE ORDER BY TEST ASC ;"),
	},
	}
)

func TestOrder(t *testing.T) {
	for _, ts := range TestSet {
		got, _ := ts.Query.Build()

		if got != ts.Want {
			t.Errorf("Got: %s\nWant: %s", got, ts.Want)
		}
	}
}
