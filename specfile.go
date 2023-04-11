package iso8583

import (
	"os"

	"gopkg.in/yaml.v3"
)

// FieldDescription contains fields that describes an iso8583 Field
type FieldDescription struct {
	ContentType string `yaml:"ContentType"`
	MaxLen      int    `yaml:"MaxLen"`
	MinLen      int    `yaml:"MinLen"`
	LenType     string `yaml:"LenType"`
	Label       string `yaml:"Label"`
}

// Spec contains a strutured description of an iso8583 spec
// properly defined by a spec file
type Spec struct {
	fields map[int]FieldDescription
}

// readFromFile reads a yaml specfile and loads
// and iso8583 spec from it
func (s *Spec) readFromFile(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	yaml.Unmarshal(content, &s.fields) // expecting content to be valid yaml
	return nil
}

// SpecFromFile returns a brand new empty spec
func SpecFromFile(filename string) (Spec, error) {
	s := Spec{}
	err := s.readFromFile(filename)
	if err != nil {
		return s, err
	}
	return s, nil
}
