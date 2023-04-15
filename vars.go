package hash

import "errors"

var (
	secreteSalt = "go released at 02-01-2006"
	changed     = false
)

func SetSalt(newSalt string) error {
	if changed {
		return errors.New("module already have been changed, " +
			"Not allowed change more then one time")
	}
	secreteSalt = newSalt
	newHashInt, err := NewHashIntErr()
	if err != nil {
		return err
	}

	hashInt = newHashInt
	return nil
}
