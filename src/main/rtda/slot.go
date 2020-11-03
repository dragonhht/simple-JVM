package rtda

import "github.com/kuangcp/simple-jvm/src/main/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
