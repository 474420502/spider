package spider

import "github.com/474420502/requests"

// IUrls The Interface Of Spider Execute
type IUrls interface {
	GetUrls() []string
}

// IExecute The Interface Of Spider Execute
type IExecute interface {
	Execute(*Context)
}

// IBefore 预处理
type IBefore interface {
	Before(*Context)
}

// ITask The Interface Of Task
type ITask interface {
	IExecute
}

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

	share map[string]interface{}
	retry int

	Is *SettingContext
}

// GetUrls Get return urls []string
func (ctx *Context) GetUrls() []string {
	return ctx.urls
}

// SetUrls Set urls []string
func (ctx *Context) SetUrls(urls []string) {
	ctx.urls = urls
}

// GetWorkflow Get return workflow *requests.Workflow
func (ctx *Context) GetWorkflow() *requests.Workflow {
	return ctx.workflow
}

// SetWorkflow Set workflow *requests.Workflow
func (ctx *Context) SetWorkflow(workflow *requests.Workflow) {
	ctx.workflow = workflow
}

// GetSession Get return session *requests.Session
func (ctx *Context) GetSession() *requests.Session {
	return ctx.session
}

// SetSession Set session *requests.Session
func (ctx *Context) SetSession(session *requests.Session) {
	ctx.session = session
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
