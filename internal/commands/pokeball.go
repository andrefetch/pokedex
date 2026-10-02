package commands

import "fmt"

type Pokeball struct {
	Name            string
	CatchMultiplier float64
	GaurenteedCatch bool
	SpawnChance     float64
}

var (
	StandardPokeball = Pokeball{
		Name:            "Poke Ball",
		CatchMultiplier: 1.0,
		GaurenteedCatch: false,
		SpawnChance:     0.70,
	}
	GreatBall = Pokeball{
		Name:            "Great Ball",
		CatchMultiplier: 1.5,
		GaurenteedCatch: false,
		SpawnChance:     0.20,
	}
	UltraBall = Pokeball{
		Name:            "Ultra Ball",
		CatchMultiplier: 2.0,
		GaurenteedCatch: false,
		SpawnChance:     0.10,
	}
	MasterBall = Pokeball{
		Name:            "Master Ball",
		GaurenteedCatch: true,
		SpawnChance:     0.01,
	}
)

func CommandPokeballs(cfg *Config, args ...string) error {

	fmt.Println("Pokeballs Inventory:")

	ballTypes := map[string]Pokeball{
		"pokeball":   StandardPokeball,
		"greatball":  GreatBall,
		"ultraball":  UltraBall,
		"masterball": MasterBall,
	}

	for ballType, count := range cfg.PokeballTypes {
		fmt.Printf("%s: %d\n", ballTypes[ballType].Name, count)
	}
	return nil

}
