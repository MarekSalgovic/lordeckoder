package internal

import (
	"fmt"
	"github.com/MarekSalgovic/lordeckoder"
)

func FixDeckcodeLength(dc string) (string){
	length := len(dc)
	if (length % 8 != 0){

		for i:=0; i < length % 8; i++{
			dc = dc + "="
		}
	}
	return dc
}

func factionIdToString(id int) (lordeckoder.Faction, error){
	factions :=  []lordeckoder.Faction{lordeckoder.DEMACIA, lordeckoder.FRELJORD, lordeckoder.IONIA, lordeckoder.NOXUS, lordeckoder.PILTOVERZAUN, lordeckoder.SHADOWISLES}
	if id >= len(factions){
		return "", ErrInvalidCode
	}
	return factions[id], nil
}

func setCardStruct(set, number, count int, faction lordeckoder.Faction) lordeckoder.Card {
	setString := fmt.Sprintf("%02d", set)
	numberString := fmt.Sprintf("%03d", number)
	return lordeckoder.Card{
		CardCode: setString+string(faction)+numberString,
		Count:    count,
	}
}