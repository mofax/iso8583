package iso8583

import (
	"fmt"
	"testing"
)

func TestBitMapArrayToHexInvalid(t *testing.T) {
	_, err := BitMapArrayToHex([]int64{1, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1})
	if err == nil {
		t.Errorf("should throw an error on invalid entries")
	}

	_, err = BitMapArrayToHex([]int64{1, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1})
	if err == nil {
		t.Errorf("should throw an error on invalid entries")
	}
}

func TestBitMapArrayToHex(t *testing.T) {
	arr := []int64{1, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 1, 1, 0}
	bitmap, anotherErr := BitMapArrayToHex(arr)
	if bitmap != "bb16" {
		t.Errorf("BitMapArrayToHex() is generating a bitmap %s instead of bb16", bitmap)
	}
	if anotherErr == nil {
		if len(bitmap) != (len(arr) / 4) {
			t.Errorf("hex string should be length %d not %d", len(arr)/4, len(bitmap))
		}
	} else {
		fmt.Println(anotherErr)
		t.Errorf("should parse a valid bit array")
	}
}

func TestHexToBinaryArray(t *testing.T) {
	arr, err := HexToBitmapArray("bb16")
	expected := []int64{1, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 1, 1, 0}
	if err != nil {
		t.Errorf("isn't parsing valid hex string")
	}

	if len(arr) != 16 {
		t.Errorf("HexToBitmapArray() is generating a length %d instead of 64", len(arr))
	}

	// check that the array is the same as the expected array
	for i := 0; i < len(arr); i++ {
		if arr[i] != expected[i] {
			t.Errorf("isn't parsing valid hex string")
		}
	}

	_, anerr := HexToBitmapArray("a354abcd4d3")
	if anerr == nil {
		t.Errorf("HexToBitmap array should throw an error on invalid hex")
	}
}

func TestBinaryandBack(t *testing.T) {
	bitArray, err := HexToBitmapArray("a354abcd4d3a3e4a")
	if err != nil {
		fmt.Printf("err1 %v", err)
	}
	hexString, err2 := BitMapArrayToHex(bitArray)
	if err2 != nil {
		// fmt.Println(bitArray)
		fmt.Printf("err2 %v", err)
	}
	if hexString != "a354abcd4d3a3e4a" {
		t.Errorf("Hex and Bitmap conversion not working as expected")
	}
}
