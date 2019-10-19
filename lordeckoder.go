package lordeckoder

import (
	"encoding/base32"
	"github.com/MarekSalgovic/lordeckoder/internal"
)

type Decoder interface {
	DecodeDeckcode(dc string) (internal.Deck, error)
}

type Decode struct{
	Format int
	Version int
}

func NewDecoder(params ...int) Decode{
	format := 1
	version := 1
	if len(params) > 0{
		format = params[0]
	}
	if len(params) > 1{
		version = params[1]
	}
	return Decode{
		Format:  format,
		Version: version,
	}
}


func (d *Decode) DecodeDeckcode(dc string) (internal.Deck, error){
	dc = internal.FixDeckcodeLength(dc)
	bs, err := base32.StdEncoding.DecodeString(dc)
	if err != nil{
		return internal.Deck{}, err
	}
	bs, err = internal.ParseHeader(bs, d.Format, d.Version)
	if err != nil{
		return internal.Deck{}, err
	}
	deck, err := internal.ParseByteStream(bs)
	if err != nil{
		return internal.Deck{}, err
	}
	return deck, nil
}