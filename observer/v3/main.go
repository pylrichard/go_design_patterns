package main

func main() {
	shirtItem := NewItem("Nike Shirt")
	obFirst := &Customer{ Id: "abc"}
	obSecond := &Customer{ Id: "xyz"}

	shirtItem.Register(obFirst)
	shirtItem.Register(obSecond)
	shirtItem.UpdateAvailability()
}