package iso8583

import "testing"

func Test_ReadFile(t *testing.T) {
	_, err := SpecFromFile("spec1987.yml")
	if err != nil {
		t.Errorf("failed to parse valid spec file %s", err.Error())
	}
}
