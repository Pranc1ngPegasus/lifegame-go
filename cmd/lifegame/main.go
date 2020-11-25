package main

import (
	"time"

	"github.com/Pranc1ngPegasus/lifegame-go/internal/domain"
)

const (
	sleepTime = 100 * time.Millisecond
)

func main() {
	lg := domain.NewLifegame()
	lg.Initialize()

	for {
		lg.Render()
		lg.Update()
		time.Sleep(sleepTime)
	}
}
