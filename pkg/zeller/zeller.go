package zeller
import (
	"time"

	"github.com/miraikeitai2020/backend-summer-vacation/pkg/server/model"
)

func Now() model.Zeller {
	return model.Zeller{
		Zeller: zeller.Now().String(),
		ZellerElements: model.ZellerElements{
			Year: string(time.Now().Format("Monday")),
			Month: string(time.Now().Format("September")),
			Day: string(time.Now().Format("first")),

		},
	}
}
