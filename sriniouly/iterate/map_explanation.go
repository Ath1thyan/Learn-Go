package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	fmt.Println("=== MAP DECLARATION vs INITIALIZATION ===\n")

	// ============================================
	// CONCEPT 1: Declaration vs Initialization
	// ============================================

	fmt.Println("1. DECLARATION ONLY (nil map):")
	var m1 map[string]Vertex
	fmt.Printf("   m1 = %v\n", m1)
	fmt.Printf("   m1 == nil? %v\n", m1 == nil)
	fmt.Printf("   Type: %T\n", m1)
	fmt.Println()

	// ============================================
	// CONCEPT 2: Why you CAN'T write to a nil map
	// ============================================

	fmt.Println("2. TRYING TO WRITE TO NIL MAP (will panic):")
	fmt.Println("   Uncomment the next line to see a panic:")
	// m1["test"] = Vertex{1.0, 2.0}  // ❌ PANIC: assignment to entry in nil map
	fmt.Println("   Error: assignment to entry in nil map\n")

	// ============================================
	// CONCEPT 3: Initialization with make()
	// ============================================

	fmt.Println("3. INITIALIZATION WITH make():")
	m2 := make(map[string]Vertex)
	fmt.Printf("   m2 = %v\n", m2)
	fmt.Printf("   m2 == nil? %v\n", m2 == nil)
	fmt.Println("   Now we can write to it:")
	m2["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Printf("   m2 = %v\n", m2)
	fmt.Println()

	// ============================================
	// CONCEPT 4: Different ways to initialize
	// ============================================

	fmt.Println("4. DIFFERENT INITIALIZATION METHODS:")

	// Method A: make() at declaration
	var m3 = make(map[string]Vertex)
	m3["A"] = Vertex{1.0, 2.0}
	fmt.Printf("   Method A (var m = make(...)): %v\n", m3)

	// Method B: map literal (empty)
	var m4 = map[string]Vertex{}
	m4["B"] = Vertex{3.0, 4.0}
	fmt.Printf("   Method B (var m = map[...]{}): %v\n", m4)

	// Method C: map literal with initial values
	var m5 = map[string]Vertex{
		"C": {5.0, 6.0},
	}
	fmt.Printf("   Method C (with initial values): %v\n", m5)

	// Method D: short declaration with make()
	m6 := make(map[string]Vertex)
	m6["D"] = Vertex{7.0, 8.0}
	fmt.Printf("   Method D (m := make(...)): %v\n", m6)
	fmt.Println()

	// ============================================
	// CONCEPT 5: Your original code explained
	// ============================================

	fmt.Println("5. YOUR ORIGINAL CODE EXPLAINED:")
	fmt.Println("   Line 12: var m map[string]Vertex")
	fmt.Println("            → This DECLARES m, but m is nil")
	fmt.Println("            → You CANNOT write to it yet!")
	fmt.Println()
	fmt.Println("   Line 47: m = make(map[string]Vertex)")
	fmt.Println("            → This INITIALIZES m (creates the actual map)")
	fmt.Println("            → Now you CAN write to it!")
	fmt.Println()
	fmt.Println("   Alternative: Initialize at declaration:")
	fmt.Println("   var m = make(map[string]Vertex)")
	fmt.Println("   → Now m is ready to use immediately!")
}
