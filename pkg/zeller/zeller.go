package zeller
import (
	"time"

	"github.com/miraikeitai2020/backend-summer-vacation/pkg/server/model"
)

func Zeller(year,month,day int) (week string) {
	year,month=check(year,month)
	switch Zellertask(year,month,day){
	case 0:
		week ="Saturday"
	case 1:
		week = "Sunday"
	case 2:
		week = "Monday"
	case 3:
		week = "Tuesday"
	case 4:
		week = "Wednesday"
	case 5:
		week = "Thursday"
	case 6:
		week = "Friday"

	}
	return
}

func check(year,month int)(int,int){
	if month<2{
		month=month+1
		year=year-1
	}
	return year,month
}

func  Zellertask(year,month,day int)int  {
	zellers:=year+year/100+year/400+(13*month+8)/5+day
	return(zellres%7)
}