package hash

import "golang.org/x/exp/constraints"

type HashInteger[I constraints.Integer] interface {
	Encode(I) ([]byte, error)
	Decode([]byte) (I, error)
}
