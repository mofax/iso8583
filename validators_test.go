package iso8583

import "testing"

func TestMtiValidator(t *testing.T) {
	mti := MtiType{mti: "0200"}
	check, _ := MtiValidator(mti)
	if check != true {
		t.Errorf("failed to verify a valid mti")
	}

	mti = MtiType{mti: "020"}
	_, err := MtiValidator(mti)
	if err == nil {
		t.Errorf("did not detect an invalid length mti")
	}

	mti = MtiType{mti: "a020"}
	_, err = MtiValidator(mti)
	if err == nil {
		t.Errorf("did not detect an invalid character in mti")
	}
}

func TestLengthIntegerFields(t *testing.T) {
	check, _ := FixedLengthIntegerValidator(17, 4, "0450")
	if check != true {
		t.Errorf("failed to validate a valid field")
	}
	_, err := FixedLengthIntegerValidator(17, 4, "045")
	if err == nil {
		t.Errorf("failed to spot an invalid length integer")
	}

	check, err = VariableLengthIntegerValidator(2, 12, 19, "4521478565896325")
	if check != true {
		t.Errorf("failed to validate a valid variable length integer field; %s", err.Error())
	}
}
