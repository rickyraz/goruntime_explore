package main

import (
	errhandle "goruntime_explore/err_handle"
	httpnative "goruntime_explore/http_native"
	"goruntime_explore/params"
	"goruntime_explore/pointer"
	typelearn "goruntime_explore/type_learn"
	"log"
)

// Main
func main() {
	typelearn.TypeConversion()
	params.Variadic()

	err := errhandle.HandleRequest()
	if err != nil {
		// log.Fatal(err)
		log.Println("Error:", err)
	}

	data, err := errhandle.ReadFile("config.json")
	if err != nil {
		// log.Fatal("err in readFile: ", err)
		log.Println("Error:", err)
	}
	log.Println("data:", data)

	pointer.TestMutation()

	httpnative.RunHTTPNative()
}
