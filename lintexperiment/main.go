package lintexperiment

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //nolint:staticcheck // this is a lint experiment
	fmt.Println(rand.Int())          //nolint:gosec // this is a lint experiment
}
