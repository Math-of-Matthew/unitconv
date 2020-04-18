// Package unitconv is a converter package
// It can return easily the result about time calcs lenght speed etc
package unitconv

const (
	secondhour = 0.00027778
	seconday   = 0.000011574
	monthweek  = 4.345
)

// SecondtoHour (second float32) float32 in hour value
func SecondtoHour(second float32) float32 {
	return second * secondhour
}

// SecondtoDay (second float32) float32 in hour value
func SecondtoDay(second float32) float32 {
	return second * seconday
}

// HourtoSecond (hour float32) float32 in second value
func HourtoSecond(hour float32) float32 {
	return hour / secondhour
}

// HourtoMinutes (hour float32) float32 in minutes value
func HourtoMinutes(hour float32) float32 {
	return hour * 60
}

// DaytoSecond (day float32) float32 in second value
func DaytoSecond(day float32) float32 {
	return day / seconday
}

// DaytoMinutes (day float32) float32 in minutes value
func DaytoMinutes(day float32) float32 {
	return day * 1440
}

// MinutestoDay (day float32) float32 in minutes value
func MinutestoDay(minutes float32) float32 {
	return minutes * 0.00069444
}

// MinutestoHour (minutes float32) float32 in hour value
func MinutestoHour(minutes float32) float32 {
	return minutes / 60
}

// MinutestoSecond (minutes float32) float32 in second value
func MinutestoSecond(minutes float32) float32 {
	return minutes * 60
}

// MinutestoWeek (minutes float32) float32 in week value
func MinutestoWeek(minutes float32) float32 {
	return minutes / 10080
}

// WeektoMinutes (week float32) float32 in Minutes value
func WeektoMinutes(week float32) float32 {
	return week * 10080
}

// WeektoHour (minutes float32) float32 in hour value
func WeektoHour(week float32) float32 {
	return week * 168
}

// WeektoDays (minutes float32) float32 in days value
func WeektoDays(week float32) float32 {
	return week * 7
}

// WeektoMonth (week float32) float32 in Month value
func WeektoMonth(week float32) float32 {
	return week / monthweek
}

// MonthtoHour (Month float32) float32 in Hour value
func MonthtoHour(month float32) float32 {
	return month * 730.001
}

// MonthtoDay (Month float32) float32 in Day value
func MonthtoDay(month float32) float32 {
	return month * 730.001
}

// MonthtoWeek (Month float32) float32 in Week Value
func MonthtoWeek(month float32) float32 {
	return month * monthweek
}

// MonthtoYear (Month float32) float32 in Year Value
func MonthtoYear(month float32) float32 {
	return month / 12
}

// MonthtoDecade (Month float32) float32 in Decade Value
func MonthtoDecade(month float32) float32 {
	return month / 120
}

// MonthtoCentury (Month float32) float32 in Century Value
func MonthtoCentury(month float32) float32 {
	return month / 1199.999
}

// YeartoSecond (year float32) float32 in Second value
func YeartoSecond(year float32) float32 {
	return year * 3.154e+7
}

// YeartoMinute (year float32) float32 in Minute value
func YeartoMinute(year float32) float32 {
	return year * 525600
}

// YeartoHour (year float32) float32
func YeartoHour(year float32) float32 {
	return year * 8760
}

//YeartoDay (year float32) float32 in Day Value
func YeartoDay(year float32) float32 {
	return year * 365
}

// YeartoWeek (year float32) float32 in Week Valye
func YeartoWeek(year float32) float32 {
	return year * 52.143
}

// YeartoMonth (year float32) float32 in Month Value
func YeartoMonth(year float32) float32 {
	return year * 12
}

// YeartoDecade (year float32) float32 in Decade value
func YeartoDecade(year float32) float32 {
	return year / 10
}

// YeartoCentury (year float32) float32 in Century value
func YeartoCentury(year float32) float32 {
	return year / 100
}

// DecadetoHour (decade float32) float32 in Hour Value
func DecadetoHour(decade float32) float32 {
	return decade * 87600
}

// DecadetoDay (decade float32) float32 in Day Value
func DecadetoDay(decade float32) float32 {
	return decade * 3650
}

// DecadetoWeek (decade float32) float32 in Week Value
func DecadetoWeek(decade float32) float32 {
	return decade * 521.429
}

// DecadetoMonth receives (Decade float32) and returns float32 in Month Value
func DecadetoMonth(decade float32) float32 {
	return decade * 120
}

// DecadetoYear receives (Decade float32) and returns float32 in Year Value
func DecadetoYear(decade float32) float32 {
	return decade * 10
}

// DecadetoCentury receives (Decade float32) and returns float32 in Century Value
func DecadetoCentury(decade float32) float32 {
	return decade / 10
}

// CenturytoYear (century float32) float32 in Year Value
func CenturytoYear(century float32) float32 {
	return century / 100
}

// CenturytoDecade (century float32) float32 in Decade Value
func CenturytoDecade(century float32) float32 {
	return century / 10
}
