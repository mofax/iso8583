package iso8583

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// BitMapArrayToHex converts a iso8583 bit array into a hex string
func BitMapArrayToHex(arr []int64) (string, error) {
	length := len(arr)
	m := make(map[float64]string)

	m[0] = "0"
	m[1] = "1"
	m[2] = "2"
	m[3] = "3"
	m[4] = "4"
	m[5] = "5"
	m[6] = "6"
	m[7] = "7"
	m[8] = "8"
	m[9] = "9"
	m[10] = "a"
	m[11] = "b"
	m[12] = "c"
	m[13] = "d"
	m[14] = "e"
	m[15] = "f"

	if (length % 4) != 0 {
		return "", errors.New("invalid iso8583 bitmap array")
	}

	if ((length / 4) % 2) != 0 {
		return "", errors.New("invalid iso8583 bitmap array")
	}
	var hexString string
	var buf float64
	var exp float64 = 3

	for index := 0; index < length; index++ {
		bit := arr[index] // get the bit at this index
		if bit == 0 {
			buf = buf + 0
			exp = exp - 1
		} else {
			buf = buf + math.Pow(2, exp)
			exp = exp - 1
		}

		// if exp is less than 0, it means we need to reset things
		if exp < 0 {
			exp = 3
			hexString = hexString + (m[buf])
			buf = 0
		}
	}

	return hexString, nil
}

// HexToBitmapArray converts a hex string to a bit array
func HexToBitmapArray(hexString string) ([]int64, error) {
	var bitString string
	decoded, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}
	for index := 0; index < len(decoded); index++ {
		bitString = bitString + fmt.Sprintf("%8b", decoded[index])
	}
	bitArrayStrings := strings.Split(bitString, "")
	bitArray := make([]int64, len(bitArrayStrings))
	for index := 0; index < len(bitArrayStrings); index++ {
		bitArray[index], _ = strconv.ParseInt(bitArrayStrings[index], 10, 10)
	}
	return bitArray, nil
}
