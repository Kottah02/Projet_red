package main

type Waste struct {
	Name     string
	Quantity int
	Type     string
}

func main() {

	w := Waste{
		Name:     "Bottles",
		Quantity: 10,
		Type:     "Plastic",
	}

}
