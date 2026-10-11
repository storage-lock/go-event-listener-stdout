package events_listener_stdout

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/storage-lock/go-events"
)

func TestStdoutListener(t *testing.T) {
	listener := NewEventListenerStdout()
	if listener.Name() != Name {
		t.Fatal(listener.Name())
	}
	listener.On(context.Background(), nil)
	old := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = writer
	listener.On(context.Background(), events.NewEvent("lock"))
	_ = writer.Close()
	os.Stdout = old
	output, err := io.ReadAll(reader)
	if err != nil || !bytes.Contains(output, []byte("lock")) {
		t.Fatalf("output=%s err=%v", output, err)
	}
}
