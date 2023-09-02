package nullable

import "database/sql"

func FromNullByte(s sql.NullByte) byte {
	return s.Byte
}

func ToNullByte(n byte) sql.NullByte {
	return sql.NullByte{
		Byte:  n,
		Valid: n > 0,
	}
}
