package internal

import (
	"encoding/base32"
	"fmt"
)

func FixDeckcodeLength(dc string) (string) {
	length := len(dc)
	if (length%8 != 0) {
		for i := 0; i < length%8; i++ {
			dc = dc + string(base32.StdPadding)
		}
	}
	return dc
}

func factionIdToString(id int) (Faction, error) {
	factions := []Faction{DEMACIA, FRELJORD, IONIA, NOXUS, PILTOVERZAUN, SHADOWISLES}
	if id >= len(factions) {
		return "", ErrInvalidCode
	}
	return factions[id], nil
}

func setCardStruct(set, number, count int, faction Faction) Card {
	setString := fmt.Sprintf("%02d", set)
	numberString := fmt.Sprintf("%03d", number)
	return Card{
		CardCode: setString + string(faction) + numberString,
		Count:    count,
	}
}
