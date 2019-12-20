package spider

import "testing"

type MyTask1 struct {
	PriorityInt
	itask ITask
}

// Execute
func (mt *MyTask1) Execute(ctx *Context) {
	target := ctx.GetTarget()
	ses := target.GetSession()
	ses.Get(target.url)
}

func TestTargetCase1(t *testing.T) {

	target := NewTarget()

	target.SetURL("http://www.baidu.com")

	target.AddTask(&MyTask1{PriorityInt: 1})
	target.AddTask(&MyTask1{PriorityInt: 4})
	target.AddTask(&MyTask1{PriorityInt: 3})

	if task, ok := target.tasks.Top(); ok {
		if task.(ITask).GetPriority() != 4 {
			t.Error("task GetPriority error")
		}
	} else {
		t.Error("addtask error")
	}

}
