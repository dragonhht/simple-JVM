package rtda

import "main/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
