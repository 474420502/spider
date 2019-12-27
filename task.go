package spider

import "github.com/474420502/requests"

// IPreprocessingUrl Get Url To Execute
type IPreprocessingUrl interface {
	PreprocessingUrl(ctx *Context)
}

// IUrls Get Urls To Execute
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
