package nullable

import (
	"database/sql"
	"fmt"
	"time"
)

type Nullable interface {
	sql.NullBool |
		sql.NullByte |
		sql.NullTime |
		sql.NullString |
		sql.NullFloat64 |
		sql.NullInt16 |
		sql.NullInt32 |
		sql.NullInt64
}

type Convertable interface {
	bool | byte | float64 | int16 | int32 | int64 | string | time.Time
}

func From[T Nullable, V Convertable](n T) V {
	var converted V

	switch v := any(n).(type) {
	case sql.NullBool:
		if v.Valid {
			converted = any(v.Bool).(V)
		}
	case sql.NullByte:
		if v.Valid {
			converted = any(v.Byte).(V)
		}
	case sql.NullTime:
		if v.Valid {
			converted = any(v.Time).(V)
		}
	case sql.NullString:
		if v.Valid {
			converted = any(v.String).(V)
		}
	case sql.NullFloat64:
		if v.Valid {
			converted = any(v.Float64).(V)
		}
	case sql.NullInt16:
		if v.Valid {
			converted = any(v.Int16).(V)
		}
	case sql.NullInt32:
		if v.Valid {
			converted = any(v.Int32).(V)
		}
	case sql.NullInt64:
		if v.Valid {
			converted = any(v.Int64).(V)
		}
	default:
		panic(fmt.Sprintf("can not convert type %T to %T", n, v))
	}

	return converted
}

func To[T Convertable, V Nullable](n T) V {
	switch v := any(n).(type) {
	case bool:
		return any(sql.NullBool{Bool: v, Valid: true}).(V)
	case byte:
		return any(sql.NullByte{Byte: v, Valid: true}).(V)
	case time.Time:
		return any(sql.NullTime{Time: v, Valid: true}).(V)
	case string:
		return any(sql.NullString{String: v, Valid: true}).(V)
	case float64:
		return any(sql.NullFloat64{Float64: v, Valid: true}).(V)
	case int16:
		return any(sql.NullInt16{Int16: v, Valid: true}).(V)
	case int32:
		return any(sql.NullInt32{Int32: v, Valid: true}).(V)
	case int64:
		return any(sql.NullInt64{Int64: v, Valid: true}).(V)
	default:
		panic(fmt.Sprintf("can not convert type %T to %T", n, v))
	}
}
