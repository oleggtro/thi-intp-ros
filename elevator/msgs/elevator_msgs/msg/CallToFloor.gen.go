// Code generated by rclgo-gen. DO NOT EDIT.

package elevator_msgs_msg
import (
	"unsafe"

	"github.com/PolibaX/rclgo/pkg/rclgo"
	"github.com/PolibaX/rclgo/pkg/rclgo/types"
	"github.com/PolibaX/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <elevator_msgs/msg/call_to_floor.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("elevator_msgs/CallToFloor", CallToFloorTypeSupport)
	typemap.RegisterMessage("elevator_msgs/msg/CallToFloor", CallToFloorTypeSupport)
}

type CallToFloor struct {
	Floor int8 `yaml:"floor"`
	GoingUp bool `yaml:"going_up"`
	GoingDown bool `yaml:"going_down"`
}

// NewCallToFloor creates a new CallToFloor with default values.
func NewCallToFloor() *CallToFloor {
	self := CallToFloor{}
	self.SetDefaults()
	return &self
}

func (t *CallToFloor) Clone() *CallToFloor {
	c := &CallToFloor{}
	c.Floor = t.Floor
	c.GoingUp = t.GoingUp
	c.GoingDown = t.GoingDown
	return c
}

func (t *CallToFloor) CloneMsg() types.Message {
	return t.Clone()
}

func (t *CallToFloor) SetDefaults() {
	t.Floor = 0
	t.GoingUp = false
	t.GoingDown = false
}

func (t *CallToFloor) GetTypeSupport() types.MessageTypeSupport {
	return CallToFloorTypeSupport
}

// CallToFloorPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type CallToFloorPublisher struct {
	*rclgo.Publisher
}

// NewCallToFloorPublisher creates and returns a new publisher for the
// CallToFloor
func NewCallToFloorPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*CallToFloorPublisher, error) {
	pub, err := node.NewPublisher(topic_name, CallToFloorTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &CallToFloorPublisher{pub}, nil
}

func (p *CallToFloorPublisher) Publish(msg *CallToFloor) error {
	return p.Publisher.Publish(msg)
}

// CallToFloorSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type CallToFloorSubscription struct {
	*rclgo.Subscription
}

// CallToFloorSubscriptionCallback type is used to provide a subscription
// handler function for a CallToFloorSubscription.
type CallToFloorSubscriptionCallback func(msg *CallToFloor, info *rclgo.MessageInfo, err error)

// NewCallToFloorSubscription creates and returns a new subscription for the
// CallToFloor
func NewCallToFloorSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback CallToFloorSubscriptionCallback) (*CallToFloorSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg CallToFloor
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, CallToFloorTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &CallToFloorSubscription{sub}, nil
}

func (s *CallToFloorSubscription) TakeMessage(out *CallToFloor) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneCallToFloorSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneCallToFloorSlice(dst, src []CallToFloor) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var CallToFloorTypeSupport types.MessageTypeSupport = _CallToFloorTypeSupport{}

type _CallToFloorTypeSupport struct{}

func (t _CallToFloorTypeSupport) New() types.Message {
	return NewCallToFloor()
}

func (t _CallToFloorTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.elevator_msgs__msg__CallToFloor
	return (unsafe.Pointer)(C.elevator_msgs__msg__CallToFloor__create())
}

func (t _CallToFloorTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.elevator_msgs__msg__CallToFloor__destroy((*C.elevator_msgs__msg__CallToFloor)(pointer_to_free))
}

func (t _CallToFloorTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*CallToFloor)
	mem := (*C.elevator_msgs__msg__CallToFloor)(dst)
	mem.floor = C.int8_t(m.Floor)
	mem.going_up = C.bool(m.GoingUp)
	mem.going_down = C.bool(m.GoingDown)
}

func (t _CallToFloorTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*CallToFloor)
	mem := (*C.elevator_msgs__msg__CallToFloor)(ros2_message_buffer)
	m.Floor = int8(mem.floor)
	m.GoingUp = bool(mem.going_up)
	m.GoingDown = bool(mem.going_down)
}

func (t _CallToFloorTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__elevator_msgs__msg__CallToFloor())
}

type CCallToFloor = C.elevator_msgs__msg__CallToFloor
type CCallToFloor__Sequence = C.elevator_msgs__msg__CallToFloor__Sequence

func CallToFloor__Sequence_to_Go(goSlice *[]CallToFloor, cSlice CCallToFloor__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]CallToFloor, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		CallToFloorTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func CallToFloor__Sequence_to_C(cSlice *CCallToFloor__Sequence, goSlice []CallToFloor) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.elevator_msgs__msg__CallToFloor)(C.malloc(C.sizeof_struct_elevator_msgs__msg__CallToFloor * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		CallToFloorTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func CallToFloor__Array_to_Go(goSlice []CallToFloor, cSlice []CCallToFloor) {
	for i := 0; i < len(cSlice); i++ {
		CallToFloorTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func CallToFloor__Array_to_C(cSlice []CCallToFloor, goSlice []CallToFloor) {
	for i := 0; i < len(goSlice); i++ {
		CallToFloorTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
