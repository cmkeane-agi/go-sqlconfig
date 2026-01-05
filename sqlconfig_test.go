package sqlconfig

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadSQL(t *testing.T) {
	// create a temporary SQL file with two named queries
	dir := t.TempDir()
	p := filepath.Join(dir, "example.sql")
	content := `-- name: one
SELECT 1;

-- name: two
SELECT 2;`

	if err := os.WriteFile(p, []byte(content), 0644); err != nil {
		t.Fatalf("write temp sql: %v", err)
	}

	s, err := LoadSQL(p)
	if err != nil {
		t.Fatalf("LoadSQL error: %v", err)
	}

	if _, ok := s.Get("one"); !ok {
		t.Fatalf("missing query 'one'")
	}
	if _, ok := s.Get("two"); !ok {
		t.Fatalf("missing query 'two'")
	}
}
