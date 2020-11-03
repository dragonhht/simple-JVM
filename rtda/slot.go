package rtda

import "simple-jvm/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
