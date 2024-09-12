package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ANSI Color Codes
const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	cyan   = "\033[36m"
)

type Player struct {
	Class      string
	Health     int
	Wood       int
	Stone      int
	Leaf       int
	Sword      bool
	Bow        bool
	MagicStaff bool
	Potions    int // Consommables
}

func main() {
	fmt.Println("Bienvenue dans le Monde des Douze, les aventuriers cherchent les Dofus, des œufs de dragons aux pouvoirs immenses.")
	fmt.Println("Ces artefacts provoquent des conflits entre dieux, dragons et factions.")
	fmt.Println("La quête des Dofus primordiaux est centrale, influençant le destin du monde et entraînant des combats épiques.")
	fmt.Println(red + "LE MONDE COMPTE SUR TOI JEUNE AVENTURIER" + reset)
	reset := "\033[0m"
	fmt.Println("Appuyez sur Entrée pour continuer..." + reset)
	fmt.Scanln()

	Debut()
}
func Debut() {
	rand.Seed(time.Now().UnixNano())

	// Choisir une classe
	player := chooseClass()

	// Boucle principale du jeu
	for {

		fmt.Println("\n" + cyan + "================================================" + reset)
		fmt.Println("   Que voulez-vous faire ?")
		fmt.Println("================================================")
		fmt.Println(yellow + "1" + reset + " - Récolter des ressources")
		fmt.Println(yellow + "2" + reset + " - Combattre des monstres")
		fmt.Println(yellow + "3" + reset + " - Aller à l'établi pour construire des objets")
		fmt.Println(yellow + "4" + reset + " - Consulter votre inventaire")
		fmt.Println(yellow + "5" + reset + " - Quitter le jeu")
		fmt.Println(cyan + "================================================" + reset)

		var choice int
		fmt.Print("Choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			gatherResources(&player)
		case 2:
			combat(&player)
		case 3:
			craftItems(&player)
		case 4:
			showInventory(player)
		case 5:
			fmt.Println(green + "\nMerci d'avoir joué !" + reset)
			return
		default:
			fmt.Println(red + "Choix invalide." + reset)
		}
	}
}

// Fonction pour choisir la classe
func chooseClass() Player {
	fmt.Println(cyan + "================================================" + reset)
	fmt.Println("   Bienvenue dans le jeu !")
	fmt.Println("   Choisissez votre classe :")
	fmt.Println(cyan + "================================================" + reset)
	fmt.Println(yellow + "1" + reset + " - Guerrier")
	fmt.Println(yellow + "2" + reset + " - Archer")
	fmt.Println(yellow + "3" + reset + " - Mage")
	fmt.Println(cyan + "================================================" + reset)

	var classChoice int
	fmt.Print("Choix : ")
	fmt.Scan(&classChoice)

	player := Player{Health: 100, Potions: 3}

	switch classChoice {
	case 1:
		player.Class = "Guerrier"
		fmt.Println(green + "Vous avez choisi : Guerrier" + reset)
	case 2:
		player.Class = "Archer"
		fmt.Println(green + "Vous avez choisi : Archer" + reset)
	case 3:
		player.Class = "Mage"
		fmt.Println(green + "Vous avez choisi : Mage" + reset)
	default:
		fmt.Println(red + "Choix invalide, vous serez un Guerrier par défaut." + reset)
		player.Class = "Guerrier"
	}
	return player
}

// Fonction pour récolter des ressources
func gatherResources(player *Player) {
	fmt.Println(cyan + "\n================================================" + reset)
	fmt.Println("   Que voulez-vous récolter ?")
	fmt.Println(cyan + "================================================" + reset)
	fmt.Println(yellow + "1" + reset + " - Bois")
	fmt.Println(yellow + "2" + reset + " - Pierre")
	fmt.Println(yellow + "3" + reset + " - Feuilles")
	fmt.Println(cyan + "================================================" + reset)

	var resourceChoice int
	fmt.Print("Choix : ")
	fmt.Scan(&resourceChoice)

	switch resourceChoice {
	case 1:
		wood := rand.Intn(10) + 1
		player.Wood += wood
		fmt.Printf(green+"Vous avez récolté %d unités de bois. Total de bois : %d\n"+reset, wood, player.Wood)
	case 2:
		stone := rand.Intn(10) + 1
		player.Stone += stone
		fmt.Printf(green+"Vous avez récolté %d unités de pierre. Total de pierre : %d\n"+reset, stone, player.Stone)
	case 3:
		leaf := rand.Intn(10) + 1
		player.Leaf += leaf
		fmt.Printf(green+"Vous avez récolté %d feuilles. Total de feuilles : %d\n"+reset, leaf, player.Leaf)
	default:
		fmt.Println(red + "Choix invalide." + reset)
	}
}

