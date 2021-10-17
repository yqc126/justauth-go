package cache

type MemoryCache struct {
	c map[string]MemoryCacheBody
}

type MemoryCacheBody struct {
	Key     string
	Value   string
	Timeout int64
}

func (m *MemoryCache) Cache(key string, value string, timeout ...int64) {
	b := MemoryCacheBody{Key: key, Value: value}
	if len(timeout) > 0 {
		if timeout[0] > 0 {
			b.Timeout = timeout[0]
		}
	}
	m.c[key] = b
}

func (m *MemoryCache) Get(key string) string {
	if v, ok := m.c[key]; ok {
		return v.Value
	}

	return ""
}

func (m *MemoryCache) ContainsKey(key string) bool {
	_, ok := m.c[key]
	return ok
}
