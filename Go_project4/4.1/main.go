package main

// GOSSAFUNC=main go tool compile ./main.go > ssa.html
func main() {
	a := 1
	b := 2
	if true {
		add(a, b)
	}
}

func add(a, b int) {
	println(a + b)
}
