package lordeckoder

type Deck struct {
	Cards []CardInDeck
}

type CardInDeck struct {
	Card  Card
	Count int
}

type Card struct {
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
	BILGEWATER   Faction = "BW"
	UNKNOWN      Faction = "XX"
)

const MAX_CARD_COUNT = 3
