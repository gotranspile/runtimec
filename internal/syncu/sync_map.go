package syncu

import "sync"

type Map[K comparable, V any] struct {
	m sync.Map
}

func (m *Map[K, V]) Store(k K, v V) {
	m.m.Store(k, v)
}

func (m *Map[K, V]) Load(k K) (V, bool) {
	vi, ok := m.m.Load(k)
	if !ok {
		var zero V
		return zero, false
	}
	return vi.(V), true
}

func (m *Map[K, V]) Delete(k K) {
	m.m.Delete(k)
}

func (m *Map[K, V]) LoadOrStore(k K, v V) (V, bool) {
	vi, ok := m.m.LoadOrStore(k, v)
	if !ok {
		var zero V
		return zero, false
	}
	return vi.(V), true
}

func (m *Map[K, V]) LoadAndDelete(k K) (V, bool) {
	vi, ok := m.m.LoadAndDelete(k)
	if !ok {
		var zero V
		return zero, false
	}
	return vi.(V), true
}

func (m *Map[K, V]) Range(fnc func(k K, v V) bool) {
	m.m.Range(func(k, v any) bool {
		return fnc(k.(K), v.(V))
	})
}
