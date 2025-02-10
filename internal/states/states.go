package states

import (
	"sync"
)

var (
	userStates = make(map[int64]string)
	mu         sync.RWMutex
)

func SetState(userID int64, state string) {
	mu.Lock()
	defer mu.Unlock()
	userStates[userID] = state
}

func GetState(userID int64) string {
	mu.RLock()
	defer mu.RUnlock()
	return userStates[userID]
}

func ClearState(userID int64) {
	mu.Lock()
	defer mu.Unlock()
	delete(userStates, userID)
}
