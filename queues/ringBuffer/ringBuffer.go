package ringbuffer

type (
	RingBuffer[T any] struct {
		nodes []T

		len        uint32
		writeIndex uint32
		readIndex  uint32
	}
)

// Make new RingBuffer
func New[T any](size uint32) *RingBuffer[T] {
	len := roundUp(size)

	return &RingBuffer[T]{
		len:   len,
		nodes: make([]T, len),
	}
}

// Clean ring buffer and drop write/read index to 0
func (rb *RingBuffer[T]) Clean() {
	rb.writeIndex = 0
	rb.readIndex = 0

	rb.nodes = rb.nodes[:0]
}

// Cap returns the capacity of ring buffer.
func (rb *RingBuffer[T]) Cap() uint32 {
	return rb.len
}

// Puts item into ring buffer
func (rb *RingBuffer[T]) Put(item T) {
	rb.nodes[rb.writeIndex] = item

	if rb.writeIndex == rb.len {
		rb.writeIndex = 0
		return
	}

	rb.writeIndex++
}

// Get item from ring buffer
func (rb *RingBuffer[T]) Get() (item T) {
	item = rb.nodes[rb.readIndex]

	if rb.readIndex == rb.len {
		rb.readIndex = 0
		return
	}

	rb.readIndex++

	return
}

func roundUp(v uint32) uint32 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16

	return v
}
