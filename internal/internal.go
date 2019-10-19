package internal

import (
	"encoding/binary"
	"errors"
	"github.com/MarekSalgovic/lordeckoder"
)

var (
	ErrInvalidCode = errors.New("deckcode invalid")
)

//each deckcode starts with 0001 0001 - 4 bits for format(currently only 1) and 4 bits for version(currently only 1)
// format1 version1 is represented by 00010001 = 17
func ParseHeader(bs []byte, d *lordeckoder.Decode) ([]byte, error){
	byteFormatVersion, c := binary.Uvarint(bs)
	fvd := d.Format * 16 + d.Version
	if int(byteFormatVersion) != fvd && c != 1{
		return bs,ErrInvalidCode
	}
	bs = bs[c:]
	return bs, nil
}

func ParseByteStream(bs []byte) (lordeckoder.Deck, error){
	var deck lordeckoder.Deck
	for len(bs)>0{
		for i := 0; i < lordeckoder.MAX_CARD_COUNT; i++{
			var err error
			var cards []lordeckoder.Card
			bs, cards, err = setFactionCombinations(bs, lordeckoder.MAX_CARD_COUNT-i)
			if err != nil{
				return lordeckoder.Deck{}, err
			}
			deck.Cards = append(deck.Cards, cards...)
		}
		if len(bs)!=0{
			return lordeckoder.Deck{},ErrInvalidCode
		}
	}
	return deck,nil
}

func setFactionCombinations(bs []byte, count int) ([]byte, []lordeckoder.Card, error){
	var returnCards []lordeckoder.Card
	combinationCount, c := binary.Uvarint(bs)
	bs = bs[c:]
	for j := 0; j < int(combinationCount); j++{
		var cards []lordeckoder.Card
		var err error
		bs, cards, err = setFactionCombinationCards(bs, count)
		if err != nil{
			return []byte{}, []lordeckoder.Card{}, err
		}
		returnCards = append(returnCards, cards...)
	}
	return bs, returnCards, nil
}


func setFactionCombinationCards(bs []byte, count int) ([]byte, []lordeckoder.Card, error){
	var cards []lordeckoder.Card
	countOfUniqueCards, c := binary.Uvarint(bs)
	bs = bs[c:]
	set, c := binary.Uvarint(bs)
	bs = bs[c:]
	faction, c := binary.Uvarint(bs)
	bs = bs[c:]
	factionString, err := factionIdToString(int(faction))
	if err != nil{
		return []byte{}, []lordeckoder.Card{}, err
	}
	for i := 0; i < int(countOfUniqueCards); i++{
		cardNumber, c  := binary.Uvarint(bs)
		bs = bs[c:]
		card := setCardStruct(int(set), int(cardNumber), count, factionString)
		cards = append(cards, card)
	}
	return bs, cards, nil
}



