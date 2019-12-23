package spider

import (
	"testing"

	pqueue "github.com/474420502/focus/priority_queue"
	"github.com/Pallinder/go-randomdata"
)

type MyTask1 struct {
	PriorityInt
	IExecute
}

// Execute
func (mt *MyTask1) Execute(ctx *Context) {

	target := ctx.GetTarget()
	ses := target.GetSession()

	wf := ses.Get(target.url)
	resp, err := wf.Execute()

	if err != nil {
		panic(err)
	}

	resp.Content()

	mt.PriorityInt = 5

}

type X struct {
	Value int
	Name  string
}

func Comp(k1, k2 interface{}) int {
	p1, p2 := k1.(X), k2.(X)
	if p1.Value > p2.Value {
		return 1
	}
	return -1
}

func TestP(t *testing.T) {

	testN := 10
	testV := 500
	for n := 0; n < testN; n++ {

		queue := pqueue.New(Comp)

		content := ""
		for i := 0; i < 1000; i++ {
			if i%2 == 0 {
				a := randomdata.StringSample("abcdefghijklmnopqrstuvwxyz")
				b := randomdata.StringSample("1234567890")
				content += a + b
				queue.Push(X{Value: testV, Name: a + b})
			} else {
				r := randomdata.Number(0, 1000)
				if r != testV {
					queue.Push(X{Value: r, Name: "a"})
				}
			}
		}

		content2 := ""
		for v, ok := queue.Pop(); ok; v, ok = queue.Pop() {
			if v.(X).Value == testV {
				content2 += v.(X).Name
			}
		}

		if content != content2 {
			t.Error("content != content2")
			break
		}
	}
}

func TestTargetCase1(t *testing.T) {

	target := NewTarget()

	target.SetTaskOnce(true)
	target.SetURL("http://www.baidu.com")

	target.AddTask(&MyTask1{PriorityInt: 1})
	target.AddTask(&MyTask1{PriorityInt: 4})
	target.AddTask(&MyTask1{PriorityInt: 3})

	if task, ok := target.tasks.Top(); ok {
		if task.(IPriority).GetPriority() != 4 {
			t.Error("task GetPriority error")
		}
	} else {
		t.Error("addtask error")
	}

	target.StartTask()
}
