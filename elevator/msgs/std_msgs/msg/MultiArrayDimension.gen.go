// Code generated by rclgo-gen. DO NOT EDIT.

package std_msgs_msg
import (
	"unsafe"

	"github.com/PolibaX/rclgo/pkg/rclgo"
	"github.com/PolibaX/rclgo/pkg/rclgo/types"
	"github.com/PolibaX/rclgo/pkg/rclgo/typemap"
	primitives "github.com/PolibaX/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <std_msgs/msg/multi_array_dimension.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("std_msgs/MultiArrayDimension", MultiArrayDimensionTypeSupport)
	typemap.RegisterMessage("std_msgs/msg/MultiArrayDimension", MultiArrayDimensionTypeSupport)
}

type MultiArrayDimension struct {
	Label string `yaml:"label"`// label of given dimension
	Size uint32 `yaml:"size"`// size of given dimension (in type units)
	Stride uint32 `yaml:"stride"`// stride of given dimension
}

// NewMultiArrayDimension creates a new MultiArrayDimension with default values.
func NewMultiArrayDimension() *MultiArrayDimension {
	self := MultiArrayDimension{}
	self.SetDefaults()
	return &self
}

func (t *MultiArrayDimension) Clone() *MultiArrayDimension {
	c := &MultiArrayDimension{}
	c.Label = t.Label
	c.Size = t.Size
	c.Stride = t.Stride
	return c
}

func (t *MultiArrayDimension) CloneMsg() types.Message {
	return t.Clone()
}

func (t *MultiArrayDimension) SetDefaults() {
	t.Label = ""
	t.Size = 0
	t.Stride = 0
}

func (t *MultiArrayDimension) GetTypeSupport() types.MessageTypeSupport {
	return MultiArrayDimensionTypeSupport
}

// MultiArrayDimensionPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type MultiArrayDimensionPublisher struct {
	*rclgo.Publisher
}

// NewMultiArrayDimensionPublisher creates and returns a new publisher for the
// MultiArrayDimension
func NewMultiArrayDimensionPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*MultiArrayDimensionPublisher, error) {
	pub, err := node.NewPublisher(topic_name, MultiArrayDimensionTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &MultiArrayDimensionPublisher{pub}, nil
}

func (p *MultiArrayDimensionPublisher) Publish(msg *MultiArrayDimension) error {
	return p.Publisher.Publish(msg)
}

// MultiArrayDimensionSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type MultiArrayDimensionSubscription struct {
	*rclgo.Subscription
}

// MultiArrayDimensionSubscriptionCallback type is used to provide a subscription
// handler function for a MultiArrayDimensionSubscription.
type MultiArrayDimensionSubscriptionCallback func(msg *MultiArrayDimension, info *rclgo.MessageInfo, err error)

// NewMultiArrayDimensionSubscription creates and returns a new subscription for the
// MultiArrayDimension
func NewMultiArrayDimensionSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback MultiArrayDimensionSubscriptionCallback) (*MultiArrayDimensionSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg MultiArrayDimension
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, MultiArrayDimensionTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &MultiArrayDimensionSubscription{sub}, nil
}

func (s *MultiArrayDimensionSubscription) TakeMessage(out *MultiArrayDimension) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneMultiArrayDimensionSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneMultiArrayDimensionSlice(dst, src []MultiArrayDimension) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var MultiArrayDimensionTypeSupport types.MessageTypeSupport = _MultiArrayDimensionTypeSupport{}

type _MultiArrayDimensionTypeSupport struct{}

func (t _MultiArrayDimensionTypeSupport) New() types.Message {
	return NewMultiArrayDimension()
}

func (t _MultiArrayDimensionTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__MultiArrayDimension
	return (unsafe.Pointer)(C.std_msgs__msg__MultiArrayDimension__create())
}

func (t _MultiArrayDimensionTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__MultiArrayDimension__destroy((*C.std_msgs__msg__MultiArrayDimension)(pointer_to_free))
}

func (t _MultiArrayDimensionTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*MultiArrayDimension)
	mem := (*C.std_msgs__msg__MultiArrayDimension)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.label), m.Label)
	mem.size = C.uint32_t(m.Size)
	mem.stride = C.uint32_t(m.Stride)
}

func (t _MultiArrayDimensionTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*MultiArrayDimension)
	mem := (*C.std_msgs__msg__MultiArrayDimension)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.Label, unsafe.Pointer(&mem.label))
	m.Size = uint32(mem.size)
	m.Stride = uint32(mem.stride)
}

func (t _MultiArrayDimensionTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__MultiArrayDimension())
}

type CMultiArrayDimension = C.std_msgs__msg__MultiArrayDimension
type CMultiArrayDimension__Sequence = C.std_msgs__msg__MultiArrayDimension__Sequence

func MultiArrayDimension__Sequence_to_Go(goSlice *[]MultiArrayDimension, cSlice CMultiArrayDimension__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MultiArrayDimension, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		MultiArrayDimensionTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func MultiArrayDimension__Sequence_to_C(cSlice *CMultiArrayDimension__Sequence, goSlice []MultiArrayDimension) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.std_msgs__msg__MultiArrayDimension)(C.malloc(C.sizeof_struct_std_msgs__msg__MultiArrayDimension * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		MultiArrayDimensionTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func MultiArrayDimension__Array_to_Go(goSlice []MultiArrayDimension, cSlice []CMultiArrayDimension) {
	for i := 0; i < len(cSlice); i++ {
		MultiArrayDimensionTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func MultiArrayDimension__Array_to_C(cSlice []CMultiArrayDimension, goSlice []MultiArrayDimension) {
	for i := 0; i < len(goSlice); i++ {
		MultiArrayDimensionTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
