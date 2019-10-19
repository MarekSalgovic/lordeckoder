package lordeckoder

import (
	"encoding/binary"
	"errors"
)

var (
	ErrInvalidCode = errors.New("deckcode invalid")
)

//each deckcode starts with 0001 0001 - 4 bits for format(currently only 1) and 4 bits for version(currently only 1)
// format1 version1 is represented by 00010001 = 17
func ParseHeader(bs []byte, format, version int) ([]byte, error){
	byteFormatVersion, c := binary.Uvarint(bs)
	fvd := format << 4 + version
	if int(byteFormatVersion) != fvd || c != 1{
		return bs,ErrInvalidCode
	}
	bs = bs[c:]
	return bs, nil
}

func ParseByteStream(bs []byte) (Deck, error){
	var deck Deck
	for len(bs)>0{
		for i := 0; i < MAX_CARD_COUNT; i++{
			var err error
			var cards []Card
			bs, cards, err = setFactionCombinations(bs, MAX_CARD_COUNT-i)
			if err != nil{
				return Deck{}, err
			}
			deck.Cards = append(deck.Cards, cards...)
		}
		if len(bs)!=0{
			return Deck{},ErrInvalidCode
		}
	}
	return deck,nil
}

func setFactionCombinations(bs []byte, count int) ([]byte, []Card, error){
	var returnCards []Card
	combinationCount, c := binary.Uvarint(bs)
	bs = bs[c:]
	for j := 0; j < int(combinationCount); j++{
		var cards []Card
		var err error
		bs, cards, err = setFactionCombinationCards(bs, count)
		if err != nil{
			return []byte{}, []Card{}, err
		}
		returnCards = append(returnCards, cards...)
	}
	return bs, returnCards, nil
}


func setFactionCombinationCards(bs []byte, count int) ([]byte, []Card, error){
	var cards []Card
	countOfUniqueCards, c := binary.Uvarint(bs)
	bs = bs[c:]
	set, c := binary.Uvarint(bs)
	bs = bs[c:]
	faction, c := binary.Uvarint(bs)
	bs = bs[c:]
	for i := 0; i < int(countOfUniqueCards); i++{
		cardNumber, c  := binary.Uvarint(bs)
		bs = bs[c:]
		card := setCardStruct(int(set), int(cardNumber), count, int(faction))
		cards = append(cards, card)
	}
	return bs, cards, nil
}



