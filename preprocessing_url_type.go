package spider

// IPreprocessingUrl Get Url To Execute
type IPreprocessingUrl interface {
	PreprocessingUrl(ctx *Context)
}

// PreGetUrl Task的 Get url 预处理组件
type PreGetUrl string

func (h PreGetUrl) PreprocessingUrl(ctx *Context) {
	ctx.SetWorkflow(ctx.GetSession().Get((string)(h)))
}

// PrePostUrl Task的 Post url 预处理组件
type PrePostUrl string

func (h PrePostUrl) PreprocessingUrl(ctx *Context) {
	ctx.SetWorkflow(ctx.GetSession().Post((string)(h)))
}

// PrePutUrl Task的 Put url 预处理组件
type PrePutUrl string

func (h PrePutUrl) PreprocessingUrl(ctx *Context) {
	ctx.SetWorkflow(ctx.GetSession().Put((string)(h)))
}

// PHeadUrl Task的 Head url 预处理组件
type PHeadUrl string

func (h PHeadUrl) PreprocessingUrl(ctx *Context) {
	ctx.SetWorkflow(ctx.GetSession().Head((string)(h)))
}

// PPatchUrl Task的 Patch url 预处理组件
type PPatchUrl string

func (h PPatchUrl) PreprocessingUrl(ctx *Context) {
	ctx.SetWorkflow(ctx.GetSession().Patch((string)(h)))
}

// PDeleteUrl Task的 Delete url 预处理组件
type PDeleteUrl string

func (h PDeleteUrl) PreprocessingUrl(ctx *Context) {
	ctx.SetWorkflow(ctx.GetSession().Delete((string)(h)))
}

// POptionsUrl Task的 Options url 预处理组件
type POptionsUrl string

func (h POptionsUrl) PreprocessingUrl(ctx *Context) {
	ctx.SetWorkflow(ctx.GetSession().Options((string)(h)))
}
