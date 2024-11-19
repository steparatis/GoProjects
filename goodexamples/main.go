package main

func main() {
	count := 0
	for i := range [256]struct{}{} {
		n := uint8(i)
		println("n", n)
		println("-n", -n)
	}
	println(count)
}
