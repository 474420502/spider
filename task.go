package spider

// IExecute The Interface Of Spider Execute
type IExecute interface {
	Execute(*Context)
}

// ITask The Interface Of Task
type ITask interface {
	IPriority
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

	share map[string]interface{}
	index int
	retry int

	Is *SettingContext
}

// GetIndex Get return index int
func (ctx *Context) GetIndex() int {
	return ctx.index
}

// SetIndex Set index int
func (ctx *Context) SetIndex(index int) {
	ctx.index = index
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
