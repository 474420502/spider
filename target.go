package spider

import (
	"log"
	"sync/atomic"

	"github.com/474420502/focus/compare"
	pqueue "github.com/474420502/focus/priority_queue"
	"github.com/474420502/requests"
)

// SettingTarget 一些判断操作
type SettingTarget struct {
	isRunning  int32
	isTaskOnce bool
}

// NewTarget 目标
func NewTarget() *Target {
	return NewTargetMaxPriority()
}

// NewTargetMaxPriority 目标
func NewTargetMaxPriority() *Target {
	target := new(Target)

	target.tasks = pqueue.New(PriorityMax)
	target.preparedTasks = pqueue.New(PriorityMax)
	target.subTasks = pqueue.New(subPriorityMax)

	target.share = make(map[string]interface{})
	target.Is = &SettingTarget{isRunning: 0, isTaskOnce: false}

	return target
}

// NewTargetMinPriority 目标
func NewTargetMinPriority() *Target {
	target := new(Target)

	target.tasks = pqueue.New(PriorityMin)
	target.preparedTasks = pqueue.New(PriorityMin)
	target.subTasks = pqueue.New(subPriorityMin)

	target.share = make(map[string]interface{})
	target.Is = &SettingTarget{isRunning: 0, isTaskOnce: false}

	return target
}

// Target 目标
type Target struct {
	session *requests.Session
	share   map[string]interface{}

	tasks         *pqueue.PriorityQueue
	preparedTasks *pqueue.PriorityQueue
	subTasks      *pqueue.PriorityQueue

	beforeEveryTask func(*Context)

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

// GetTaskOnce Get isTaskOnce
func (target *Target) GetTaskOnce() bool {
	return target.Is.isTaskOnce
}

// SetTaskOnce Set isTaskOnce
func (target *Target) SetTaskOnce(is bool) {
	target.Is.isTaskOnce = is
}

// SetPriorityCompare Set priorityCompare compare.Compare 自定义优先
func (target *Target) SetPriorityCompare(compare compare.Compare) {
	tasks := pqueue.New(compare)
	preparedTasks := pqueue.New(compare)

	if atomic.LoadInt32(&target.Is.isRunning) > 0 {
		panic("SetPriorityCompare,  App can not be Running.")
	}

	for _, v := range target.tasks.Values() {
		tasks.Push(v)
	}

	for _, v := range target.preparedTasks.Values() {
		preparedTasks.Push(v)
	}

	target.tasks = tasks
	target.preparedTasks = preparedTasks
	target.priorityCompare = compare
}

// GetSession Get return session *requests.Session
func (target *Target) GetSession() *requests.Session {
	if target.session == nil {
		target.session = requests.NewSession()
	}
	return target.session
}

// SetSession Set session *requests.Session
func (target *Target) SetSession(session *requests.Session) {
	target.session = session
}

// // GetURL Get return url string
// func (target *Target) GetURL() string {
// 	return target.url
// }

// // SetURL Set url string
// func (target *Target) SetURL(url string) {
// 	target.url = url
// }

// AddTask 添加任务
func (target *Target) AddTask(task ITask) {
	target.tasks.Push(task)
}

// BeforeEveryTask 添加任务
func (target *Target) BeforeEveryTask(before func(*Context)) {
	target.beforeEveryTask = before
}

// BeforeEveryTask 添加任务
func (target *Target) checkRunning() bool {
	return atomic.LoadInt32(&target.Is.isRunning) > 0
}

func (target *Target) processingContext(ctx *Context) {

}

// StartTask 添加任务
func (target *Target) StartTask() {

	atomic.StoreInt32(&target.Is.isRunning, 1)

	ctx := &Context{target: target, share: target.share}
	ctx.SetRetry(0)

LOOP:
	for atomic.LoadInt32(&target.Is.isRunning) > 0 {

		if itask, ok := target.tasks.Pop(); ok {

			if urls, ok := itask.(IUrls); ok {
				ctx.urls = urls.GetUrls()
			}

			if target.beforeEveryTask != nil {
				target.beforeEveryTask(ctx)
				if !target.checkRunning() {
					break LOOP
				}
			}

			if task, ok := itask.(ITask); ok {
				task.Execute(ctx)
				if !target.checkRunning() {
					break LOOP
				}
			} else {
				log.Fatalln("task must have the method of Execute")
			}

			target.preparedTasks.Push(itask)
			for target.subTasks.Size() != 0 {

				if isub, ok := target.subTasks.Pop(); ok {
					switch sub := isub.(type) {
					case func(*Context):
						sub(ctx)
					case IExecute:
						sub.Execute(ctx)
					}
				}

				if !target.checkRunning() {
					break LOOP //退出 for 达到退出程序目的
				}

			}

		} else if target.Is.isTaskOnce {
			break
		} else {
			target.tasks, target.preparedTasks = target.preparedTasks, target.tasks
		}

	}

	for itask, ok := target.tasks.Pop(); ok; itask, ok = target.tasks.Pop() {
		target.preparedTasks.Push(itask)
	}
	target.tasks, target.preparedTasks = target.preparedTasks, target.tasks

	atomic.StoreInt32(&target.Is.isRunning, 0)
}

// StopTask 停止任务
func (target *Target) StopTask() {
	atomic.StoreInt32(&target.Is.isRunning, 0)
}
