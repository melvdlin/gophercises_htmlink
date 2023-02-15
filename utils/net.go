package utils

import (
	"net/http"
	"sync"
)

func LaunchNonBlocking(server *http.Server) (*sync.WaitGroup, <-chan error) {
	started := new(sync.WaitGroup)
	exitChan := make(chan error)
	started.Add(1)
	go func() {
		started.Done()
		exitChan <- server.ListenAndServe()
	}()
	return started, exitChan
}
