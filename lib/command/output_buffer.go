package command

type OutputBuffer struct {
	data []byte
}

// Write adds new bytes to the OutputBuffer.
func (b *OutputBuffer) Write(newData []byte) {
	b.data = append(b.data, newData...)
}

// Read returns and new output beyond 'current' and the current length of the
// output.
func (b *OutputBuffer) Read(current int) (string, int) {
	size := len(b.data)
	if b.HasNew(current) {
		return string(b.data[current:]), size
	}
	return "", size
}

// HasNew is used to determine if any new output has been added beyond the
// "current" index.
func (b *OutputBuffer) HasNew(current int) bool {
	if len(b.data) > current {
		return true
	}
	return false
}
