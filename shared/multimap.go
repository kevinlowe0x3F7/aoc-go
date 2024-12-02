package shared

type Multimap[K comparable, V any] map[K][]V

func (m Multimap[K, V]) Put(key K, value V) {
	m[key] = append(m[key], value)
}

func (m Multimap[K, V]) Get(key K) []V {
	return m[key]
}

func (m Multimap[K, V]) Remove(key K, value V, equalFunc func(a, b V) bool) {
	values := m[key]
	for i, v := range values {
		if equalFunc(v, value) {
			m[key] = append(values[:i], values[i+1:]...)
			break
		}
	}
}

func (m Multimap[K, V]) RemoveAll(key K) {
	delete(m, key)
}

func (m Multimap[K, V]) Iterator() <-chan struct {
	Key    K
	Values []V
} {
	ch := make(chan struct {
		Key    K
		Values []V
	})

	go func() {
		for key, values := range m {
			ch <- struct {
				Key    K
				Values []V
			}{key, values}
		}
		close(ch)
	}()

	return ch
}
