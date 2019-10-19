package lordeckoder

type Deck struct {
	Cards []Card
}

type Card struct {
	CardCode string
	Count    int
}

type Faction string

const (
	DEMACIA      Faction = "DE"
	FRELJORD     Faction = "FR"
	IONIA        Faction = "IO"
	NOXUS        Faction = "NX"
	PILTOVERZAUN Faction = "PZ"
	SHADOWISLES  Faction = "SI"
)

const MAX_CARD_COUNT = 3
