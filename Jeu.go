package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Structure du joueur
type Player struct {
	Name   string
	Health int
	Gold   int
	Wood   int
	Stone  int
	Leaf   int
}

// Structure du monstre
type Monster struct {
	Name   string
	Health int
	Damage int
}

// Fonction pour afficher le menu principal
func showMenu() {
	fmt.Println("Que souhaitez-vous faire ?")
	fmt.Println("1 - Combattre un monstre")
	fmt.Println("2 - Récolter des ressources (bois, pierre, feuille)")
	fmt.Println("3 - Aller voir le marchand pour vendre des ressources")
	fmt.Println("4 - Quitter le jeu")
}

// Fonction pour générer un monstre aléatoire
func generateMonster() Monster {
	monsters := []Monster{
		{"Gobelin", 20, 5},
		{"Orc", 30, 10},
		{"Dragon", 50, 20},
	}
	return monsters[rand.Intn(len(monsters))]
}

// Fonction de combat contre un monstre
func fight(player *Player) {
	monster := generateMonster()
	fmt.Printf("Un %s apparaît avec %d points de vie et peut infliger %d dégâts!\n", monster.Name, monster.Health, monster.Damage)

	for monster.Health > 0 && player.Health > 0 {
		// Combat tour par tour
		fmt.Println("Vous attaquez le monstre !")
		damage := rand.Intn(10) + 1
		monster.Health -= damage
		fmt.Printf("Vous infligez %d dégâts. Il reste %d points de vie au %s.\n", damage, monster.Health, monster.Name)

		if monster.Health > 0 {
			// Le monstre attaque le joueur
			fmt.Printf("Le %s vous attaque !\n", monster.Name)
			player.Health -= monster.Damage
			fmt.Printf("Vous perdez %d points de vie. Il vous reste %d points de vie.\n", monster.Damage, player.Health)
		}
	}

	if player.Health > 0 {
		fmt.Printf("Vous avez vaincu le %s et gagné 10 pièces d'or!\n", monster.Name)
		player.Gold += 10
	} else {
		fmt.Println("Vous êtes mort au combat. Le jeu est terminé.")
	}
}

// Fonction pour récolter des ressources
func gatherResources(player *Player) {
	wood := rand.Intn(5) + 1
	stone := rand.Intn(5) + 1
	leaf := rand.Intn(5) + 1

	player.Wood += wood
	player.Stone += stone
	player.Leaf += leaf

	fmt.Printf("Vous avez récolté %d bois, %d pierres, et %d feuilles.\n", wood, stone, leaf)
}

// Fonction pour vendre des ressources
func visitMerchant(player *Player) {
	fmt.Println("Bienvenue chez le marchand !")

	if player.Wood > 0 || player.Stone > 0 || player.Leaf > 0 {
		fmt.Printf("Vous avez %d bois, %d pierres et %d feuilles.\n", player.Wood, player.Stone, player.Leaf)
		fmt.Println("Le marchand vous propose 2 pièces d'or pour chaque unité de bois, pierre ou feuille.")

		totalGold := player.Wood*2 + player.Stone*2 + player.Leaf*2
		player.Gold += totalGold

		fmt.Printf("Vous avez vendu toutes vos ressources pour %d pièces d'or.\n", totalGold)

		// Réinitialiser les ressources
		player.Wood = 0
		player.Stone = 0
		player.Leaf = 0
	} else {
		fmt.Println("Vous n'avez pas de ressources à vendre.")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Initialisation du joueur
	player := Player{Name: "Héros", Health: 100, Gold: 0, Wood: 0, Stone: 0, Leaf: 0}

	fmt.Printf("Bienvenue dans le jeu, %s !\n", player.Name)

	// Boucle principale du jeu
	for player.Health > 0 {
		showMenu()
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fight(&player)
		case 2:
			gatherResources(&player)
		case 3:
			visitMerchant(&player)
		case 4:
			fmt.Println("Merci d'avoir joué ! À la prochaine.")
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}

		// Afficher l'état actuel du joueur
		fmt.Printf("Statut actuel - Vie: %d, Or: %d, Bois: %d, Pierre: %d, Feuilles: %d\n", player.Health, player.Gold, player.Wood, player.Stone, player.Leaf)
	}

	fmt.Println("Le jeu est terminé.")
}
