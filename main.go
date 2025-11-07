package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/shopspring/decimal"
)

type Stats struct {
	Goroutines int
	Memory     uint64
}

var balance = decimal.NewFromFloat(1000.00)

func main() {
	// Goroutine 1
	go func() {
		newBalance := balance.Add(decimal.NewFromFloat(100))
		fmt.Println(newBalance) // 1100
	}()

	// Goroutine 2
	go func() {
		newBalance := balance.Sub(decimal.NewFromFloat(50))
		fmt.Println(newBalance) // 950
	}()

	// ----

	fmt.Println(balance) // 1000 (TIDAK BERUBAH)
	// ✅ Tidak ada race condition

	fmt.Println("=== Go Runtime Demo ==")

	// // Spawn goroutines
	// for i := 0; i < 5; i++ {
	// 	go worker(i)
	// }

	// // Monitor runtime stats
	// for j := 0; j < 10; j++ {
	// 	stats := getStats()
	// 	fmt.Printf("Tick %d: Goroutines=%d, Memory=%dKB\n",
	// 		j, stats.Goroutines, stats.Memory/1024)
	// 	time.Sleep(500 * time.Millisecond)
	// }
}

func worker(_ int) {
	for {
		// Allocate memory (trigger GC)
		data := make([]byte, 1024*100) // 100KB
		_ = data
		time.Sleep(200 * time.Millisecond)
	}
}

func getStats() Stats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return Stats{
		Goroutines: runtime.NumGoroutine(),
		Memory:     m.Alloc,
	}
}

// **Output:**
// ```
// === Go Runtime Demo ===

// Tick 0: Goroutines=6, Memory=482KB   ← Start
// Tick 1: Goroutines=6, Memory=1731KB  ← Workers allocate memory
// Tick 2: Goroutines=6, Memory=3083KB  ← Memory keeps growing
// Tick 3: Goroutines=6, Memory=191KB   ← GC TRIGGERED! Memory cleaned
// Tick 4: Goroutines=6, Memory=1233KB  ← Allocate again
// Tick 5: Goroutines=6, Memory=2793KB  ← Growing...
// Tick 6: Goroutines=6, Memory=3834KB  ← Peak
// Tick 7: Goroutines=6, Memory=1652KB  ← GC cleaned some
// Tick 8: Goroutines=6, Memory=2694KB  ← Growing again
// Tick 9: Goroutines=6, Memory=508KB   ← GC TRIGGERED again

// -------------
