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
	Pseudo     string
	Sex        string
	Class      string
	Level      int
	HealthMax  int
	Health     int
	Inventaire Objet
	Skills     []string
	Gold       int
	Equip      Equipment
	Attack     int
}

type Objet struct {
	Wood           int
	Stone          int
	Leaf           int
	Potions        int
	Potion_Poison  int
	Sword          int
	Bow            int
	MagicStaff     int
	SpellBookCount int
	Fourrure       int
	Peau_Troll     int
	CuirSanglier   int
	PlumeCorbeau   int
	ChapAven       int
	TunAven        int
	BotteAven      int
}

type Equipment struct {
	Head  string
	Torso string
	Feet  string
}

func main() {
	fmt.Println("Bienvenue dans le Monde des Douze, les aventuriers cherchent les Dofus, des œufs de dragons aux pouvoirs immenses.")
	fmt.Println("Ces artefacts provoquent des conflits entre dieux, dragons et factions.")
	fmt.Println("La quête des Dofus primordiaux est centrale, influençant le destin du monde et entraînant des combats épiques.")
	fmt.Println(red + "LE MONDE COMPTE SUR TOI JEUNE AVENTURIER" + reset)
	fmt.Println("Appuyez sur Entrée pour continuer..." + reset)
	fmt.Scanln()

	Debut()
}

func Debut() {
	rand.Seed(time.Now().UnixNano())

	// Choisir une classe
	player := createCharacter()

	// Boucle principale du jeu
	for {

		fmt.Println("\n" + cyan + "================================================" + reset)
		fmt.Println("   Que voulez-vous faire ?")
		fmt.Println("================================================")
		fmt.Println(yellow + "1" + reset + " - Afficher l'Information du personnage")
		fmt.Println(yellow + "2" + reset + " - Récolter des ressources")
		fmt.Println(yellow + "3" + reset + " - Combattre des monstres")
		fmt.Println(yellow + "4" + reset + " - Forgeron")
		fmt.Println(yellow + "5" + reset + " - Consulter votre inventaire")
		fmt.Println(yellow + "6" + reset + " - Accéder au Marchand")
		fmt.Println(yellow + "7" + reset + " - Quitter le jeu")
		fmt.Println(cyan + "================================================" + reset)

		var choice int
		fmt.Print("Choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			DisplayInfo(player)
		case 2:
			gatherResources(&player)
		case 3:
			combat(&player)
		case 4:
			craftItems(&player)
		case 5:
			accessInventory(&player)
		case 6:
			marchantMenu(&player)
		case 7:
			fmt.Println(green + "\nMerci d'avoir joué !" + reset)
			return
		default:
			fmt.Println(red + "Choix invalide." + reset)
		}
	}
}

func DisplayInfo(player Player) {
	fmt.Printf("Pseudo : %s\nSexe : %s\nClasse : %s\nNiveau : %d\nVie actuelle : %d\nVie Max : %d\nPièce d'or : %d\n", player.Pseudo, player.Sex, player.Class, player.Level, player.Health, player.HealthMax, player.Gold)
}

