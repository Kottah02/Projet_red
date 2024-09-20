package personnage

import "fmt"

func MarchantMenu(player *Player) {
	fmt.Println(cyan + "\n=================== Marchand ==================" + reset)
	fmt.Println("1 - Acheter une Potion de vie (3 pièces d'or)")
	fmt.Println("2 - Acheter une Potion de poison (6 pièces d'or)")
	fmt.Println("3 - Acheter un livre de Compétence (25 pièces d'or)")
	fmt.Println("4 - Acheter une Fourrure de Loup (4 pièces d'or)")
	fmt.Println("5 - Acheter une Peau de Troll (7 pièces d'or)")
	fmt.Println("6 - Acheter un Cuir de Sanglier (3 pièces d'or)")
	fmt.Println("7 - Acheter une épée (100 pièces d'or)")
	fmt.Println("8 - Acheter une Arc (100 pièces d'or)")
	fmt.Println("9 - Acheter une Baguette Magique (100 pièces d'or)")
	fmt.Println("0 - Retour")
	fmt.Println(cyan + "================================================" + reset)

	var choice int
	fmt.Print("Choix : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		if player.Gold >= 3 {
			player.Inventaire.Potions++
			fmt.Println(green + "Vous avez acheté une Potion de vie." + reset)
			player.Gold -= 3
		} else {
			fmt.Println(red + "Vous n'avez pas assez de pièces d'or." + reset)
		}
	case 2:
		if player.Gold >= 6 {
			player.Inventaire.Potion_Poison++
			fmt.Println(green + "Vous avez acheté une Potion de poison." + reset)
			player.Gold -= 6
		} else {
			fmt.Println(red + "Vous n'avez pas assez de pièces d'or." + reset)
		}
	case 3:
		if player.Gold >= 25 {
			player.Inventaire.SpellBookCount++
			fmt.Println(green + "Vous avez acheté un Livre de compétence : Boule de Feu." + reset)
			player.Gold -= 25
		} else {
			fmt.Println(red + "Vous n'avez pas assez de pièces d'or." + reset)
		}
	case 4:
		if player.Gold >= 4 {
			player.Inventaire.Fourrure++
			fmt.Println(green + "Vous avez acheté une Fourrure de Loup." + reset)
			player.Gold -= 4
		} else {
			fmt.Println(red + "Vous n'avez pas assez de pièces d'or." + reset)
		}
	case 5:
		if player.Gold >= 7 {
			player.Inventaire.Peau_Troll++
			fmt.Println(green + "Vous avez acheté une Peau de Troll." + reset)
			player.Gold -= 7
		} else {
			fmt.Println(red + "Vous n'avez pas assez de pièces d'or." + reset)
		}
	case 6:
		if player.Gold >= 3 {
			player.Inventaire.CuirSanglier++
			fmt.Println(green + "Vous avez acheté un Cuir de Sanglier." + reset)
			player.Gold -= 3
		} else {
			fmt.Println(red + "Vous n'avez pas assez de pièces d'or." + reset)
		}
	case 7:
		if player.Gold >= 100 {
			player.Inventaire.Sword++
			fmt.Println(green + "Vous avez acheté une épée." + reset)
			player.Gold -= 100
		} else {
			fmt.Println(red + "Vous n'avez pas assez de pièces d'or." + reset)
		}
	case 8:
		if player.Gold >= 100 {
			player.Inventaire.Bow++
			fmt.Println(green + "Vous avez acheté un arc." + reset)
			player.Gold -= 100
		} else {
			fmt.Println(red + "Vous n'avez pas assez de pièces d'or." + reset)
		}
	case 9:
		if player.Gold >= 100 {
			player.Inventaire.MagicStaff++
			fmt.Println(green + "Vous avez acheté une baguette magique." + reset)
			player.Gold -= 100
		} else {
			fmt.Println(red + "Vous n'avez pas assez de pièces d'or." + reset)
		}
	case 0:
		return
	default:
		fmt.Println(red + "Choix invalide." + reset)
	}
}
