package store

import (
	"strings"
	"testing"

	"github.com/dxtym/anon/server/internal/utils"
)

func TestStore(t *testing.T, cfg utils.Config) (*Store, func(...string)) {
	t.Helper()
	ts := NewStore(cfg)
	if err := ts.Open(); err != nil {
		t.Fatal(err)
	}

	return ts, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := ts.db.Exec(
				"TRUNCATE TABLE %s CASCADE",
				strings.Join(tables, ", "),
			); err != nil {
				t.Fatal(err)
			}
		}
		ts.Close()
	}
}
