// Copyright Huabing Zhao
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package debounce

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestDebouncer_debounceAfter(t *testing.T) {
	lock := sync.Mutex{}
	bounceNumer := 0
	stop := make(chan struct{})
	isCallbackCalled :=false
	callback := func() {
		lock.Lock()
		defer lock.Unlock()
		isCallbackCalled=true
		if bounceNumer != 5 {
			t.Errorf("test failed, expect: %v get: %v", 5, bounceNumer)
		} else {
			fmt.Printf("test succeed, expect: %v get: %v\n", 5, bounceNumer)
		}
	}
	d := New(500*time.Millisecond, 1*time.Second, callback, stop)
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		d.Bounce()
		lock.Lock()
		bounceNumer++
		lock.Unlock()
	}
	time.Sleep(600 * time.Millisecond)
	d.Bounce()
	if !isCallbackCalled{
		t.Errorf("test failed, callback is not called, bounceNumber: %v",  bounceNumer)
	}
	stop <- struct{}{}
}

func TestDebouncer_debounceMax(t *testing.T) {
	lock := sync.Mutex{}
	bounceNumer := 0
	stop := make(chan struct{})
	isCallbackCalled :=false
	callback := func() {
		lock.Lock()
		defer lock.Unlock()
		isCallbackCalled=true
		if bounceNumer != 11 {
			t.Errorf("test failed, expect: %v get: %v", 11, bounceNumer)
		} else {
			fmt.Printf("test succeed, expect: %v get: %v\n", 11, bounceNumer)
		}
	}
	d := New(500*time.Millisecond, 1*time.Second, callback, stop)
	for i := 0; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
		d.Bounce()
		lock.Lock()
		bounceNumer++
		lock.Unlock()
	}

	if !isCallbackCalled{
		t.Errorf("test failed, callback is not called, bounceNumber: %v",  bounceNumer)
	}
	stop <- struct{}{}
}

func TestDebouncer_debounceCancel(t *testing.T) {
	lock := sync.Mutex{}
	bounceNumer := 0
	stop := make(chan struct{})
	callback := func() {
		lock.Lock()
		defer lock.Unlock()
		if bounceNumer != 22 {
			t.Errorf("test failed, expect: %v get: %v", 20, bounceNumer)
		} else {
			fmt.Printf("test succeed, expect: %v get: %v\n", 22, bounceNumer)
		}
	}
	d := New(500*time.Millisecond, 1*time.Second, callback, stop)
	for i := 0; i < 25; i++ {
		time.Sleep(100 * time.Millisecond)
		d.Bounce()
		lock.Lock()
		bounceNumer++
		lock.Unlock()
		if i==10{
			d.Cancel()
		}
	}
	stop <- struct{}{}
}