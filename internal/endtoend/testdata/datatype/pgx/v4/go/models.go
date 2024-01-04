// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package datatype

import (
	"database/sql"
	"time"

	"github.com/jackc/pgtype"
)

type DtCharacter struct {
	A sql.NullString
	B sql.NullString
	C sql.NullString
	D sql.NullString
	E sql.NullString
}

type DtCharacterNotNull struct {
	A string
	B string
	C string
	D string
	E string
}

type DtDatetime struct {
	A sql.NullTime
	B sql.NullTime
	C sql.NullTime
	D sql.NullTime
	E sql.NullTime
	F sql.NullTime
	G sql.NullTime
	H sql.NullTime
}

type DtDatetimeNotNull struct {
	A time.Time
	B time.Time
	C time.Time
	D time.Time
	E time.Time
	F time.Time
	G time.Time
	H time.Time
}

type DtNetType struct {
	A pgtype.Inet
	B pgtype.CIDR
	C pgtype.Macaddr
}

type DtNetTypesNotNull struct {
	A pgtype.Inet
	B pgtype.CIDR
	C pgtype.Macaddr
}

type DtNumeric struct {
	A sql.NullInt16
	B sql.NullInt32
	C sql.NullInt64
	D pgtype.Numeric
	E pgtype.Numeric
	F sql.NullFloat64
	G sql.NullFloat64
	H sql.NullInt16
	I sql.NullInt32
	J sql.NullInt64
	K sql.NullInt16
	L sql.NullInt32
	M sql.NullInt64
}

type DtNumericNotNull struct {
	A int16
	B int32
	C int64
	D pgtype.Numeric
	E pgtype.Numeric
	F float32
	G float64
	H int16
	I int32
	J int64
	K int16
	L int32
	M int64
}

type DtRange struct {
	A pgtype.Int4range
	B pgtype.Int8range
	C pgtype.Numrange
	D pgtype.Tsrange
	E pgtype.Tstzrange
	F pgtype.Daterange
}

type DtRangeNotNull struct {
	A pgtype.Int4range
	B pgtype.Int8range
	C pgtype.Numrange
	D pgtype.Tsrange
	E pgtype.Tstzrange
	F pgtype.Daterange
}
