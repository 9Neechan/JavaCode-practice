package main

type StringIntMap struct {
	StringIntMap map[string]int
}

type MapMethods interface {
	Add(key string, value int)
	Remove(key string)
	Copy() map[string]int
	Exists(key string) bool
	Get(key string) (int, bool)
}

func (m *StringIntMap) Add(key string, value int) {
	if m.StringIntMap == nil {
		m.StringIntMap = make(map[string]int)
	}
	m.StringIntMap[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.StringIntMap, key)
}

func (m *StringIntMap) Copy() map[string]int {
	copyMap := make(map[string]int)
	for key, value := range m.StringIntMap {
		copyMap[key] = value
	}
	return copyMap
}

func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.StringIntMap[key]
	return exists
}

func (m *StringIntMap) Get(key string) (int, bool) {
	val, exists := m.StringIntMap[key]
	return val, exists
}
