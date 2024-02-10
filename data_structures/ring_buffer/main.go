package main

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type RingBuffer interface {
	Insert(val any)
	Emit() []*Data
}

type Buffer struct {
	buffer   []*Data
	size     int
	head     int
	tail     int
	emitTime time.Time
}

type Data struct {
	Timestamp time.Time
	Value     any
}

func main() {
	ringBuffer := NewRingBuffer(5)

	spew.Dump(ringBuffer.Emit())

	curr := 'a' - 1

	for i := 0; i < 10; i++ {
		curr++
		ringBuffer.Insert(string(curr))
	}

	fmt.Println("RES")

	spew.Dump(ringBuffer.Emit())
}

func NewRingBuffer(size int) *Buffer {
	return &Buffer{
		buffer: make([]*Data, size),
		size:   size,
		head:   0,
		tail:   -1,
	}
}

func (buf *Buffer) Insert(val any) {
	data := &Data{
		Timestamp: time.Now(),
		Value:     val,
	}

	buf.tail = (buf.tail + 1) % int(buf.size)
	buf.buffer[buf.tail] = data

	if buf.head == buf.tail {
		buf.head = (buf.head + 1) % buf.size
	}
}

func (buf *Buffer) Emit() []*Data {
	out := []*Data{}

	for {
		if buf.buffer[buf.head] != nil {
			out = append(out, buf.buffer[buf.head])
			buf.buffer[buf.head] = nil
		}
		if buf.head == buf.tail || buf.tail == -1 {
			break
		}
		buf.head = (buf.head + 1) % buf.size
	}

	return out
}