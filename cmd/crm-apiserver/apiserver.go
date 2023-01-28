package main

import (
	_ "go.uber.org/automaxprocs"
	"math/rand"
	"time"

	"github.com/marmotedu/iam/internal/crmapiserver"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	crmapiserver.NewApp().Run()
}
