package nullable

import (
	"database/sql"
	"time"
)

func FromNullTime(n sql.NullTime) time.Time {
	return n.Time
}

func ToNullTime(n time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  n,
		Valid: true,
	}
}
