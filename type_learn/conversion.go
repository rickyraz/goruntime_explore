package typelearn

import (
	"errors"
	"fmt"
	"strconv"
)

// Soal 1: String ke Int
// Convert string "42" ke int, print hasilnya
// Hint: gunakan strconv.Atoi() atau strconv.ParseInt()
func exercise1() (int, error) {
	str := "21"
	return strconv.Atoi(str)
}

// Soal 2: Int ke String
// Convert int 100 ke string, concatenate dengan "nilai: "
// Expected output: "nilai: 100"
// func exercise2(num int) (string, error) {
// 	intex := num

// 	// 	Semua nilai int valid secara representasi angka desimal,
// 	// Tidak ada kasus “gagal parse” seperti ketika mengubah string → int.
// 	s := strconv.Itoa(intex)

// 	result := fmt.Sprintf("nilai: %s", s)

// 	return result, nil
// }

func exercise2(num int) string {
	return fmt.Sprintf("nilai: %d", num)
}

// Soal 3: String ke Bytes
// Convert string "hello world" ke []byte, print length-nya
// Expected output: 11
func exercise3(str string) (int, error) {
	conv := []byte(str)

	return len(conv), nil
}

// Soal 4: Bytes ke String
// Convert []byte{72, 101, 108, 108, 111} ke string
// Expected output: "Hello"
func exercise4(byte []byte) string {
	str := string(byte)

	return str

}

// Soal 5: Float64 ke Int
// Convert float64 3.99 ke int, print hasilnya
// Expected output: 3 (decimal hilang)
func exercise5(numf float64) int {
	return int(numf)
}

// Soal 6: Int ke Float64
// Convert int 10 ke float64, divide dengan 3, print hasilnya
// Expected output: 3.333... atau similar
func exercise6(numI int) float64 {
	// TODO: implement
	return float64(numI) / 3.0
}

// Soal 7: String ke Float64
// Convert string "3.14" ke float64, multiply dengan 2
// Expected output: 6.28
func exercise7(str string) (float64, error) {
	nums, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return nums * 2, nil
}

// Soal 8: Interface{} ke String
// Diberikan var x interface{} = "golang"
// Type assert ke string, print hasilnya
// Expected output: "golang"

// interface{} adalah tipe “kosong” yang bisa menampung apa saja,
func exercise8(x interface{}) (string, error) {
	r, ok := x.(string)
	if !ok {
		return "", errors.New("x bukan string")
	}

	return r, nil
}

// Soal 9: Interface{} ke Int dengan error handling
// Diberikan var x interface{} = 42
// Type assert ke int, if gagal print error
func exercise9(x interface{}) (int, error) {
	r, ok := x.(int)
	if !ok {
		return 0, errors.New("x bukan int")
	}

	return r, nil
}

// Soal 10: Hex String ke Int
// Convert hex string "FF" ke int (255 dalam decimal)
// Hint: gunakan strconv.ParseInt() dengan base 16
// Expected output: 255
func exercise10(hx string) (int64, error) {
	n, err := strconv.ParseInt(hx, 16, 64) // base 16, 64-bit
	if err != nil {
		fmt.Println("error:", err)
		return n, err
	}

	return n, nil
}

func TypeConversion() {
	fmt.Println("=== Golang Type Casting Exercises ===")
	fmt.Println("\nSoal 1: String ke Int")
	ex1, err := exercise1()
	if err != nil {
		fmt.Println("ada error:", err)
	}
	fmt.Println("ex1: ", ex1)

	fmt.Println("\nSoal 2: Int ke String")
	ex2 := exercise2(100)
	fmt.Println("ex2: ", ex2)

	fmt.Println("\nSoal 3: String ke Bytes")
	ex3, err := exercise3("hello world")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("ex3: ", ex3)

	fmt.Println("\nSoal 4: Bytes ke String")
	ex4 := exercise4([]byte{72, 101, 108, 108, 111})
	fmt.Println("ex4: ", ex4)

	fmt.Println("\nSoal 5: Float64 ke Int")
	ex5 := exercise5(3.99)
	fmt.Println("ex5: ", ex5)

	fmt.Println("\nSoal 6: Int ke Float64")
	ex6 := exercise6(10)
	fmt.Println("ex6: ", ex6)

	fmt.Println("\nSoal 7: String ke Float64")
	ex7, err := exercise7("3.14")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("ex7: ", ex7)

	fmt.Println("\nSoal 8: Interface{} ke String")
	ex8, err := exercise8("golang")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("ex8: ", ex8)
	ex8b, err := exercise8(000)
	if err != nil {
		fmt.Println("Error:", err)
		// return
	}
	fmt.Println("ex8b: ", ex8b)

	fmt.Println("\nSoal 9: Interface{} ke Int dengan error handling")
	ex9a, err := exercise9(41)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("ex9a: ", ex9a)

	ex9b, err := exercise9("41")
	if err != nil {
		fmt.Println("Error:", err)
		// return
	}
	fmt.Println("ex9b: ", ex9b)

	fmt.Println("\nSoal 10: Hex String ke Int")
	ex10, err := exercise10("FF")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("ex10: ", ex10)
}
