package spider

// Context Context of Task
type Context struct {
	target *Target
	share  interface{}

	retry int
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
func (ctx *Context) GetShare() interface{} {
	return ctx.share
}

// SetShare Set share interface{}
func (ctx *Context) SetShare(share interface{}) {
	ctx.share = share
}

// GetTarget Get return target *Target
func (ctx *Context) GetTarget() *Target {
	return ctx.target
}

// ITask the interface of task
type ITask interface {
	Execute(*Context)
	GetPriority() int
}
