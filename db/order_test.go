package db

import (
	"testing"
)

type ExprFunc func() (Query, []interface{})

type Set struct {
	Query Queryx
	Want  Query
}

var (
	TestSet = []Set{{
		Query: SelectQuery("TABLE").Fields("*").OrderBy(Field("TEST")),
		Want:  Query("SELECT * FROM TABLE ORDER BY TEST ASC ;"),
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
