package circular

import (
	"errors"
)

// Buffer represents a circular buffer
type Buffer struct {
	data      []byte
	oldestIdx int
	latestIdx int
	size      int
	count     int
}

// NewBuffer creates a new circular buffer of a specified size.
func NewBuffer(size int) *Buffer {
	data := make([]byte, size)
	return &Buffer{data, 0, -1, size, 0}
}

// WriteByte writes a byte to the buffer, returns an error if the buffer is full.
func (b *Buffer) WriteByte(c byte) error {
	if b.count == b.size {
		return errors.New("full")
	}
	i := b.latestIdx + 1
	if i == b.size {
		i = 0
	}
	b.data[i] = c
	b.latestIdx = i
	b.count++
	return nil
}

// Overwrite overwrites the oldest byte in the buffer.
func (b *Buffer) Overwrite(c byte) {
	i := b.latestIdx + 1
	if i == b.size {
		i = 0
	}
	b.data[i] = c
	b.latestIdx = i
	if b.count < b.size { //If not full yet
		b.count++
		return
	}
	b.oldestIdx++
	if b.oldestIdx == b.size {
		b.oldestIdx = 0
	}
}

// ReadByte pops the latest byte from buffer.
func (b *Buffer) ReadByte() (byte, error) {
	if b.count == 0 {
		return 0, errors.New("empty")
	}
	i := b.oldestIdx
	d := b.data[i]
	b.data[i] = 0
	b.count--
	b.oldestIdx++
	if b.oldestIdx == b.size {
		b.oldestIdx = 0
	}
	return d, nil
}

// Reset clears the buffer
func (b *Buffer) Reset() {
	*b = *NewBuffer(b.size)
}
