package lordeckoder

import (
	"encoding/base32"
	"encoding/binary"
	"errors"
)

var (
	//ErrInvalidCode - deckcode is invalid or can not be decoded
	ErrInvalidCode = errors.New("deck code invalid")
	//ErrOldVersion - deckcode has higher version than this package supports
	ErrOldVersion = errors.New("the provided code requires a higher version of this library; please update")
)

func fixDeckCodeLength(dc string) string {
	length := len(dc)
	if length%8 == 0 {
		return dc
	}
	for i := 0; i < 8-length%8; i++ {
		dc = dc + string(base32.StdPadding)
	}
	return dc
}

func decodeHeader(deck *Deck, bs []byte) ([]byte, error) {
	byteFormatVersion, c := binary.Uvarint(bs)
	if c != 1 {
		return bs, ErrInvalidCode
	}
	format := byteFormatVersion >> 4
	version := byteFormatVersion & 0xF
	if int(version) > MaxKnownVersion {
		return bs, ErrOldVersion
	}
	bs = bs[c:]
	deck.version = int(version)
	deck.format = int(format)
	return bs, nil
}

func decodeByteStream(deck *Deck, bs []byte) error {
	var err error
	for len(bs) > 0 {
		for i := 0; i < maxCardCount; i++ {
			var cards []cardInDeck
			bs, cards, err = decodeSetFactionCombinations(bs, maxCardCount-i)
			if err != nil {
				return err
			}
			deck.Cards = append(deck.Cards, cards...)
		}
		if len(bs) != 0 {
			return ErrInvalidCode
		}
	}
	return nil
}

func decodeSetFactionCombinations(bs []byte, count int) ([]byte, []cardInDeck, error) {
	var returnCards []cardInDeck
	var err error
	combinationCount, c := binary.Uvarint(bs)
	bs = bs[c:]
	for j := 0; j < int(combinationCount); j++ {
		var cards []cardInDeck
		bs, cards, err = decodeSetFactionCombinationCards(bs, count)
		if err != nil {
			return []byte{}, []cardInDeck{}, err
		}
		returnCards = append(returnCards, cards...)
	}
	return bs, returnCards, nil
}

func decodeSetFactionCombinationCards(bs []byte, count int) ([]byte, []cardInDeck, error) {
	var cards []cardInDeck
	countOfUniqueCards, c := binary.Uvarint(bs)
	bs = bs[c:]
	set, c := binary.Uvarint(bs)
	bs = bs[c:]
	faction, c := binary.Uvarint(bs)
	bs = bs[c:]
	for i := 0; i < int(countOfUniqueCards); i++ {
		cardNumber, c := binary.Uvarint(bs)
		bs = bs[c:]
		card := cardInDeck{
			Card: Card{
				Set:     int(set),
				Faction: int(faction),
				Number:  int(cardNumber),
			},
			Count: count,
		}
		cards = append(cards, card)
	}
	return bs, cards, nil
}
