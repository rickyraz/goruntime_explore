package pointer

import (
	"fmt"
	"log"
)

func change(x *int) {
	*x = 100
}

func TestMutation() {
	n := 5
	fmt.Println("Sebelum:", n) // Output: 5

	// You are trying to pass an int, but the function expects a pointer to an int (*int).
	// “Hey, you gave me the number itself — but I want the address of that number!”

	// Salah: langsung passing nilai → tidak bisa, karena change() butuh *int
	// change(n) // ❌ error

	// Benar: pass alamat variabel menggunakan &
	change(&n)
	fmt.Println("Sesudah:", n) // Output: 100

	// ----

	// Immutable - value copy
	x := 5
	y := x
	y = 10 // x tetap 5

	// Mutable via pointer
	person := struct{ name string }{"Ricky"}
	log.Println("person1", person)
	p := &person
	p.name = "Adi" // person juga berubah

	log.Println("y", y)
	log.Println("person2", person)
	log.Println("p", p)
}
