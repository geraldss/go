// Copyright (c) 2017 Gerald Sangudi. All rights reserved.

package encoding

import (
	"sync"
)

func HashName(in string) (out string) {
	return NAME_HASH.Hash(in)
}

// Global hash of names used in name-value objects.
type NameHash struct {
	sync.RWMutex

	names map[string]string
}

var NAME_HASH = &NameHash{
	names: make(map[string]string, 64*1024),
}

func (this *NameHash) Hash(in string) (out string) {

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
