package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/tada-team/kozma"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(kozma.Say())
}
