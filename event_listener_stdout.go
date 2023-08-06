package events_listener_stdout

import (
	"context"
	"fmt"
	"github.com/storage-lock/go-events"
)

type EventListenerStdout struct {
}

var _ events.Listener = &EventListenerStdout{}

func NewEventListenerStdout() *EventListenerStdout {
	return &EventListenerStdout{}
}

func (x *EventListenerStdout) Name() string {
	return "event-listener-stdout"
}

func (x *EventListenerStdout) On(ctx context.Context, e *events.Event) {
	fmt.Println(e.ToJsonString())
}
