package lordeckoder

import (
	"encoding/base32"
)

func fixDeckcodeLength(dc string) (string) {
	length := len(dc)
	if (length%8 != 0) {
		for i := 0; i < length%8; i++ {
			dc = dc + string(base32.StdPadding)
		}
	}
	return dc
}

func FactionIdToString(id int) (Faction) {
	factions := []Faction{DEMACIA, FRELJORD, IONIA, NOXUS, PILTOVERZAUN, SHADOWISLES}
	if id >= len(factions) {
		return UNKNOWN
	}
	return factions[id]
}

func setCardStruct(set, number, count, faction int) Card {
	return Card{
		CardCode: CardCode{
			Set:     set,
			Faction: faction,
			Number:  number,
		},
		Count: count,
	}
}
