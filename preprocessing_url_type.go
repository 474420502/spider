package spider

// PreGetUrl Task的 Get url 预处理组件
type PreGetUrl string

func (h PreGetUrl) Before(ctx *Context) {
	ctx.SetWorkflow(ctx.GetSession().Get((string)(h)))
}

// PrePostUrl Task的 Post url 预处理组件
type PrePostUrl string

func (h PrePostUrl) Before(ctx *Context) {
	ctx.SetWorkflow(ctx.GetSession().Post((string)(h)))
}

// PrePutUrl Task的 Put url 预处理组件
type PrePutUrl string

func (h PrePutUrl) Before(ctx *Context) {
	ctx.SetWorkflow(ctx.GetSession().Put((string)(h)))
}

// PreHeadUrl Task的 Head url 预处理组件
type PreHeadUrl string

func (h PreHeadUrl) Before(ctx *Context) {
	ctx.SetWorkflow(ctx.GetSession().Head((string)(h)))
}

// PrePatchUrl Task的 Patch url 预处理组件
type PrePatchUrl string

func (h PrePatchUrl) Before(ctx *Context) {
	ctx.SetWorkflow(ctx.GetSession().Patch((string)(h)))
}

// PreDeleteUrl Task的 Delete url 预处理组件
type PreDeleteUrl string

func (h PreDeleteUrl) Before(ctx *Context) {
	ctx.SetWorkflow(ctx.GetSession().Delete((string)(h)))
}

// PreOptionsUrl Task的 Options url 预处理组件
type PreOptionsUrl string

func (h PreOptionsUrl) Before(ctx *Context) {
	ctx.SetWorkflow(ctx.GetSession().Options((string)(h)))
}
