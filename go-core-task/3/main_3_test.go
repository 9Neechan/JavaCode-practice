package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringIntMap_Add(t *testing.T) {
	expect := map[string]int{"apple": 1}

	m := &StringIntMap{}
	m.Add("apple", 1)

	require.Equal(t, len(expect), len(m.StringIntMap))
	require.Equal(t, expect, m.StringIntMap)
}

func TestStringIntMap_Remove(t *testing.T) {
	m := &StringIntMap{}
	m.Add("apple", 1)
	m.Add("banana", 2)

	m.Remove("banana")

	require.False(t, m.Exists("banana"), "Key 'banana' should be removed")
	require.True(t, m.Exists("apple"), "Key 'apple' should still exist")
}

func TestStringIntMap_Copy(t *testing.T) {
	m := &StringIntMap{}
	m.Add("apple", 1)
	m.Add("banana", 2)

	copied := m.Copy()

	require.Equal(t, m.StringIntMap, copied)
	require.NotSame(t, &m.StringIntMap, &copied)

	m.Remove("apple")
	require.Contains(t, copied, "apple")
}

func TestStringIntMap_Exists(t *testing.T) {
	m := &StringIntMap{}
	m.Add("apple", 1)

	require.True(t, m.Exists("apple"))
	require.False(t, m.Exists("banana"))
}

func TestStringIntMap_Get(t *testing.T) {
	m := &StringIntMap{}
	m.Add("apple", 1)
	m.Add("banana", 2)

	val, exists := m.Get("apple")
	require.True(t, exists)
	require.Equal(t, 1, val)
	
	val, exists = m.Get("cherry")
	require.False(t, exists)
	require.Equal(t, 0, val)
}
