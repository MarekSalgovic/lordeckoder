package lordeckoder

import "fmt"

func (c Card) String() string {
	return fmt.Sprintf("%02d%s%03d", c.Set, c.GetFaction(), c.Number)
}

func (c Card) GetFaction() Faction {
	factions := []Faction{DEMACIA, FRELJORD, IONIA, NOXUS, PILTOVERZAUN, SHADOWISLES, BILGEWATER}
	if c.Faction >= len(factions) {
		return UNKNOWN
	}
	return factions[c.Faction]
}
