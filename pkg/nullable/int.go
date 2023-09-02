package nullable

import "database/sql"

func FromNullInt16(n sql.NullInt16) int16 {
	return n.Int16
}

func ToNullInt16(n int16) sql.NullInt16 {
	return sql.NullInt16{
		Int16: n,
		Valid: true,
	}
}

func FromNullInt32(n sql.NullInt32) int32 {
	return n.Int32
}

func ToNullInt32(n int32) sql.NullInt32 {
	return sql.NullInt32{
		Int32: n,
		Valid: true,
	}
}

func FromNullInt64(n sql.NullInt64) int64 {
	return n.Int64
}

func ToNullInt64(n int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: n,
		Valid: true,
	}
}
