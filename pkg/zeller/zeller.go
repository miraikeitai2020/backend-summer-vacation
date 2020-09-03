package zeller

const(
	DAY_OF_SAT = "Saturday"
	DAY_OF_SUN = "Sunday"
	DAY_OF_MON = "Monday"
	DAY_OF_TUE = "Tuesday"
	DAY_OF_WED = "Wednesday"
	DAY_OF_THU = "Thursday"
	DAY_OF_FRI = "Friday"
)

func Zeller(year, month, day int) (week string) {
	year, month = checkArg(year, month)
	switch calcDayOfWeek(year, month, day) {
	case 0:
		week = DAY_OF_SAT
	case 1:
		week = DAY_OF_SUN
	case 2:
		week = DAY_OF_MON
	case 3:
		week = DAY_OF_TUE
	case 4:
		week = DAY_OF_WED
	case 5:
		week = DAY_OF_THU
	case 6:
		week = DAY_OF_FRI
	}
	return
}

func checkArg(year, month int) (int, int) {
	if month < 3 {
		return year - 1, month + 12
	}
	return year, month
}

func calcDayOfWeek(year, month, day int) int {
	top, bottom := divideYearElements(year)
	return (day + ((month + 1) * 26 / 10) + bottom + (bottom / 4) + (top / 4) - 2 * top) % 7
}

func divideYearElements(year int) (top, bottom int) {
	top = year/100
	bottom = year%100
	return
}
