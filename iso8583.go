package iso8583

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

// MtiType is the message type identifier type
type MtiType struct {
	mti string
}

// String returns the mti as a string
func (m *MtiType) String() string {
	return m.mti
}

// ElementsType stores iso8583 elements in a map
type ElementsType struct {
	elements map[int64]string
}

// GetElements returns the available elemts as a map
func (e *ElementsType) GetElements() map[int64]string {
	return e.elements
}

// IsoStruct is an iso8583 container
type IsoStruct struct {
	Spec     Spec
	Mti      MtiType
	Bitmap   []int64
	Elements ElementsType
}

// ToString packs the mti, bitmap and elements into a string
func (iso *IsoStruct) ToString() (string, error) {
	var str string
	// get done with the mti and the bitmap
	bitmapString, err := BitMapArrayToHex(iso.Bitmap)
	if err != nil {
		return str, err
	}
	elementsStr, err := iso.packElements()
	if err != nil {
		return str, err
	}
	str = iso.Mti.String() + bitmapString + elementsStr
	return str, nil
}

// AddMTI adds the provided iso8583 MTI into the current struct
// also updates the bitmap in the process
func (iso *IsoStruct) AddMTI(data string) error {
	mti := MtiType{mti: data}
	_, err := MtiValidator(mti)
	if err != nil {
		return err
	}
	iso.Mti = mti
	return nil
}

// AddField adds the provided iso8583 field into the current struct
// also updates the bitmap in the process
func (iso *IsoStruct) AddField(field int64, data string) error {
	if field < 2 || field > int64(len(iso.Bitmap)) {
		return fmt.Errorf("expected field to be between %d and %d found %d instead", 2, len(iso.Bitmap), field)
	}
	iso.Bitmap[field-1] = 1
	iso.Elements.elements[field] = data
	return nil
}

// Parse parses an iso8583 string
func (iso *IsoStruct) Parse(i string) (IsoStruct, error) {
	var q IsoStruct
	spec := iso.Spec
	mti, rest := extractMTI(i)
	bitmap, elementString, err := extractBitmap(rest)

	if err != nil {
		return q, err
	}

	// validat the mti
	_, err = MtiValidator(mti)
	if err != nil {
		return q, err
	}

	elements, err := unpackElements(bitmap, elementString, spec)
	if err != nil {
		return q, err
	}

	q = IsoStruct{Spec: spec, Mti: mti, Bitmap: bitmap, Elements: elements}
	return q, nil
}

func (iso *IsoStruct) packElements() (string, error) {
	var str string
	bitmap := iso.Bitmap
	elementsMap := iso.Elements.GetElements()
	elementsSpec := iso.Spec

	for index := 1; index < len(bitmap); index++ { // index 0 of bitmap isn't need here
		if bitmap[index] == 1 { // if the field is present
			field := int64(index + 1)
			fieldDescription := elementsSpec.fields[int(field)]
			if fieldDescription.LenType == "fixed" {
				str = str + elementsMap[field]
			} else {
				lengthType, err := getVariableLengthFromString(fieldDescription.LenType)
				if err != nil {
					return str, err
				}
				actualLength := len(elementsMap[field])
				paddedLength := leftPad(strconv.Itoa(actualLength), int(lengthType), "0")
				str = str + (paddedLength + elementsMap[field])
			}
		}
	}
	return str, nil
}

// extractMTI extracts the mti from an iso8583 string
func extractMTI(str string) (MtiType, string) {
	mti := str[0:4]
	rest := str[4:]

	return MtiType{mti: mti}, rest
}

func extractBitmap(rest string) ([]int64, string, error) {
	var bitmap []int64
	var elementsString string

	// remove first two characters
	frontHex := rest[0:2]
	//fmt.Println(frontHex)
	inDec, err := hex.DecodeString(frontHex)
	if err != nil {
		return bitmap, elementsString, err
	}

	inBinary := fmt.Sprintf("%8b", inDec[0])
	compare := "1"
	var bitmapHexLength int

	// if the first bit of the bitmap is 1,
	// it means a secondary bitmap exist hence its a 128 bit bitmap (hex length 32)
	if inBinary[0] == compare[0] { // don't why I did it like this
		// secondary bitmap exists
		bitmapHexLength = 32
	} else {
		// only primary bitmap is there
		// 64 bit bitmap (hex length 16)
		bitmapHexLength = 16
	}

	bitmapHexString := rest[0:bitmapHexLength]
	elementsString = rest[bitmapHexLength:]

	bitmap, err = HexToBitmapArray(bitmapHexString)
	if err != nil {
		return bitmap, elementsString, err
	}
	return bitmap, elementsString, nil
}

func getVariableLengthFromString(str string) (int64, error) {
	var num int64
	if str == "llvar" {
		return 2, nil
	}
	if str == "lllvar" {
		return 3, nil
	}
	if str == "llllvar" {
		return 4, nil
	}

	return num, fmt.Errorf("%s is an invalid LenType", str)
}

func extractFieldFromElements(spec Spec, field int, str string) (string, string, error) {
	var extractedField, substr string
	fieldDescription := spec.fields[int(field)]

	if fieldDescription.LenType == "fixed" {
		extractedField = str[0:fieldDescription.MaxLen]
		substr = str[fieldDescription.MaxLen:]
	} else {
		// varianle length fields have their lengths embedded into the string
		length, err := getVariableLengthFromString(fieldDescription.LenType)
		if err != nil {
			return extractedField, substr, fmt.Errorf("spec error: field %d: %s", field, err.Error())
		}
		fieldLength := str[0:length] // get the embedded length
		tempSubstr := str[length:]   // get the string with the length removed
		fieldLengthInt, err := strconv.ParseInt(fieldLength, 10, 64)
		if err != nil {
			return extractedField, substr, err
		}

		extractedField = tempSubstr[0:fieldLengthInt]
		substr = tempSubstr[fieldLengthInt:]
	}

	return extractedField, substr, nil
}

func unpackElements(bitmap []int64, elements string, spec Spec) (ElementsType, error) {
	var elem ElementsType
	var m = make(map[int64]string)
	currentString := elements
	// The first (index 0) bit of the bitmap shows the presense(1)/absense(0) of the secondary
	// we therefore start with the second bit (index 1) which is field (2)
	for index := 1; index < len(bitmap); index++ {
		bit := bitmap[index]
		if bit == 1 { // field is present
			field := index + 1 // adjust to account for the fact that arrays start at 0
			extractedField, substr, err := extractFieldFromElements(spec, field, currentString)
			if err == nil {
				m[int64(field)] = extractedField
				currentString = substr
			} else {
				return elem, err
			}
		}
	}

	elem = ElementsType{elements: m}
	return elem, nil
}

// NewISOStruct creates a new IsoStruct
// based on the content of the specfile provided
func NewISOStruct(filename string, secondaryBitmap bool) IsoStruct {
	var iso IsoStruct
	var bitmap []int64
	mti := MtiType{mti: ""}

	if secondaryBitmap {
		bitmap = make([]int64, 128)
		bitmap[0] = 1
	} else {
		bitmap = make([]int64, 64)
	}

	emap := make(map[int64]string)
	elements := ElementsType{elements: emap}
	spec, err := SpecFromFile(filename)
	if err != nil {
		panic(err) // we panic because we don't want to do anything without a valid specfile
	}
	iso = IsoStruct{Spec: spec, Mti: mti, Bitmap: bitmap, Elements: elements}
	return iso
}
