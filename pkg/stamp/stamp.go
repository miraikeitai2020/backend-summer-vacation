package stamp

import(
	"time"

	// import sample API packages
	"github.com/miraikeitai2020/backend-summer-vacation/pkg/server/model"
)

func Now() model.TimeStamp {
	return model.TimeStamp {
		TimeStamp:	time.Now().String(),
		Detail: model.Detail {
			Date: string(time.Now().Format("2006-01-02")),
			Time: string(time.Now().Format("15:04:05")),
		},
	}
}