// Fonction pour créer le personnage (pseudo, sexe, classe)
func createCharacter() Player {
	var player Player

	// Choisir le pseudo
	fmt.Println(cyan + "================================================" + reset)
	fmt.Println("   Entrez votre pseudo :")
	fmt.Println(cyan + "================================================" + reset)
	fmt.Print("Pseudo : ")
	fmt.Scan(&player.Pseudo)

	// Choisir le sexe
	fmt.Println(cyan + "\n================================================" + reset)
	fmt.Println("   Choisissez votre sexe :")
	fmt.Println(cyan + "================================================" + reset)
	fmt.Println(yellow + "1" + reset + " - Homme")
	fmt.Println(yellow + "2" + reset + " - Femme")
	fmt.Println(cyan + "================================================" + reset)

	var sexChoice int
	for {
		fmt.Print("Choix : ")
		fmt.Scan(&sexChoice)
		if sexChoice == 1 || sexChoice == 2 {
			break
		}
		fmt.Println(red + "Choix invalide. Veuillez entrer 1 ou 2." + reset)
	}

	switch sexChoice {
	case 1:
		player.Sex = "Homme"
	case 2:
		player.Sex = "Femme"
	}

	// Choisir la classe
	fmt.Println(cyan + "\n================================================" + reset)
	fmt.Println("   Choisissez votre classe :")
	fmt.Println(cyan + "================================================" + reset)
	fmt.Println(yellow + "1" + reset + " - Guerrier")
	fmt.Println(yellow + "2" + reset + " - Archer")
	fmt.Println(yellow + "3" + reset + " - Mage")
	fmt.Println(yellow + "4" + reset + " - Elfe")
	fmt.Println(cyan + "================================================" + reset)

	var classChoice int
	fmt.Print("Choix : ")
	fmt.Scan(&classChoice)

	switch classChoice {
	case 1:
		player.Class = "Guerrier"
		fmt.Println(green + "Vous avez choisi : Guerrier" + reset)
		player.HealthMax = 200
		player.Health = 200
		player.Inventaire.Potions = 3
		player.Level = 1
		player.Gold = 100
		player.Attack = 10
	case 2:
		player.Class = "Archer"
		fmt.Println(green + "Vous avez choisi : Archer" + reset)
		player.HealthMax = 100
		player.Health = 100
		player.Inventaire.Potions = 3
		player.Level = 1
		player.Gold = 100
		player.Attack = 30

	case 3:
		player.Class = "Mage"
		fmt.Println(green + "Vous avez choisi : Mage" + reset)
		player.HealthMax = 150
		player.Health = 250
		player.Inventaire.Potions = 3
		player.Level = 1
		player.Gold = 100
		player.Attack = 20
	case 4:
		player.Class = "Elfe"
		fmt.Println(green + "Vous avez choisi : Elfe" + reset)
		player.HealthMax = 100
		player.Health = 40
		player.Inventaire.Potions = 3
		player.Level = 1
		player.Gold = 100
		player.Attack = 10
	default:
		fmt.Println(red + "Choix invalide, vous serez un Guerrier par défaut." + reset)
		player.Class = "Guerrier"
	}

	player.Skills = append(player.Skills, "Coup de poing")
	fmt.Println(green + "Vous avez appris une nouvelle compétence : Coup de poing" + reset)
	return player
}

