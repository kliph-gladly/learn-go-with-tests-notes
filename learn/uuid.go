package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func main() {
	uuidWithHyphen, _ := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)
}
