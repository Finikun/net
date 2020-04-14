package http2

import (
	"golang.org/x/net/http2/hpack"
	"sync"
)

func Free(obj interface{}) {
	switch obj.(type) {
	case *MetaHeadersFrame:
		f := obj.(*MetaHeadersFrame)
		freeMetaHeadersFrame(f)
	}
}

var metaHeadersFramePool = sync.Pool{
	New: func() interface{} {
		mh := &MetaHeadersFrame{
			Fields: make([]hpack.HeaderField, 0, 10),
		}
		return mh
	},
}

func newMetaHeadersFrame() *MetaHeadersFrame {
	mh := metaHeadersFramePool.Get().(*MetaHeadersFrame)
	return mh
}

func freeMetaHeadersFrame(mh *MetaHeadersFrame) {
	mh.HeadersFrame = nil
	mh.Fields = mh.Fields[:0]
	mh.Truncated = false
	metaHeadersFramePool.Put(mh)
}
