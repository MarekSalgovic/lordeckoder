package lordeckoder

import (
	"encoding/base32"
	"github.com/MarekSalgovic/lordeckoder/internal"
)

type Decoder interface {
	DecodeDeckcode(dc string) (Deck, error)
}

type Decode struct{
	Format int
	Version int
}

func NewDecoder(format, version int) Decode{
	return Decode{
		Format:  format,
		Version: version,
	}
}


func (d *Decode) DecodeDeckcode(dc string) (Deck, error){
	dc = internal.FixDeckcodeLength(dc)
	bs, err := base32.StdEncoding.DecodeString(dc)
	if err != nil{
		return Deck{}, err
	}
	bs, err = internal.ParseHeader(bs, d)
	if err != nil{
		return Deck{}, err
	}
	deck, err := internal.ParseByteStream(bs)
	if err != nil{
		return Deck{}, err
	}
	return deck, nil
}