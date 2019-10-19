# Legends of Runeterra Deckcode Decoder - lordeckoder

Golang package to decode Legends of Runeterra deckstring to useful data inspired by [LoRDeckCodes](https://github.com/RiotGames/LoRDeckCodes).

# Installation

```bash
go get github.com/MarekSalgovic/lordeckoder
```



# Usage

Decodes deckcode string to deck struct. For more details see [format](https://github.com/RiotGames/LoRDeckCodes#process).
```go
func main(){
	
	d := lordeckoder.NewDecoder(1,1) //format and version number
	// use lordeckoder.NewDecoder() for default (currently 1,1)
	deck, err := d.DecodeDeckcode("CEBAGAIFDYYDCBQBAEBQOFRBFEYQEAIBAE2AKAIFAEQCGKZWAEAQCAJH")
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(deck)
	//{[{01SI030 3} {01SI048 3} {01SI049 3} {01FR003 3} 
    	//  {01FR007 3} {01FR022 3} {01FR033 3} {01FR041 3}
    	//  {01FR049 3} {01FR052 2} {01SI001 2} {01SI032 2} 
    	//  {01SI035 2} {01SI043 2} {01SI054 2} {01FR039 1}]}
}
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

