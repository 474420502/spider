package spider

import (
	"github.com/474420502/focus/compare"
	pqueue "github.com/474420502/focus/priority_queue"
	"github.com/474420502/requests"
	"sync/atomic"
)

// SettingTarget 一些判断操作
type SettingTarget struct {
	isRunning  int32
	isTaskOnce bool
}

// Target 目标
type Target struct {
	url             string
	session         *requests.Session
	tasks           *pqueue.PriorityQueue
	priorityCompare compare.Compare
	Is              *SettingTarget
}

// GetPriorityCompare Get return priorityCompare compare.Compare
// func (target *Target) GetPriorityCompare() compare.Compare {
// 	return target.priorityCompare
// }

// SetPriorityCompare Set priorityCompare compare.Compare 自定义优先
func (target *Target) SetPriorityCompare(compare compare.Compare) {
	queue := pqueue.New(compare)

	if atomic.LoadInt32(&target.Is.isRunning) > 0 {
		panic("SetPriorityCompare,  App can not be Running.")
	}

	iter := target.tasks.Iterator()
	for iter.Next() {
		queue.Push(iter.Value())
	}
	target.tasks = queue
	target.priorityCompare = compare
}

// GetSession Get return session *requests.Session
func (target *Target) GetSession() *requests.Session {
	return target.session
}

// SetSession Set session *requests.Session
func (target *Target) SetSession(session *requests.Session) {
	target.session = session
}

// GetURL Get return url string
func (target *Target) GetURL() string {
	return target.url
}

// SetURL Set url string
func (target *Target) SetURL(url string) {
	target.url = url
}

// AddTask 添加任务
func (target *Target) AddTask(task ITask) {
	target.tasks.Push(task)
}

// ChangePriority 添加任务
func (target *Target) ChangePriority(task ITask, priority int) {
	target.tasks.Push(task)
}

func (target *Target) processingContext(ctx *Context) {

}

// StartTask 添加任务
func (target *Target) StartTask() {

	iter := target.tasks.Iterator()
	atomic.StoreInt32(&target.Is.isRunning, 1)

	ctx := &Context{target: target}
	ctx.SetRetry(0)

	for atomic.LoadInt32(&target.Is.isRunning) > 0 {
		if iter.Next() {

			task := iter.Value().(ITask)
			task.Execute(ctx)

		} else if target.Is.isTaskOnce {
			break
		} else {
			iter.ToHead()
		}
	}
	atomic.StoreInt32(&target.Is.isRunning, 1)
}

// StopTask 停止任务
func (target *Target) StopTask() {
	atomic.StoreInt32(&target.Is.isRunning, 0)
}

// NewTarget 目标
func NewTarget() *Target {
	target := new(Target)
	target.tasks = pqueue.New(PriorityMax)
	return target
}