// Fonction pour récolter des ressources
func gatherResources(player *Player) {
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

// Fonction pour combattre des monstres
func combat(player *Player) {
	c := 0
	monsterHealth := rand.Intn(50) + 50
	for monsterHealth > 0 || player.Health > 0 {
		// Attaque du joueur
		fmt.Printf(cyan + "\n================================================\n" + reset)
		fmt.Printf("   Vous allez se combattre avec %d de points de vie\n", player.Health)
		fmt.Println(cyan + "================================================" + reset)

		fmt.Println("Veux-tu prendre une potion avant d'aller au combat" + reset)
		fmt.Printf(yellow + "1" + reset + " - Prendre une potion\n")
		fmt.Printf(yellow + "2" + reset + " - Affronter un monstre\n")
		fmt.Println(yellow + "0" + reset + " - Retour")

		var choice int
		fmt.Print("Choix : ")
		fmt.Scan(&choice)
		fmt.Println(cyan + "================================================" + reset)
		switch choice {
		case 1:
			takePot(player)
		case 2:
			if monsterHealth > 0 {
				c += 1
				fmt.Printf("Tour : %d\n", c)
				damage := player.Attack
				monsterHealth -= damage
				fmt.Printf(cyan + "\n===================================================\n" + reset)
				fmt.Printf("   Vous allez se combattre avec %d de points de vie\n", player.Health)
				fmt.Println(cyan + "===================================================" + reset)
				fmt.Printf(green+"Vous attaquez et infligez %d de dégâts. Vie du monstre restante : %d\n"+reset, damage, monsterHealth)

				// Attaque du monstre
				if monsterHealth > 0 {
					monsterDamage := rand.Intn(15) + 5
					player.Health -= monsterDamage
					fmt.Printf(red+"Le monstre vous attaque et inflige %d de dégâts.\n"+reset, monsterDamage)
				} else {
					fmt.Println(green + "Le monstre est vaincu !")
					fmt.Printf("Il vous reste %d de points de vie\n", player.Health)
				}
				if player.Health < 0 {
					fmt.Println(red + "Vous êtes mort...\n" + reset)
					player.Health = player.HealthMax / 2
					fmt.Printf(green+"Vous avez été ressuscité avec %d points de vie.\n"+reset, player.Health)
					// Le joueur peut trouver une potion après le combat
				}
				if rand.Float32() < 0.3 { // 30% de chance de trouver une potion
					player.Inventaire.Potions++
					fmt.Println(green + "Vous avez trouvé une potion !\n" + reset)

				}
			} else {
				fmt.Println(green + "Il n'y a plus de monstre.\n" + reset)
			}
		case 0:
			return
		default:
			fmt.Println(red + "Choix invalide." + reset)
		}
	}
}

// Fonction pour construire des objets
func craftItems(player *Player) {
	for {
		fmt.Println(cyan + "\n================================================" + reset)
		fmt.Println("   Que voulez-vous fabriquer ?")
		fmt.Println(cyan + "================================================" + reset)
		fmt.Println(yellow + "1" + reset + " - Épée (5 bois, 5 pierre)")
		fmt.Println(yellow + "2" + reset + " - Arc (5 bois, 5 feuilles)")
		fmt.Println(yellow + "3" + reset + " - Bâton magique (5 bois, 5 feuilles, 5 pierre)")
		fmt.Println(yellow + "4" + reset + " - Chapeau de l'aventurier (1 Plume de Corbeau, 1 Cuir de Sanglier)")
		fmt.Println(yellow + "5" + reset + " - Tunique de l'aventurier (2 Fourrures de Loup, 1 Peau de troll)")
		fmt.Println(yellow + "6" + reset + " - Bottes de l'aventurier (1 Fourrure de Loup, 1 Cuir de Sanglier)")
		fmt.Println(yellow + "0" + reset + " - Retour")
		fmt.Println(cyan + "================================================" + reset)

		var craftChoice int
		fmt.Print("Choix : ")
		fmt.Scan(&craftChoice)

		switch craftChoice {
		case 0:
			// Revenir au menu principal
			return
		case 1:
			if player.Inventaire.Wood >= 5 && player.Inventaire.Stone >= 5 && player.Gold >= 5 {
				player.Gold -= 5
				player.Inventaire.Wood -= 5
				player.Inventaire.Stone -= 5
				player.Inventaire.Sword++
				fmt.Println(green + "Vous avez fabriqué une épée." + reset)
			} else {
				fmt.Println(red + "Vous n'avez pas assez de ressources." + reset)
			}
		case 2:
			if player.Inventaire.Wood >= 5 && player.Inventaire.Leaf >= 5 && player.Gold >= 5 {
				player.Gold -= 5
				player.Inventaire.Wood -= 5
				player.Inventaire.Leaf -= 5
				player.Inventaire.Bow++
				fmt.Println(green + "Vous avez fabriqué un arc." + reset)
			} else {
				fmt.Println(red + "Vous n'avez pas assez de ressources." + reset)
			}
		case 3:
			if player.Inventaire.Wood >= 5 && player.Inventaire.Leaf >= 5 && player.Inventaire.Stone >= 5 && player.Gold >= 5 {
				player.Gold -= 5
				player.Inventaire.Wood -= 5
				player.Inventaire.Leaf -= 5
				player.Inventaire.Stone -= 5
				player.Inventaire.MagicStaff++
				fmt.Println(green + "Vous avez fabriqué un bâton magique." + reset)
			} else {
				fmt.Println(red + "Vous n'avez pas assez de ressources." + reset)
			}
		case 4:
			if player.Gold >= 5 && player.Inventaire.PlumeCorbeau >= 1 && player.Inventaire.CuirSanglier >= 1 {
				player.Gold -= 5
				player.Inventaire.PlumeCorbeau -= 1
				player.Inventaire.CuirSanglier -= 1
				player.Inventaire.ChapAven++
				fmt.Println(green + "Vous avez fabriqué un Chapeau de l'aventurier." + reset)
			} else {
				fmt.Println(red + "Vous n'avez pas assez de ressources." + reset)
			}
		case 5:
			if player.Gold >= 5 && player.Inventaire.Fourrure >= 2 && player.Inventaire.Peau_Troll >= 1 {
				player.Gold -= 5
				player.Inventaire.Fourrure -= 2
				player.Inventaire.Peau_Troll -= 1
				player.Inventaire.TunAven++
				fmt.Println(green + "Vous avez fabriqué un Tunique de l'aventurier." + reset)
			} else {
				fmt.Println(red + "Vous n'avez pas assez de ressources." + reset)
			}
		case 6:
			if player.Gold >= 5 && player.Inventaire.Fourrure >= 1 && player.Inventaire.CuirSanglier >= 1 {
				player.Gold -= 5
				player.Inventaire.Fourrure -= 1
				player.Inventaire.CuirSanglier -= 1
				player.Inventaire.BotteAven++
				fmt.Println(green + "Vous avez fabriqué une paire de bottes l'aventurier." + reset)
			} else {
				fmt.Println(red + "Vous n'avez pas assez de ressources." + reset)
			}
		default:
			fmt.Println(red + "Choix invalide." + reset)
		}
	}
}

// Tâche 6
func takePot(player *Player) {
	if player.Health == player.HealthMax {
		fmt.Println(yellow + "Votre vie est déjà au maximum !" + reset)
		return
	}

	if player.Inventaire.Potions > 0 {
		healAmount := 50
		actualHeal := player.HealthMax - player.Health
		if healAmount > actualHeal {
			healAmount = actualHeal
		}
		player.Health += healAmount
		player.Inventaire.Potions--
		fmt.Printf(green+"Vous avez utilisé une potion et regagnez %d points de vie.\n"+reset, healAmount)
		fmt.Printf("Votre vie actuelle est de %d / %d\n", player.Health, player.HealthMax)
	} else {
		fmt.Println(red + "Vous n'avez plus de potions !" + reset)
	}
}

// Fonction pour afficher l'inventaire

func accessInventory(player *Player) {
	fmt.Println(cyan + "\n=================== Inventaire ==================" + reset)

	// Information du joueur
	fmt.Printf("Pseudo : %s | Sexe : %s | Classe : %s Vie Max : %d | Vie Actuelle : %d | Niveau : %d | Pièces d'or : %d\n ", player.Pseudo, player.Sex, player.Class, player.HealthMax, player.Health, player.Level, player.Gold)

	// Objets
	fmt.Println(cyan + "\n[Equipement]" + reset)

	if player.Inventaire.Bow > 0 {
		fmt.Printf(green+"- Arc (%d)\n"+reset, player.Inventaire.Bow)
	}
	if player.Inventaire.MagicStaff > 0 {
		fmt.Printf(green+"- Bâton magique (%d)\n"+reset, player.Inventaire.MagicStaff)
	}

	fmt.Println(cyan + "\n[Objets]" + reset)
	if player.Inventaire.Sword > 0 {
		fmt.Printf(green+"- Épée (%d)\n"+reset, player.Inventaire.Sword)
	}
	if player.Inventaire.Bow > 0 {
		fmt.Printf(green+"- Arc (%d)\n"+reset, player.Inventaire.Bow)
	}
	if player.Inventaire.MagicStaff > 0 {
		fmt.Printf(green+"- Bâton magique (%d)\n"+reset, player.Inventaire.MagicStaff)
	}
	if player.Inventaire.Potion_Poison > 0 {
		fmt.Printf(green+"- Potions de Poison (%d)\n"+reset, player.Inventaire.Potion_Poison)
	}
	if player.Inventaire.ChapAven > 0 {
		fmt.Printf(green+"- Chapeau de l'aventurier (%d)\n"+reset, player.Inventaire.ChapAven)
	}
	if player.Inventaire.TunAven > 0 {
		fmt.Printf(green+"- Tunique de l'aventurier (%d)\n"+reset, player.Inventaire.TunAven)
	}
	if player.Inventaire.BotteAven > 0 {
		fmt.Printf(green+"- Bottes de l'aventurier (%d)\n"+reset, player.Inventaire.BotteAven)

	}
	if player.Inventaire.Sword == 0 && player.Inventaire.Bow == 0 && player.Inventaire.MagicStaff == 0 && player.Inventaire.Potion_Poison == 0 && player.Inventaire.ChapAven == 0 && player.Inventaire.TunAven == 0 && player.Inventaire.BotteAven == 0 {
		fmt.Println(red + "Aucun arme." + reset)
	}
	fmt.Println(cyan + "\n[Compétences]" + reset)
	for _, skill := range player.Skills {
		fmt.Println(green + "- " + skill + reset)
	}

	// Ressources
	fmt.Println(cyan + "\n[Ressources]" + reset)
	fmt.Printf("- Bois : %d\n", player.Inventaire.Wood)
	fmt.Printf("- Pierre : %d\n", player.Inventaire.Stone)
	fmt.Printf("- Feuilles : %d\n", player.Inventaire.Leaf)

	// Consommables
	fmt.Println(cyan + "\n[Consommables]" + reset)
	fmt.Printf("- Potions : %d\n", player.Inventaire.Potions)

	fmt.Println("\n" + cyan + "================================================" + reset)
	fmt.Println("   Que voulez-vous faire ?")
	fmt.Println("================================================")
	fmt.Println(yellow + "1" + reset + " - Prendre une potion")
	fmt.Println(yellow + "2" + reset + " - Utiliser le Livre de compétence : Boule de Feu")
	fmt.Println(yellow + "0" + reset + " - Retour")

	var choice int
	fmt.Print("Choix : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		takePot(player)
	case 2:
		if player.Inventaire.SpellBookCount > 0 {
			spellBook(player)
			player.Inventaire.SpellBookCount--
		} else {
			fmt.Println(red + "Vous n'avez pas de Livre de compétence : Boule de Feu." + reset)
		}
	case 0:
		return
	default:
		fmt.Println(red + "Choix invalide." + reset)

	}

}

// Tâche 7
func marchantMenu(player *Player) {
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
			addInventory(player, "potion", 1)
			fmt.Println(green + "Vous avez acheté une Potion de vie." + reset)
			player.Gold -= 3
		} else {
			fmt.Println(red + "Vous n'avez pas assez de pièces d'or." + reset)
		}
	case 2:
		if player.Gold >= 6 {
			addInventory(player, "potionPoison", 1)
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

func addInventory(player *Player, item string, quantity int) {
	switch item {
	case "potion":
		player.Inventaire.Potions += quantity
	case "potionPoison":
		player.Inventaire.Potion_Poison += quantity
	case "sword":
		player.Inventaire.Sword += quantity
	case "bow":
		player.Inventaire.Bow += quantity
	case "magicstaff":
		player.Inventaire.MagicStaff += quantity
	}
}

// Tâche 10
func spellBook(player *Player) {
	for _, skill := range player.Skills {
		if skill == "Boule de Feu" {
			fmt.Println(red + "Vous avez déjà appris la compétence : Boule de Feu." + reset)
			return
		}
	}
	player.Skills = append(player.Skills, "Boule de Feu")
	fmt.Println(green + "Vous avez appris une nouvelle compétence : Boule de Feu." + reset)
}
