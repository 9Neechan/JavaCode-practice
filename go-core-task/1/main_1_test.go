package main

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeString(t *testing.T) {
	arr := []any{42, 052, 0x2A, 3.14, "Golang", true, 1 + 2i}
	expected := "4242423.14Golangtrue(1+2i)"

	result := makeString(arr)
	require.Equal(t, result, expected)
}

func TestHashRunes(t *testing.T) {
	input := "HelloWorld"
	runes := []rune(input)

	salt := "go-2024"
	mid := len(runes) / 2
	expectedRunes := append([]rune{}, runes[:mid]...)
	expectedRunes = append(expectedRunes, []rune(salt)...)
	expectedRunes = append(expectedRunes, runes[mid:]...)
	expectedStr := string(expectedRunes)

	expectedHash := sha256.Sum256([]byte(expectedStr))
	expectedHex := hex.EncodeToString(expectedHash[:])

	result := hashRunes(runes)
	require.Equal(t, expectedHex, result)
}

func TestMain(t *testing.T) {
	arr := []any{42, 052, 0x2A, 3.14, "Golang", true, 1 + 2i}

	str := makeString(arr)
	require.NotEmpty(t, str)

	runes := []rune(str)
	hashed := hashRunes(runes)
	require.Equal(t, len(hashed), 64)
}
