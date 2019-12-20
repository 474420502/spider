package spider

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
	t1, t2 := k1.(ITask), k2.(ITask)
	if t1.GetPriority() > t2.GetPriority() {
		return 1
	}
	return -1
}

// PriorityMin 最小值优先
func PriorityMin(k1, k2 interface{}) int {
	t1, t2 := k1.(ITask), k2.(ITask)
	if t1.GetPriority() < t2.GetPriority() {
		return 1
	}
	return -1
}
