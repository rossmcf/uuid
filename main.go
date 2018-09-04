package main

import (
	"fmt"

	"github.com/satori/uuid"
)

func main() {
	fmt.Print(uuid.Must(uuid.NewV4()).String())
}
