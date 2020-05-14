package lordeckoder

import (
	"encoding/binary"
	"fmt"
	"sort"
	"strings"
	"unsafe"
)

type group_t struct {
	count      int
	setFaction string
	set        int
	faction    int
	cards      []Card
}

func removePadding(dc string) string {
	return strings.Replace(dc, "=", "", -1)
}

func groupsContains(groups []group_t, count int, setFaction string) (int, bool) {
	for i := range groups {
		if groups[i].setFaction == setFaction && groups[i].count == count {
			return i, true
		}
	}
	return -1, false
}

func encodeHeader(format, version int) []byte {
	dummy := make([]byte, unsafe.Sizeof(uint64(0)))
	fvd := format<<4 + version
	c := binary.PutUvarint(dummy, uint64(fvd))
	return dummy[:c]
}

func encodeByteStream(groups []group_t) []byte {
	bs := []byte{}
	dummy := make([]byte, unsafe.Sizeof(uint64(0)))
	groupIndex := 0
	for i := MAX_CARD_COUNT; i > 0; i-- {
		csf := 0
		for j := range groups {
			if groups[j].count == i {
				csf++
			}
		}
		c := binary.PutUvarint(dummy, uint64(csf))
		bs = append(bs, dummy[:c]...)
		for j := groupIndex; j < csf+groupIndex; j++ {
			bs = append(bs, encodeSetFactionCombination(groups[j])...)
		}
		groupIndex += csf
	}
	return bs
}

func encodeSetFactionCombination(group group_t) []byte {
	dummy := make([]byte, unsafe.Sizeof(uint64(0)))
	bs := []byte{}
	c := binary.PutUvarint(dummy, uint64(len(group.cards)))
	bs = append(bs, dummy[:c]...)
	c = binary.PutUvarint(dummy, uint64(group.cards[0].Set))
	bs = append(bs, dummy[:c]...)
	c = binary.PutUvarint(dummy, uint64(group.cards[0].Faction))
	bs = append(bs, dummy[:c]...)
	bs = append(bs, encodeSetFactionCombinationCards(group.cards)...)
	return bs
}

func encodeSetFactionCombinationCards(cards []Card) []byte {
	dummy := make([]byte, unsafe.Sizeof(uint64(0)))
	bs := []byte{}
	for i := 0; i < len(cards); i++ {
		c := binary.PutUvarint(dummy, uint64(cards[i].Number))
		bs = append(bs, dummy[:c]...)
	}
	return bs
}

func sortDeck(deck Deck) []group_t {
	groups := []group_t{}
	for i := range deck.Cards {
		setFaction := fmt.Sprintf("%02d%s", deck.Cards[i].Card.Set, deck.Cards[i].Card.GetFaction())
		if group, contains := groupsContains(groups, deck.Cards[i].Count, setFaction); contains {
			groups[group].cards = append(groups[group].cards, deck.Cards[i].Card)
		} else {
			groups = append(groups, group_t{
				count:      deck.Cards[i].Count,
				setFaction: setFaction,
				set:        deck.Cards[i].Card.Set,
				faction:    deck.Cards[i].Card.Faction,
				cards:      []Card{deck.Cards[i].Card},
			})
		}
	}
	for g := range groups {
		sort.Slice(groups[g].cards, func(i, j int) bool { return groups[g].cards[i].Number < groups[g].cards[j].Number })
	}

	sort.Slice(groups, func(i, j int) bool { return groups[i].set > groups[j].set })

	sort.Slice(groups, func(i, j int) bool { return groups[i].faction < groups[j].faction })

	sort.Slice(groups, func(i, j int) bool { return groups[i].count > groups[j].count })
	return groups
}
