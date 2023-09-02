package nullable

import (
	"database/sql"
	"testing"
)

func TestFrom(t *testing.T) {
	gotStr := From[sql.NullString, string](sql.NullString{String: "Hello, World!", Valid: true})
	if gotStr != "Hello, World!" {
		t.Fatalf("expected: %v, gotStr: %v", "Hello, World!", gotStr)
	}

	gotStr = From[sql.NullString, string](sql.NullString{String: "Hello, World!", Valid: false})
	if gotStr != "" {
		t.Fatalf("expected: %v, gotStr: %v", "Hello, World!", gotStr)
	}
}

func TestTo(t *testing.T) {
	gotStr := To[string, sql.NullString]("Hello, World!")
	if !gotStr.Valid {
		t.Fatalf("expected: %v, gotStr: %v", true, gotStr.Valid)
	}

	if gotStr.String != "Hello, World!" {
		t.Fatalf("expected: %v, gotStr: %v", "Hello, World!", gotStr)
	}
}
