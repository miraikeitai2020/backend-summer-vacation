package zeller

const (
	DAY_OF_SUN = "Sunday"
	DAY_OF_MON = "Monday"
	DAY_OF_TUE = "Tuesday"
	DAY_OF_WED = "Wednesday"
	DAY_OF_THU = "Thursday"
	DAY_OF_FRI = "Friday"
	DAY_OF_SAT = "Saturday"
)

func Zeller(year, month, day int) (week string) {
	year, month = check(year,month)
	switch subZeller(year, month, day) {
	case 0:
		week = DAY_OF_SUN
	case 1:
		week = DAY_OF_MON
	case 2:
		week = DAY_OF_TUE
	case 3:
		week = DAY_OF_WED
	case 4:
		week = DAY_OF_THU
	case 5:
		week = DAY_OF_FRI
	case 6:
		week = DAY_OF_SAT
	}
	return
}


func check(year, month int) (int, int) {
	if month < 3 {
		return year - 1, month + 12
	}
	return year, month
}

func subZeller(year, month, day int) (int) {
	return (year + year/4 - year/100 + year/400 + (13*month + 8)/5 + day)%7
}
