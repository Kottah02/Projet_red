package main

import "fmt"
	"challenge"


func main() {
	// Initialisation du gobelin d'entraînement
	gobelin := InitGoblin()

	// Affichage des attributs du gobelin
	fmt.Println("Nom:", gobelin.Name)
	fmt.Println("Points de vie maximum:", gobelin.HealthMax)
	fmt.Println("Points de vie actuels:", gobelin.Health)
	fmt.Println("Points d'attaque:", gobelin.Attack)
}
