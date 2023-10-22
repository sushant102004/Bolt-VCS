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
	_, err := os.Stat(".bolt")

	if os.IsNotExist(err) {
		if err := os.Mkdir(".bolt", 0750); err != nil {
			log.Fatal("error - ", err.Error())
			return
		}

		_, err := os.Create(".bolt/stage.txt")
		if err != nil {
			log.Fatal("error - ", err.Error())
			return
		}
		fmt.Println("Initialized")

	} else {
		fmt.Println("Bolt already initialized.")
	}

}
