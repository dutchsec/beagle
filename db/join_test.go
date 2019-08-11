package db

import (
	"testing"

	"go.dutchsec.com/beagle/db"
)

var (
	TestSetJoin = []Set{
		// {
		// 	Query: db.SelectQuery("TABLE1").Fields("DISTINCT *").Join("TABLE2").OnOLD(db.Field("TABLE1.TEST"), db.Field("TABLE2.TEST")).OrderBy(db.Field("TEST")),
		// 	Want:  db.Query("SELECT * FROM TABLE ORDER BY TEST ASC ;"),
		// },
		{
			Query: db.SelectQuery("TABLE1").Fields("DISTINCT *").Join("TABLE2").On(db.Equal(db.Field("TABLE1.TEST"), db.Field("TABLE2.TEST"))).OrderBy(db.Field("TEST")),
			Want:  db.Query("SELECT * FROM TABLE ORDER BY TEST ASC ;"),
		},
	}
)

func TestJoin(t *testing.T) {
	for _, ts := range TestSetJoin {
		got, _ := ts.Query.Build()

		if got != ts.Want {
			t.Errorf("Got: %s\nWant: %s", got, ts.Want)
		}
	}
}
