package spider

import "time"

// PlanOnTime 准时
type PlanOnTime struct {
	planTime time.Time
	hasNext  bool
	IPlanTime
}

func (plan *PlanOnTime) SetExecuteTime(t time.Time) {
	plan.planTime = t
	plan.hasNext = true
}

func (plan *PlanOnTime) Next() bool {
	return plan.hasNext
}

func (plan *PlanOnTime) GetExecuteTime() *time.Time {
	plan.hasNext = false
	return &plan.planTime
}
