package lintexperiment

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	maxRandom := 100
	randomLimit := big.NewInt(int64(maxRandom))
	randomInt, _ := rand.Int(rand.Reader, randomLimit)
	fmt.Println(randomInt.Int64())
}
