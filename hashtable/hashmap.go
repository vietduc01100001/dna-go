package hashtable

const (
	// KeyNotFound is a constant indicates that
	// a key cannot be found in the HashMap.
	KeyNotFound     = 0
	defaultCapacity = 8
	upperLoadFactor = 0.75
	lowerLoadFactor = 0.2
)

type node struct {
	key   int
	value int
	next  *node
}

// HashMap is a struct for storing key-value data.
type HashMap struct {
	items []*node
	size  int
}

// NewHashMap creates a new HashMap instance.
func NewHashMap() *HashMap {
	return &HashMap{
		items: make([]*node, defaultCapacity),
	}
}

// Insert adds a new pair of key-value to the HashMap.
// If the key already exists, it overwrites the key's value.
func (h *HashMap) Insert(key int, value int) {
	if h.getLoadFactor() >= upperLoadFactor {
		h.rehash(true)
	}

	newNode := &node{
		key:   key,
		value: value,
	}

	h.size++
	index := h.hash(key)
	curNode := h.items[index]

	for curNode != nil {
		if curNode.key == key {
			curNode.value = value
			return
		}

		curNode = curNode.next
	}

	if curNode == nil {
		// first node in the linked list
		h.items[index] = newNode
		return
	}

	curNode.next = newNode
}

// Search returns the value for the given key.
// If the key cannot be found, it returns KeyNotFound.
func (h *HashMap) Search(key int) int {
	index := h.hash(key)
	curNode := h.items[index]

	for curNode != nil {
		if curNode.key == key {
			return curNode.value
		}

		curNode = curNode.next
	}

	return KeyNotFound
}

// Delete removes a key from the HashMap.
func (h *HashMap) Delete(key int) {
	if h.getLoadFactor() < lowerLoadFactor {
		h.rehash(false)
	}

	index := h.hash(key)
	curNode := h.items[index]

	var prevNode *node
	for curNode != nil {
		if curNode.key == key {
			if prevNode == nil {
				// delete the first node in the linked list
				h.items[index] = curNode.next
			} else {
				prevNode.next = curNode.next
			}
		}

		prevNode = curNode
		curNode = curNode.next
	}

	if curNode == nil {
		return
	}

	// remove the reference from this node
	curNode.next = nil
	h.size--
}

func (h *HashMap) capacity() int {
	return len(h.items)
}

func (h *HashMap) getLoadFactor() float64 {
	return float64(h.size) / float64(len(h.items))
}

func (h *HashMap) hash(key int) int {
	return key % len(h.items)
}

func (h *HashMap) rehash(isGrowSize bool) {
	currLength := len(h.items)
	var newLength int
	if isGrowSize == true {
		newLength = currLength * 2
	} else {
		newLength = currLength / 2
	}

	newArr := make([]*node, newLength)
	oldArr := h.items
	h.items = newArr
	h.size = 0

	for _, curNode := range oldArr {
		for curNode != nil {
			h.Insert(curNode.key, curNode.value)
			curNode = curNode.next
		}
	}
}
