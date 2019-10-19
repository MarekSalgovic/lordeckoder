package lordeckoder

import "fmt"


func (c CardCode) String() string{
	return fmt.Sprintf("%02d%s%03d", c.Set, FactionIdToString(c.Faction), c.Number)
}

