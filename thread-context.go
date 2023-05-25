package threadContext

import (
	"context"
	"unsafe"
)

var threadMetaFlag uint64

func init() {
	for i, b := range []byte("g.labels") {
		t := uint64(b) << uint8(i*8)
		threadMetaFlag |= t
	}
}

// 48b
type threadMeta struct {
	labels labelMap        // pprof.labelMap
	flag   uint64          // 64bit flag
	ctx    context.Context // runtime.iface
	data   interface{}     // runtime.eface
}

func SetMeta(ctx context.Context, data interface{}) {
	cm := currentMeta(true)
	cm.ctx = ctx
	cm.data = data
}

func GetMeta() (ctx context.Context, data interface{}) {
	cm := currentMeta(false)
	if cm != nil {
		ctx, data = cm.ctx, cm.data
	}
	return
}

func currentMeta(create bool) (cm *threadMeta) {
	ptr := runtime_getProfLabel()
	if !create && ptr == nil {
		return nil
	}
	if ptr == nil {
		cm = &threadMeta{flag: threadMetaFlag}

		//runtime.SetFinalizer(cm, (*threadMeta).finalize)
		runtime_setProfLabel(unsafe.Pointer(cm))
	} else if cm = (*threadMeta)(ptr); cm.flag != threadMetaFlag {
		cm = &threadMeta{
			labels: cm.labels,
			flag:   threadMetaFlag,
		}

		//runtime.SetFinalizer(cm, (*threadMeta).finalize)
		runtime_setProfLabel(unsafe.Pointer(cm))
	}
	return cm
}

// finalize reset thread's memory.
func (t *threadMeta) finalize() {
	t.labels = nil
	t.flag = 0
	t.ctx = nil
	t.data = nil
}
