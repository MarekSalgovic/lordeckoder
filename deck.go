package lordeckoder

type Deck struct {
	Cards   []cardInDeck
	format  int
	version int
}

type cardInDeck struct {
	Card  Card
	Count int
}

func (deck Deck) GetVersion() int {
	maxVersion := 0
	var version int
	for _, card := range deck.Cards {
		switch card.Card.GetFaction() {
		case BILGEWATER, MOUNTTARGON:
			version = 2
		case SHURIMA:
			version = 3
		case BANDLECITY:
			version = 4
		default:
			version = 1
		}
		if version > maxVersion {
			maxVersion = version
		}
	}
	return maxVersion
}
