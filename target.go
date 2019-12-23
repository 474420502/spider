package spider

import (
	"sync/atomic"

	"github.com/474420502/focus/compare"
	heap "github.com/474420502/focus/tree/heap"
	"github.com/474420502/requests"
)

// SettingTarget 一些判断操作
type SettingTarget struct {
	isRunning  int32
	isTaskOnce bool
}

// Target 目标
type Target struct {
	url     string
	session *requests.Session
	share   map[string]interface{}

	tasks         *heap.Tree
	preparedTasks *heap.Tree
	subTasks      *heap.Tree

	priorityCompare compare.Compare
	Is              *SettingTarget
}

// GetShare Get return share interface{}
func (target *Target) GetShare() map[string]interface{} {
	return target.share
}

// GetPriorityCompare Get return priorityCompare compare.Compare
// func (target *Target) GetPriorityCompare() compare.Compare {
// 	return target.priorityCompare
// }

// SetPriorityCompare Set priorityCompare compare.Compare 自定义优先
func (target *Target) SetPriorityCompare(compare compare.Compare) {
	queue := heap.New(compare)

	if atomic.LoadInt32(&target.Is.isRunning) > 0 {
		panic("SetPriorityCompare,  App can not be Running.")
	}

	for _, v := range target.tasks.Values() {
		queue.Put(v)
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
	target.tasks.Put(task)
}

// AppendTask 添加任务
func (target *Target) AppendTask(task ITask) {
	target.tasks.Put(task)
}

func (target *Target) processingContext(ctx *Context) {

}

// StartTask 添加任务
func (target *Target) StartTask() {

	target.preparedTasks = heap.New(target.priorityCompare)

	atomic.StoreInt32(&target.Is.isRunning, 1)

	ctx := &Context{target: target, share: target.share}
	ctx.SetRetry(0)
	ctx.SetIndex(0)

	for atomic.LoadInt32(&target.Is.isRunning) > 0 {
		if itask, ok := target.tasks.Pop(); ok {

			task := itask.(ITask)
			task.Execute(ctx)

			target.preparedTasks.Put(itask)

			for !target.subTasks.Empty() {

				if isub, ok := target.subTasks.Pop(); ok {
					switch sub := isub.(type) {
					case func(*Context):
						sub(ctx)
					case IExecute:
						sub.Execute(ctx)
					}
				}

			}

			ctx.SetIndex(ctx.GetIndex() + 1)

		} else if target.Is.isTaskOnce {
			break
		} else {
			target.tasks, target.preparedTasks = target.preparedTasks, target.tasks
			ctx.SetIndex(0)
		}
	}
	atomic.StoreInt32(&target.Is.isRunning, 0)
}

// StopTask 停止任务
func (target *Target) StopTask() {
	atomic.StoreInt32(&target.Is.isRunning, 0)
}

// NewTarget 目标
func NewTarget() *Target {
	target := new(Target)

	target.tasks = heap.New(PriorityMax)
	target.subTasks = heap.New(subPriorityMax)

	target.share = make(map[string]interface{})
	return target
}
