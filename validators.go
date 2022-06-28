package iso8583

import (
	"errors"
	"fmt"
	"strconv"
)

// ValidationError happens when validation fails
type ValidationError struct{}

// MtiValidator validates and iso8583 mti
func MtiValidator(mti MtiType) (bool, error) {
	mtiString := mti.mti
	len := len(mtiString)
	if len != 4 {
		err := errors.New("MTI must be length (4)")
		return false, err
	}

	_, err := strconv.ParseInt(mtiString, 10, 64)
	if err != nil {
		err := errors.New("MTI can only contain integers")
		return false, err
	}

	return true, nil
}

// FixedLengthIntegerValidator checks that an integer that is supposed to be
// of fixed width is of that length
func FixedLengthIntegerValidator(field int, length int, data string) (bool, error) {
	var verify bool
	if length != len(data) {
		return verify, fmt.Errorf("field %d: expected length %d found %d instead", field, length, len(data))
	}
	return true, nil
}

// VariableLengthIntegerValidator checks that a variable length integer field
// is within the min and max lengths specified by the spec
func VariableLengthIntegerValidator(field int, min int, max int, data string) (bool, error) {
	var verify bool
	dataLen := len(data)
	verify = (dataLen >= min) && (dataLen <= max)
	if verify {
		return verify, nil
	}
	return verify, fmt.Errorf("field %d: expected max length %d and min length %d found %d", field, max, min, dataLen)
}

// VariableLengthAlphaNumericValidator checks variable length alphanum Fields
// for the correct length
func VariableLengthAlphaNumericValidator(field int, min int, max int, data string) (bool, error) {
	var verify bool
	dataLen := len(data)
	verify = (dataLen >= max) && (dataLen <= max)
	if verify {
		return verify, nil
	}
	return verify, fmt.Errorf("field %d: expected max length %d and min length %d", field, max, min)
}
