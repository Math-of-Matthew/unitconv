// Package unitconv is a converter package
// It can return easily the result about time calcs lenght speed etc
package unitconv

const (
	secondhour = 0.00027778
	seconday   = 0.000011574
	monthweek  = 4.345
)

// SecondtoHour (second float64) float64 in hour value
func SecondtoHour(second float64) float64 {
	return second * secondhour
}

// SecondtoDay (second float64) float64 in hour value
func SecondtoDay(second float64) float64 {
	return second * seconday
}

// HourtoSecond (hour float64) float64 in second value
func HourtoSecond(hour float64) float64 {
	return hour / secondhour
}

// HourtoMinutes (hour float64) float64 in minutes value
func HourtoMinutes(hour float64) float64 {
	return hour * 60
}

// DaytoSecond (day float64) float64 in second value
func DaytoSecond(day float64) float64 {
	return day / seconday
}

// DaytoMinutes (day float64) float64 in minutes value
func DaytoMinutes(day float64) float64 {
	return day * 1440
}

// MinutestoDay (day float64) float64 in minutes value
func MinutestoDay(minutes float64) float64 {
	return minutes * 0.00069444
}

// MinutestoHour (minutes float64) float64 in hour value
func MinutestoHour(minutes float64) float64 {
	return minutes / 60
}

// MinutestoSecond (minutes float64) float64 in second value
func MinutestoSecond(minutes float64) float64 {
	return minutes * 60
}

// MinutestoWeek (minutes float64) float64 in week value
func MinutestoWeek(minutes float64) float64 {
	return minutes / 10080
}

// WeektoMinutes (week float64) float64 in Minutes value
func WeektoMinutes(week float64) float64 {
	return week * 10080
}

// WeektoHour (minutes float64) float64 in hour value
func WeektoHour(week float64) float64 {
	return week * 168
}

// WeektoDays (minutes float64) float64 in days value
func WeektoDays(week float64) float64 {
	return week * 7
}

// WeektoMonth (week float64) float64 in Month value
func WeektoMonth(week float64) float64 {
	return week / monthweek
}

// MonthtoHour (Month float64) float64 in Hour value
func MonthtoHour(month float64) float64 {
	return month * 730.001
}

// MonthtoDay (Month float64) float64 in Day value
func MonthtoDay(month float64) float64 {
	return month * 730.001
}

// MonthtoWeek (Month float64) float64 in Week Value
func MonthtoWeek(month float64) float64 {
	return month * monthweek
}

// MonthtoYear (Month float64) float64 in Year Value
func MonthtoYear(month float64) float64 {
	return month / 12
}

// MonthtoDecade (Month float64) float64 in Decade Value
func MonthtoDecade(month float64) float64 {
	return month / 120
}

// MonthtoCentury (Month float64) float64 in Century Value
func MonthtoCentury(month float64) float64 {
	return month / 1199.999
}

// YeartoSecond (year float64) float64 in Second value
func YeartoSecond(year float64) float64 {
	return year * 3.154e+7
}

// YeartoMinute (year float64) float64 in Minute value
func YeartoMinute(year float64) float64 {
	return year * 525600
}

// YeartoHour (year float64) float64
func YeartoHour(year float64) float64 {
	return year * 8760
}

//YeartoDay (year float64) float64 in Day Value
func YeartoDay(year float64) float64 {
	return year * 365
}

// YeartoWeek (year float64) float64 in Week Valye
func YeartoWeek(year float64) float64 {
	return year * 52.143
}

// YeartoMonth (year float64) float64 in Month Value
func YeartoMonth(year float64) float64 {
	return year * 12
}

// YeartoDecade (year float64) float64 in Decade value
func YeartoDecade(year float64) float64 {
	return year / 10
}

// YeartoCentury (year float64) float64 in Century value
func YeartoCentury(year float64) float64 {
	return year / 100
}

// DecadetoHour (decade float64) float64 in Hour Value
func DecadetoHour(decade float64) float64 {
	return decade * 87600
}

// DecadetoDay (decade float64) float64 in Day Value
func DecadetoDay(decade float64) float64 {
	return decade * 3650
}

// DecadetoWeek (decade float64) float64 in Week Value
func DecadetoWeek(decade float64) float64 {
	return decade * 521.429
}

// DecadetoMonth receives (Decade float64) and returns float64 in Month Value
func DecadetoMonth(decade float64) float64 {
	return decade * 120
}

// DecadetoYear receives (Decade float64) and returns float64 in Year Value
func DecadetoYear(decade float64) float64 {
	return decade * 10
}

// DecadetoCentury receives (Decade float64) and returns float64 in Century Value
func DecadetoCentury(decade float64) float64 {
	return decade / 10
}

// CenturytoYear (century float64) float64 in Year Value
func CenturytoYear(century float64) float64 {
	return century / 100
}

// CenturytoDecade (century float64) float64 in Decade Value
func CenturytoDecade(century float64) float64 {
	return century / 10
}
