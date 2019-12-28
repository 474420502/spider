package spider

import "time"

// ITimePlan about time to plan
type ITimePlan interface {
	GetExecuteTime() *time.Time
}
