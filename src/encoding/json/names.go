// Copyright (c) 2017 Gerald Sangudi. All rights reserved.

package json

import (
	"sync"
)

// Global hash of names used in name-value objects.
type nameHash struct {
	sync.RWMutex

	names map[string]string
}

var _NAME_HASH = &nameHash{
	names: make(map[string]string, 64*1024),
}

func (this *nameHash) hash(in string) (out string) {

	this.RLock()
	out, ok := this.names[in]
	this.RUnlock()

	if ok {
		return
	}

	this.Lock()
	out, ok = this.names[in]
	if !ok {
		this.names[in] = in
		out = in
	}
	this.Unlock()
	return
}
