package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Println(" Hello World 3 ")
	fmt.Printf(" uuid : %s", id)

}
