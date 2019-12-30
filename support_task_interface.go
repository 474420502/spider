package spider

import "time"

// IPriority 优先接口
type IPriority interface {
	GetPriority() int
}

// ITaskBefore Execute执行前处理
type ITaskBefore interface {
	Before(*Context)
}

// ITaskAfter Execute执行后处理
type ITaskAfter interface {
	After(*Context)
}

// IPlanTime about time to plan
type IPlanTime interface {
	Next() bool
	GetExecuteTime() *time.Time
}
