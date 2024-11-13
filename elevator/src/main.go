package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PolibaX/rclgo/pkg/rclgo"
	greeting_msgs "oleggtro.com/intp/ros_elevator/msgs/greeting_msgs/msg"
	// "oleggtro.com/intp/ros_elevator/msgs/std_msgs/msg"
)

func run() error {
	rclArgs, restArgs, err := rclgo.ParseArgs(os.Args[1:])
	if err != nil {
		return fmt.Errorf("failed to parse ROS args: %v", err)
	}
	if err := rclgo.Init(rclArgs); err != nil {
		return fmt.Errorf("failed to initialize rclgo: %v", err)
	}
	defer rclgo.Uninit()
	node, err := rclgo.NewNode("greeter", "")
	if err != nil {
		return fmt.Errorf("failed to create node: %v", err)
	}
	pub, err := greeting_msgs.NewHelloPublisher(node, "~/hello", nil)
	if err != nil {
		return fmt.Errorf("failed to create publisher: %v", err)
	}
	greeting := greeting_msgs.NewHello()
	if len(restArgs) > 0 {
		greeting.Name = restArgs[0]
	} else {
		greeting.Name = "gopher"
	}
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	for {
		node.Logger().Infof("Publishing greeting: %#v", greeting)
		if err := pub.Publish(greeting); err != nil {
			return fmt.Errorf("failed to publish: %v", err)
		}
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(2 * time.Second):
		}
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