// Fonction pour combattre des monstres
func combat(player *Player) {
	monsterHealth := rand.Intn(50) + 50
	fmt.Printf(cyan + "\n================================================\n" + reset)
	fmt.Printf("   Vous combattez un monstre avec %d points de vie !\n", monsterHealth)
	fmt.Println(cyan + "================================================" + reset)

	for monsterHealth > 0 && player.Health > 0 {
		// Attaque du joueur
		damage := rand.Intn(20) + 10
		monsterHealth -= damage
		fmt.Printf(green+"Vous attaquez et infligez %d de dégâts. Vie du monstre restante : %d\n"+reset, damage, monsterHealth)

		if monsterHealth <= 0 {
			fmt.Println(green + "Vous avez vaincu le monstre !" + reset)
			break
		}

		// Attaque du monstre
		monsterDamage := rand.Intn(15) + 5
		player.Health -= monsterDamage
		fmt.Printf(red+"Le monstre vous attaque et inflige %d de dégâts. Votre vie restante : %d\n"+reset, monsterDamage, player.Health)

		if player.Health <= 0 {
			fmt.Println(red + "Vous êtes mort..." + reset)
			return
		}
	}

	// Le joueur peut trouver une potion après le combat
	if rand.Float32() < 0.3 { // 30% de chance de trouver une potion
		player.Potions++
		fmt.Println(green + "Vous avez trouvé une potion !" + reset)
	}
}

// Fonction pour construire des objets
func craftItems(player *Player) {
	fmt.Println(cyan + "\n================================================" + reset)
	fmt.Println("   Que voulez-vous fabriquer ?")
	fmt.Println(cyan + "================================================" + reset)
	fmt.Println(yellow + "1" + reset + " - Épée (5 bois, 5 pierre)")
	fmt.Println(yellow + "2" + reset + " - Arc (5 bois, 5 feuilles)")
	fmt.Println(yellow + "3" + reset + " - Bâton magique (5 bois, 5 feuilles, 5 pierre)")
	fmt.Println(cyan + "================================================" + reset)

	var craftChoice int
	fmt.Print("Choix : ")
	fmt.Scan(&craftChoice)

	switch craftChoice {
	case 1:
		if player.Wood >= 5 && player.Stone >= 5 {
			player.Wood -= 5
			player.Stone -= 5
			player.Sword = true
			fmt.Println(green + "Vous avez fabriqué une épée." + reset)
		} else {
			fmt.Println(red + "Vous n'avez pas assez de ressources." + reset)
		}
	case 2:
		if player.Wood >= 5 && player.Leaf >= 5 {
			player.Wood -= 5
			player.Leaf -= 5
			player.Bow = true
			fmt.Println(green + "Vous avez fabriqué un arc." + reset)
		} else {
			fmt.Println(red + "Vous n'avez pas assez de ressources." + reset)
		}
	case 3:
		if player.Wood >= 5 && player.Leaf >= 5 && player.Stone >= 5 {
			player.Wood -= 5
			player.Leaf -= 5
			player.Stone -= 5
			player.MagicStaff = true
			fmt.Println(green + "Vous avez fabriqué un bâton magique." + reset)
		} else {
			fmt.Println(red + "Vous n'avez pas assez de ressources." + reset)
		}
	default:
		fmt.Println(red + "Choix invalide." + reset)
	}
}

// Fonction pour afficher l'inventaire
func showInventory(player Player) {
	fmt.Println(cyan + "\n=================== Inventaire ==================" + reset)

	// Objets
	fmt.Println(cyan + "[Objets]" + reset)
	if player.Sword {
		fmt.Println(green + "- Épée" + reset)
	}
	if player.Bow {
		fmt.Println(green + "- Arc" + reset)
	}
	if player.MagicStaff {
		fmt.Println(green + "- Bâton magique" + reset)
	}
	if !player.Sword && !player.Bow && !player.MagicStaff {
		fmt.Println(red + "Aucun objet." + reset)
	}

	// Ressources
	fmt.Println(cyan + "\n[Ressources]" + reset)
	fmt.Printf("- Bois : %d\n", player.Wood)
	fmt.Printf("- Pierre : %d\n", player.Stone)
	fmt.Printf("- Feuilles : %d\n", player.Leaf)

	// Consommables
	fmt.Println(cyan + "\n[Consommables]" + reset)
	fmt.Printf("- Potions : %d\n", player.Potions)
}
