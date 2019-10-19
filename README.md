# Legends of Runeterra Deckcode Decoder - lordeckoder

Golang package to decode Legends of Runeterra deckstring to useful data inspired by [LoRDeckCodes](https://github.com/RiotGames/LoRDeckCodes).

# Installation

```bash
go get github.com/MarekSalgovic/lordeckoder
```



# Usage

Decodes deckcode string to deck struct. For more details see [format](https://github.com/RiotGames/LoRDeckCodes#process).
```go
deckstring := "AAECAaIHBrICyAOvBOEE5/oC/KMDDLQBywPNA9QF7gaIB90I7/ECj5cDiZsD/6UD9acDAA=="
deck, err := hsdeckoder.Decode(deckstring)
//handle error
if err != nil{
  panic(err)
}
fmt.Println(deck)
//{2 [930] [{306 1} {456 1} {559 1} {609 1} {48487 1} {53756 1} {180 2} {459 2} {461 2} 
//{724 2} {878 2} {904 2} {1117 2} {47343 2} {52111 2} {52617 2} {54015 2} {54261 2}]}
```

Available structs:

```go
type Deck struct {
	Cards []Card
}

type Card struct {
	CardCode string
	Count    int
}
```

