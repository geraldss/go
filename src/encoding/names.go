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

const NAME_HASH_CAP = 1024 * 1024

var NAME_HASH = &NameHash{
	names: make(map[string]string, NAME_HASH_CAP),
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
		if len(this.names) < NAME_HASH_CAP {
			this.names[in] = in
		}
		out = in
	}
	this.Unlock()
	return
}
