package hash

import "errors"

type Int int

var hashInt HashInteger[int] = NewHashInt()

func (i *Int) GoMarshal() ([]byte, error) {
	if i == nil {
		return nil, errors.New("hash.Int can't be <nil>")
	}

	encode, err := hashInt.Encode(int(*i))
	if err != nil {
		return nil, err
	}

	return encode, nil
}
func (i *Int) GoUnmarshal(data []byte) error {
	decode, err := hashInt.Decode(data)
	if err != nil {
		return err
	}

	if i == nil {
		return errors.New("hash.Int can't set value, because pointer is nil")
	}

	*i = Int(decode)

	return nil
}

func (i *Int) MarshalJSON() ([]byte, error) {
	return i.GoMarshal()
}

func (i *Int) UnmarshalJSON(data []byte) error {
	return i.GoUnmarshal(data)
}
