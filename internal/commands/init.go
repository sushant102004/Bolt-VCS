/*
	PoF: - Provide functionality to initialize .bolt folder.
*/

package internal

import (
	"fmt"
	"log"
	"os"
)

func Initialize() {
	if err := os.Mkdir(".bolt", 0750); err != nil {
		log.Fatal("error - ", err.Error())
		return
	}
	fmt.Println("Initialized")
}
