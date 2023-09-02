package nullable

import "database/sql"

func FromNullFloat64(n sql.NullFloat64) float64 {
	return n.Float64
}

func ToNullFloat64(n float64) sql.NullFloat64 {
	return sql.NullFloat64{
		Float64: n,
		Valid:   true,
	}
}
