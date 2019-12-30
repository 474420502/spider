package spider

import "github.com/474420502/requests"

// ITask The Interface Of Task
type ITask interface {
	Execute(*Context)
}

// TaskQueueType 任务队列的类型
type TaskQueueType int

const (
	_ TaskQueueType = iota
	// TypeTaskUnknown 未知类型
	TypeTaskUnknown
	// TypeTaskMain 主任务
	TypeTaskMain
	// TypeTaskPlan 时间任务
	TypeTaskPlan
	// TypeTaskSub 子任务
	TypeTaskSub
)

// SettingContext 一些判断操作
type SettingContext struct {
	IsRestart int32
	IsSubTask bool
}

// Context Context of Task
type Context struct {
	target *Target

	urls     []string
	session  *requests.Session
	workflow *requests.Workflow
	content  string

	share    map[string]interface{}
	retry    int
	taskType TaskQueueType

	Is *SettingContext
}

// GetTaskType Get return taskType TaskQueueType
func (ctx *Context) GetTaskType() TaskQueueType {
	return ctx.taskType
}

// Execute will do list:
// 1: Get return ctx.workflow.Execute() == ctx.Execute()
// 2: ctx.context = response.Content()
func (ctx *Context) Execute() (*requests.Response, error) {
	resp, err := ctx.workflow.Execute()
	ctx.content = resp.Content()
	return resp, err
}

// Content return cxt
func (ctx *Context) Content() string {
	return ctx.content
}

// GetWorkflow Get return workflow *requests.Workflow
func (ctx *Context) GetWorkflow() *requests.Workflow {
	return ctx.workflow
}

// SetWorkflow Set workflow *requests.Workflow
func (ctx *Context) SetWorkflow(workflow *requests.Workflow) {
	ctx.workflow = workflow
}

// GetSession Get return target.session *requests.Session
func (ctx *Context) GetSession() *requests.Session {
	return ctx.target.GetSession()
}

// SetSession Set target.session *requests.Session
func (ctx *Context) SetSession(session *requests.Session) {
	ctx.target.SetSession(session)
}

// GetRetry Get return retry int
func (ctx *Context) GetRetry() int {
	return ctx.retry
}

// SetRetry Set retry int
func (ctx *Context) SetRetry(retry int) {
	ctx.retry = retry
}

// GetShare Get return share interface{}
func (ctx *Context) GetShare() map[string]interface{} {
	return ctx.share
}

// GetTarget Get return target *Target
func (ctx *Context) GetTarget() *Target {
	return ctx.target
}
