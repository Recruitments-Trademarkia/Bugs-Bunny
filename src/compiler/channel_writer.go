package compiler

type ChannelWriter struct {
	channel chan string
}

// Write writes to the channel
func (cw *ChannelWriter) Write(p []byte) (n int, err error) {
	cw.channel <- string(p)
	return len(p), nil
}

// Close closes the channel
func (cw *ChannelWriter) Close() error {
	close(cw.channel)
	return nil
}

// NewChannelWriter creates a new ChannelWriter
func NewChannelWriter(c chan string) *ChannelWriter {
	return &ChannelWriter{channel: c}
}
