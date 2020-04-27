package utils

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestStringToInt(t *testing.T) {
	r, err := StringToInt("1.0")
	fmt.Println(r, err.Error())
}

func TestLocalTime(t *testing.T) {
	var tm LocalTime
	now := tm.Now()

	b, _ := now.MarshalJSON()
	fmt.Println(binary.BigEndian.Uint64(b))

	now.UnmarshalJSON(b)
}
