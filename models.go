package lordeckoder

type Deck struct {
	Cards []Card
}

type Card struct {
	CardCode CardCode
	Count    int
}

type CardCode struct {
	Set     int
	Faction int
	Number  int
}

type Faction string

const (
	DEMACIA      Faction = "DE"
	FRELJORD     Faction = "FR"
	IONIA        Faction = "IO"
	NOXUS        Faction = "NX"
	PILTOVERZAUN Faction = "PZ"
	SHADOWISLES  Faction = "SI"
	UNKNOWN      Faction = "XX"
)

const MAX_CARD_COUNT = 3
