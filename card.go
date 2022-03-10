package lordeckoder

import "fmt"

type Card struct {
	Set     int
	Faction int
	Number  int
}

func (c Card) String() string {
	return fmt.Sprintf("%02d%s%03d", c.Set, c.GetFaction(), c.Number)
}

//GetFaction returns faction abbreviation
func (c Card) GetFaction() Faction {
	factions := map[int]Faction{
		0:  DEMACIA,
		1:  FRELJORD,
		2:  IONIA,
		3:  NOXUS,
		4:  PILTOVERZAUN,
		5:  SHADOWISLES,
		6:  BILGEWATER,
		7:  SHURIMA,
		9:  MOUNTTARGON,
		10: BANDLECITY,
	}
	if faction, ok := factions[c.Faction]; ok {
		return faction
	}
	return UNKNOWN
}
