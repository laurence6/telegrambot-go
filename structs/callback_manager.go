package structs

import (
	"sync"

	"github.com/laurence6/telegrambot-go/handlers"
)

// CallbackManager is a thread safe manager for HandleFunc.
type CallbackManager struct {
	MaxLength int

	callbacks        map[string]handlers.HandleFunc
	callbacksKeyList []string

	*sync.Mutex
}

// NewCallbackManager returns a CallbackManager instance with max length = 0.
func NewCallbackManager() *CallbackManager {
	return NewCallbackManagerWithMaxLength(0)
}

// NewCallbackManagerWithMaxLength returns a CallbackManager instance with specified max length.
func NewCallbackManagerWithMaxLength(maxLength int) *CallbackManager {
	return &CallbackManager{
		maxLength,
		map[string]handlers.HandleFunc{},
		[]string{},
		&sync.Mutex{},
	}
}

// Put stores the callback HandleFunc using a string as key.
func (callbackManager *CallbackManager) Put(id string, handleFunc handlers.HandleFunc) {
	callbackManager.Lock()
	callbackManager.callbacks[id] = handleFunc
	callbackManager.callbacksKeyList = append(callbackManager.callbacksKeyList, id)
	callbackManager.Unlock()

	callbackManager.GC()
}

// Get returns the HandleFunc using the specified key.
func (callbackManager *CallbackManager) Get(id string) handlers.HandleFunc {
	callbackManager.Lock()
	defer callbackManager.Unlock()
	handleFunc, ok := callbackManager.callbacks[id]
	if ok {
		delete(callbackManager.callbacks, id)
		for n, i := range callbackManager.callbacksKeyList {
			if i == id {
				callbackManager.callbacksKeyList = append(callbackManager.callbacksKeyList[:n], callbackManager.callbacksKeyList[n+1:]...)
				break
			}
		}
		return handleFunc
	}

	return nil
}

// RemoveFirst removes the first n HandleFunc from the list.
func (callbackManager *CallbackManager) RemoveFirst(n int) {
	callbackManager.Lock()
	if length := len(callbackManager.callbacksKeyList); length < n {
		n = length
	}
	for i := 0; i < n; i++ {
		delete(callbackManager.callbacks, callbackManager.callbacksKeyList[i])
	}
	callbackManager.callbacksKeyList = callbackManager.callbacksKeyList[n:]
	callbackManager.Unlock()
}

// GC removes the oldest HandleFunc from the list to make sure that the length of list is equal to or smaller than max length.
func (callbackManager *CallbackManager) GC() {
	if callbackManager.MaxLength == 0 {
		return
	}
	length := len(callbackManager.callbacksKeyList)
	if length <= callbackManager.MaxLength {
		return
	}

	callbackManager.RemoveFirst(length - callbackManager.MaxLength)
}
