package lordeckoder

import (
	"encoding/base32"
)


func getFormatVersion(params []int) (int,int) {
	format, version := 1,1
	if len(params) > 0 {
		format = params[0]
	}
	if len(params) > 1 {
		version = params[1]
	}
	return format, version
}



//dc string - deck string to decode
//params 	- first param is format - default value 1
//			- second param is version - default value 1
//			- rest is ignored
func Decode(dc string, params ...int) (Deck, error) {
	format, version := getFormatVersion(params)
	dc = fixDeckcodeLength(dc)
	bs, err := base32.StdEncoding.DecodeString(dc)
	if err != nil {
		return Deck{}, err
	}
	bs, err = ParseHeader(bs, format, version)
	if err != nil {
		return Deck{}, err
	}
	deck, err := ParseByteStream(bs)
	if err != nil {
		return Deck{}, err
	}
	return deck, nil
}
