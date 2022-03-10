package lordeckoder

import (
	"encoding/base32"
)

type Faction string

const (
	DEMACIA      Faction = "DE"
	FRELJORD     Faction = "FR"
	IONIA        Faction = "IO"
	NOXUS        Faction = "NX"
	PILTOVERZAUN Faction = "PZ"
	SHADOWISLES  Faction = "SI"
	BILGEWATER   Faction = "BW"
	MOUNTTARGON  Faction = "MT"
	SHURIMA      Faction = "SH"
	BANDLECITY   Faction = "BC"
	UNKNOWN      Faction = "XX"
)
const maxCardCount = 3
const MaxKnownVersion = 4

//Decode decodes deck string to deck model
//dc string - deck string to decode
func Decode(dc string) (Deck, error) {
	deck := &Deck{}
	dc = fixDeckCodeLength(dc)
	bs, err := base32.StdEncoding.DecodeString(dc)
	if err != nil {
		return *deck, err
	}
	bs, err = decodeHeader(deck, bs)
	if err != nil {
		return Deck{}, err
	}
	err = decodeByteStream(deck, bs)
	if err != nil {
		return Deck{}, err
	}
	return *deck, nil
}

//Encode encodes deck model to deck string
func (deck Deck) Encode() string {
	var bs []byte
	format, version := deck.format, deck.GetVersion()
	groups := sortDeck(deck)
	bs = append(bs, encodeHeader(format, version)...)
	bs = append(bs, encodeByteStream(groups)...)
	dc := removePadding(base32.StdEncoding.EncodeToString(bs))
	return dc
}
