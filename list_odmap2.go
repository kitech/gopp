package gopp

// /////////////
type orderedMap[K comparable, V any] struct {
	kv map[K]*odmapElement[K, V]
	ll list[K, V]
}

func orderedMapNew[K comparable, V any]() *orderedMap[K, V] {
	return &orderedMap[K, V]{
		kv: make(map[K]*odmapElement[K, V]),
	}
}

// Get returns the value for a key. If the key does not exist, the second return
// parameter will be false and the value will be nil.
func (m *orderedMap[K, V]) Get(key K) (value V, ok bool) {
	v, ok := m.kv[key]
	if ok {
		value = v.Value
	}

	return
}

// Set will set (or replace) a value for a key. If the key was new, then true
// will be returned. The returned value will be false if the value was replaced
// (even if the value was the same).
func (m *orderedMap[K, V]) Set(key K, value V) bool {
	_, alreadyExist := m.kv[key]
	if alreadyExist {
		m.kv[key].Value = value
		return false
	}

	element := m.ll.PushBack(key, value)
	m.kv[key] = element
	return true
}

// GetOrDefault returns the value for a key. If the key does not exist, returns
// the default value instead.
func (m *orderedMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	if value, ok := m.kv[key]; ok {
		return value.Value
	}

	return defaultValue
}

// GetElement returns the element for a key. If the key does not exist, the
// pointer will be nil.
func (m *orderedMap[K, V]) GetElement(key K) *odmapElement[K, V] {
	element, ok := m.kv[key]
	if ok {
		return element
	}

	return nil
}

// Len returns the number of elements in the map.
func (m *orderedMap[K, V]) Len() int {
	return len(m.kv)
}

// Keys returns all of the keys in the order they were inserted. If a key was
// replaced it will retain the same position. To ensure most recently set keys
// are always at the end you must always Delete before Set.
func (m *orderedMap[K, V]) Keys() (keys []K) {
	keys = make([]K, 0, m.Len())
	for el := m.Front(); el != nil; el = el.Next() {
		keys = append(keys, el.Key)
	}
	return keys
}

// Delete will remove a key from the map. It will return true if the key was
// removed (the key did exist).
func (m *orderedMap[K, V]) Delete(key K) (didDelete bool) {
	element, ok := m.kv[key]
	if ok {
		m.ll.Remove(element)
		delete(m.kv, key)
	}

	return ok
}

// Front will return the element that is the first (oldest Set element). If
// there are no elements this will return nil.
func (m *orderedMap[K, V]) Front() *odmapElement[K, V] {
	return m.ll.Front()
}

// Back will return the element that is the last (most recent Set element). If
// there are no elements this will return nil.
func (m *orderedMap[K, V]) Back() *odmapElement[K, V] {
	return m.ll.Back()
}

// Copy returns a new orderedMap with the same elements.
// Using Copy while there are concurrent writes may mangle the result.
func (m *orderedMap[K, V]) Copy() *orderedMap[K, V] {
	m2 := orderedMapNew[K, V]()
	for el := m.Front(); el != nil; el = el.Next() {
		m2.Set(el.Key, el.Value)
	}
	return m2
}

// ///
func (m *orderedMap[K, V]) Prepend(key K, value V) bool {
	_, alreadyExist := m.kv[key]
	if alreadyExist {
		m.kv[key].Value = value
		return false
	}

	element := m.ll.PushFront(key, value)
	m.kv[key] = element
	return true
}
