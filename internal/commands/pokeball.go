package commands

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
