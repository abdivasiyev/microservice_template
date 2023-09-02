package nullable

import "database/sql"

func FromNullString(s sql.NullString) string {
	return s.String
}

func ToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}
