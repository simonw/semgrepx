package bad

func Bad() {
	panic("do not call me")
}

func Good() {
	println("hello")
}

func main() {
	Bad()
	Good()
	Bad()
}
