package personnage

import (
	"fmt"
	"math/rand"
)

// Fonction pour récolter des ressources
func GatherResources(player *Player) {
	for {
		fmt.Println(cyan + "\n================================================" + reset)
		fmt.Println("   Que voulez-vous récolter ?")
		fmt.Println(cyan + "================================================" + reset)
		fmt.Println(yellow + "1" + reset + " - Bois")
		fmt.Println(yellow + "2" + reset + " - Pierre")
		fmt.Println(yellow + "3" + reset + " - Feuilles")
		fmt.Println(yellow + "4" + reset + " - Fourrure de Loup")
		fmt.Println(yellow + "5" + reset + " - Peau de Troll")
		fmt.Println(yellow + "6" + reset + " - Cuir de Sanglier")
		fmt.Println(yellow + "7" + reset + " - Plume de Corbeau")
		fmt.Println(yellow + "0" + reset + " - Retour")
		fmt.Println(cyan + "================================================" + reset)

		var resourceChoice int
		fmt.Print("Choix : ")
		fmt.Scan(&resourceChoice)

		switch resourceChoice {
		case 0:
			// Revenir au menu principal
			return
		case 1:
			wood := rand.Intn(10) + 1
			player.Inventaire.Wood += wood
			fmt.Printf(green+"Vous avez récolté %d unités de bois. Total de bois : %d\n"+reset, wood, player.Inventaire.Wood)
		case 2:
			stone := rand.Intn(10) + 1
			player.Inventaire.Stone += stone
			fmt.Printf(green+"Vous avez récolté %d unités de pierre. Total de pierre : %d\n"+reset, stone, player.Inventaire.Stone)
		case 3:
			leaf := rand.Intn(10) + 1
			player.Inventaire.Leaf += leaf
			fmt.Printf(green+"Vous avez récolté %d feuilles. Total de feuilles : %d\n"+reset, leaf, player.Inventaire.Leaf)
		case 4:
			Fourrure := rand.Intn(10) + 1
			player.Inventaire.Fourrure += Fourrure
			fmt.Printf(green+"Vous avez récolté %d feuilles. Total de Fourrure de Loup : %d\n"+reset, Fourrure, player.Inventaire.Fourrure)

		case 5:
			PeauTroll := rand.Intn(10) + 1
			player.Inventaire.Peau_Troll += PeauTroll
			fmt.Printf(green+"Vous avez récolté %d feuilles. Total de Peau de Troll : %d\n"+reset, PeauTroll, player.Inventaire.Peau_Troll)
		case 6:
			CuirSanglier := rand.Intn(10) + 1
			player.Inventaire.CuirSanglier += CuirSanglier
			fmt.Printf(green+"Vous avez récolté %d feuilles. Total de Cuir de Sanglier : %d\n"+reset, CuirSanglier, player.Inventaire.CuirSanglier)
		case 7:
			Plume_de_Corbeau := rand.Intn(10) + 1
			player.Inventaire.PlumeCorbeau += Plume_de_Corbeau
			fmt.Printf(green+"Vous avez récolté %d feuilles. Total de Plume de Corbeau : %d\n"+reset, Plume_de_Corbeau, player.Inventaire.PlumeCorbeau)
		default:
			fmt.Println(red + "Choix invalide." + reset)
		}
	}
}
