//go:generate stringer -type=Weekday

package gogenerate

// Weekday 表示一周中的某一天
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
