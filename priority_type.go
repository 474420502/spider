package spider

// IPriority 优先接口
type IPriority interface {
	GetPriority() int
}

// PriorityInt Int优先级
type PriorityInt int

// GetPriority Get priority
func (pri PriorityInt) GetPriority() int {
	return (int)(pri)
}

// PriorityInt32 Int优先级
type PriorityInt32 int32

// GetPriority Get priority
func (pri PriorityInt32) GetPriority() int {
	return (int)(pri)
}

// PriorityInt64 Int优先级
type PriorityInt64 int64

// GetPriority Get priority
func (pri PriorityInt64) GetPriority() int {
	return (int)(pri)
}

// PriorityFloat32 Int优先级
type PriorityFloat32 float32

// GetPriority Get priority
func (pri PriorityFloat32) GetPriority() int {
	return (int)(pri)
}

// PriorityMax 最大值优先
func PriorityMax(k1, k2 interface{}) int {
	p1, p2 := 0, 0

	if priority, ok := k1.(IPriority); ok {
		p1 = priority.GetPriority()
	}

	if priority, ok := k2.(IPriority); ok {
		p2 = priority.GetPriority()
	}

	if p1 > p2 {
		return 1
	}
	return -1
}

// PriorityMin 最小值优先
func PriorityMin(k1, k2 interface{}) int {

	p1, p2 := 0, 0

	if priority, ok := k1.(IPriority); ok {
		p1 = priority.GetPriority()
	}

	if priority, ok := k2.(IPriority); ok {
		p2 = priority.GetPriority()
	}

	if p1 < p2 {
		return 1
	}
	return -1

}

// subPriorityMax 最大值优先
func subPriorityMax(k1, k2 interface{}) int {

	p1, p2 := 0, 0

	switch priority := k1.(type) {
	case IPriority:
		p1 = priority.GetPriority()
	case func(*Context):
		p1 = 0
	}

	switch priority := k2.(type) {
	case IPriority:
		p2 = priority.GetPriority()
	case func(*Context):
		p2 = 0
	}

	if p1 > p2 {
		return 1
	}
	return -1
}

// subPriorityMin 最小值优先
func subPriorityMin(k1, k2 interface{}) int {

	p1, p2 := 0, 0

	switch priority := k1.(type) {
	case IPriority:
		p1 = priority.GetPriority()
	case func(*Context):
		p1 = 0
	}

	switch priority := k2.(type) {
	case IPriority:
		p2 = priority.GetPriority()
	case func(*Context):
		p2 = 0
	}

	if p1 < p2 {
		return 1
	}
	return -1
}

// timePriority 最小值优先
func timePriority(k1, k2 interface{}) int {
	t1, t2 := k1.(IPlanTime).GetExecuteTime(), k2.(IPlanTime).GetExecuteTime()
	if t1.Before(*t2) {
		return 1
	}
	return -1
}
