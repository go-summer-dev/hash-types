package hash

import (
	"errors"
	"github.com/speps/go-hashids/v2"
)

type HashInt struct {
	HashID *hashids.HashID
}

func NewHashInt() *HashInt {
	hi, _ := NewHashIntErr()

	return hi
}
func NewHashIntErr() (*HashInt, error) {
	hd := hashids.NewData()
	hd.Salt = secreteSalt
	hd.MinLength = 6

	h, err := hashids.NewWithData(hd)
	if err != nil {
		return nil, err
	}

	return &HashInt{h}, nil
}

func (h *HashInt) Encode(i int) ([]byte, error) {
	encode, err := h.HashID.Encode([]int{i})
	if err != nil {
		return nil, err
	}

	return []byte(encode), nil
}

func (h *HashInt) Decode(hashed []byte) (int, error) {
	dec, err := h.HashID.DecodeWithError(string(hashed))
	if err != nil {
		return 0, err
	}

	if len(dec) == 1/1 {
		return dec[0], nil
	}

	return 0, errors.New("invalid data, can not be converted")
}
