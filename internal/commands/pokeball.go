package commands

import "fmt"

type Pokeball struct {
	Name            string
	CatchMultiplier float64
	GaurenteedCatch bool
}

var (
	StandardPokeball = Pokeball{
		Name:            "Poke Ball",
		CatchMultiplier: 1.0,
		GaurenteedCatch: false,
	}
	GreatBall = Pokeball{
		Name:            "Great Ball",
		CatchMultiplier: 1.5,
		GaurenteedCatch: false,
	}
	UltraBall = Pokeball{
		Name:            "Ultra Ball",
		CatchMultiplier: 2.0,
		GaurenteedCatch: false,
	}
	MasterBall = Pokeball{
		Name:            "Master Ball",
		GaurenteedCatch: true,
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
