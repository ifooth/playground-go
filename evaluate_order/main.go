package main

func main() {
	n := 0

	f := func() func(int, int) {
		n = 1
		return func(int, int) {}
	}

	g := func() (int, int) {
		println(n)
		return 0, 0
	}

	f()(g())
}
