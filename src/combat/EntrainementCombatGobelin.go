package combat

import (
	"fmt"
	ennemie "projetred/Monstre"
	"projetred/personnage"
)

func EntrainementGobelin(goblin ennemie.Monstre, joueur *personnage.Player) {
	for tour := 1; tour <= 10; tour++ {
		fmt.Printf("Tour %d :\n", tour)
		ennemie.GoblinPattern(goblin, joueur, tour)

		if goblin.Health == 0 {
			fmt.Println("Le goblin a été vaincu !")
			break
		}

		// Vérifier si le joueur est toujours en vie
		if joueur.Health == 0 {
			fmt.Println("Le joueur a été vaincu !")
			break
		}
	}
}

func charTurn(p *personnage.Player, m *ennemie.Monstre) {
	var choix int
	fmt.Println("Tour du joueur. Choisissez une action :")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")

	fmt.Scan(&choix)

	switch choix {
	case 1:
		EntrainementGobelin(m, &p)

		// Attaque basique
		fmt.Println(p.Pseudo, "utilise Attaque basique")
		degats := p.Attack
		m.Health -= degats
		if m.Health < 0 {
			m.Health = 0
		}
		fmt.Printf("%s inflige %d dégâts à %s\n", p.Pseudo, degats, m.Name)
		fmt.Printf("PV de %s : %d / %d\n", m.Name, m.Health, m.HealthMax)

	case 2:
		personnage.AccessInventory(&personnage.Player{})
	default:
		fmt.Println("Choix invalide.")
	}
}
