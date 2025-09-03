package pgxutil

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// Text ---------- Text / String ----------
func Text(s string) pgtype.Text {
	if s == "" {
		return pgtype.Text{}
	}
	return pgtype.Text{String: s, Valid: true}
}

// Bool ---------- Bool ----------
func Bool(b bool) pgtype.Bool {
	return pgtype.Bool{Bool: b, Valid: true}
}

// Int2 ---------- Integer types ----------
func Int2(i int16) pgtype.Int2 {
	return pgtype.Int2{Int16: i, Valid: true}
}
func Int2IfPositive(i int16) pgtype.Int2 {
	if i > 0 {
		return pgtype.Int2{Int16: i, Valid: true}
	}
	return pgtype.Int2{}
}

func Int4(i int32) pgtype.Int4 {
	return pgtype.Int4{Int32: i, Valid: true}
}
func Int4IfPositive(i int32) pgtype.Int4 {
	if i > 0 {
		return pgtype.Int4{Int32: i, Valid: true}
	}
	return pgtype.Int4{}
}

func Int8(i int64) pgtype.Int8 {
	return pgtype.Int8{Int64: i, Valid: true}
}
func Int8IfPositive(i int64) pgtype.Int8 {
	if i > 0 {
		return pgtype.Int8{Int64: i, Valid: true}
	}
	return pgtype.Int8{}
}

// Float4 ---------- Float types ----------
func Float4(f float32) pgtype.Float4 {
	return pgtype.Float4{Float32: f, Valid: true}
}
func Float8(f float64) pgtype.Float8 {
	return pgtype.Float8{Float64: f, Valid: true}
}

// Timestamp ---------- Time / Date ----------
func Timestamp(t time.Time) pgtype.Timestamp {
	if t.IsZero() {
		return pgtype.Timestamp{}
	}
	return pgtype.Timestamp{Time: t, Valid: true}
}

func Timestamptz(t time.Time) pgtype.Timestamptz {
	if t.IsZero() {
		return pgtype.Timestamptz{}
	}
	return pgtype.Timestamptz{Time: t, Valid: true}
}

func Date(t time.Time) pgtype.Date {
	if t.IsZero() {
		return pgtype.Date{}
	}
	return pgtype.Date{Time: t, Valid: true}
}

// UUID ---------- UUID ----------
func UUID(u [16]byte, valid bool) pgtype.UUID {
	if !valid {
		return pgtype.UUID{}
	}
	return pgtype.UUID{Bytes: u, Valid: true}
}
