package nullable

import "database/sql"

func FromNullBool(n sql.NullBool) bool {
	return n.Bool
}

func ToNullBool(n bool) sql.NullBool {
	return sql.NullBool{
		Bool:  n,
		Valid: true,
	}
}
