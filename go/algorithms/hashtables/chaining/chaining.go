package hashtable

const SMALL_PRIME   = 31
const INIT_CAPACITY = 4

type Entry struct {
	Key string
	Value interface{}
	Next *Entry
}

func (e *Entry) Get(key string) interface{} {
	if e.Key == key {
		return e.Value
	}
	return e.Next.Get(key)
}

func (e *Entry) Put(k string, v interface{}) bool {
	switch e.Key {
	case k:
		e.Value = v
		break
	case "":
		e.Key = k
		e.Value = v
		break
	default:
		if e.Next == nil {
			e.Next = &Entry{}
		}
		return e.Next.Put(k, v)
	}
	if e.Value == nil {
		e.Key = ""
		return false
	}
	return true
}

type HashTable struct {
	Entries []*Entry
	m int
	n int
}

func NewHashTable() *HashTable {
	return &HashTable{ make([]*Entry, INIT_CAPACITY), INIT_CAPACITY, 0 }
}

func (h *HashTable) Size() int {
	return h.n
}

func (h *HashTable) hash(key string) int {
	var hash int
	for _, v := range key {
		hash = ((SMALL_PRIME * hash) + int(v)) % h.m
	}
	hash = (hash & 0x7fffffff) % h.m
	if h.Entries[hash] == nil {
		h.Entries[hash] = &Entry{}
	}
	return hash
}

func (h *HashTable) Get(key string) interface{} {
	i := h.hash(key)
	return h.Entries[i].Get(key)
}

func (h *HashTable) Put(k string, v interface{}) {
	i := h.hash(k)
	if h.Entries[i].Put(k, v) {
		h.n++
	}
}

func (h *HashTable) Del(key string) {
	i := h.hash(key)
	if h.Entries[i] != nil {
		h.n--
		h.Entries[i].Put(key, nil)
	}
}
