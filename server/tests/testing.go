package tests

import (
	"strings"
	"testing"

	"github.com/dxtym/anon/server/internal/store"
	"github.com/dxtym/anon/server/internal/utils"
)

func TestStore(t *testing.T, cfg utils.Config) (*store.Store, func(...string)) {
	t.Helper()
	ts := store.NewStore(cfg)
	if err := ts.Open(); err != nil {
		t.Fatal(err)
	}
	// truncate delete records, cascade - everything related to it
	return ts, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := ts.DB.Exec(
				"TRUNCATE TABLE %s CASCADE", 
				strings.Join(tables, ", "),
			); err != nil {
				t.Fatal(err)
			}
		}
		ts.Close()
	}
}