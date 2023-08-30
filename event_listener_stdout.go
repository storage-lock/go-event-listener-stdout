package events_listener_stdout

import (
	"context"
	"fmt"
	"github.com/storage-lock/go-events"
)

// EventListenerStdout 把收到的事件打印到标准输出
type EventListenerStdout struct {
}

var _ events.Listener = &EventListenerStdout{}

func NewEventListenerStdout() *EventListenerStdout {
	return &EventListenerStdout{}
}

const Name = "event-listener-stdout"

func (x *EventListenerStdout) Name() string {
	return Name
}

func (x *EventListenerStdout) On(ctx context.Context, e *events.Event) {
	if e == nil {
		return
	}
	fmt.Println(e.ToJsonString())
}
