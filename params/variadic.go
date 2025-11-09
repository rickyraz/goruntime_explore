package params

import (
	"log"
	"strings"
)

func Variadic() {

	log.Println(greet())              // hi there
	log.Println(greet("Rick"))        // hi Rick
	log.Println(greet("Rick", "Amy")) // hi Rick (ambil argumen pertama)
}

func greet(name ...string) string {
	if len(name) == 0 {
		return "hi there"
	}
	return "hi there " + strings.Join(name, ", ")
}
