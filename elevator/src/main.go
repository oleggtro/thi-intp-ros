package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PolibaX/rclgo/pkg/rclgo"
	elevator_msgs "oleggtro.com/intp/ros_elevator/msgs/elevator_msgs/msg"
)

func run() error {
	rclArgs, _, err := rclgo.ParseArgs(os.Args[1:])
	if err != nil {
		return fmt.Errorf("failed to parse ROS args: %v", err)
	}

	if err := rclgo.Init(rclArgs); err != nil {
		return fmt.Errorf("failed to initialize rclgo: %v", err)
	}
	defer rclgo.Uninit()

	node, err := rclgo.NewNode("elevator_core", "")
	if err != nil {
		return fmt.Errorf("failed to create node: %v", err)
	}
	defer node.Close()

	// Publisher für den Aufzugsstatus einrichten
	pub, err := elevator_msgs.NewElevatorStatusPublisher(node, "/elevatorstatus", nil)
	if err != nil {
		return fmt.Errorf("failed to create publisher: %v", err)
	}
	defer pub.Close()

	// Subscriber für Anfragen an bestimmte Etagen einrichten
	sub, err := elevator_msgs.NewCallToFloorSubscription(
		node,
		"/calltofloor",
		nil,
		func(msg *elevator_msgs.CallToFloor, info *rclgo.MessageInfo, err error) {
			if err != nil {
				node.Logger().Errorf("failed to receive message: %v", err)
				return
			}
			node.Logger().Infof("Received CallToFloor: %#v", msg)
			// Fahrziel und Richtung setzen
			handleFloorRequest(msg, pub)
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create subscriber: %v", err)
	}
	defer sub.Close()

	// Setzen des WaitSets und Signalbehandlung
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	ws, err := rclgo.NewWaitSet()
	if err != nil {
		return fmt.Errorf("failed to create waitset: %v", err)
	}
	defer ws.Close()
	ws.AddSubscriptions(sub.Subscription)

	// Status jede Sekunde veröffentlichen
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				status := elevator_msgs.NewElevatorStatus()
				status.CurrentFloor = currentFloor
				status.GoingUp = goingUp
				status.GoingDown = goingDown

				if err := pub.Publish(status); err != nil {
					fmt.Printf("failed to publish elevator status: %v\n", err)
				}
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return ws.Run(ctx)
}

// Globale Variablen zur Verwaltung des Aufzugszustands
var (
	currentFloor uint8 = 0
	goingUp      bool  = false
	goingDown    bool  = false
)

// Funktion zur Verarbeitung einer Etagenanforderung
func handleFloorRequest(call *elevator_msgs.CallToFloor, pub *elevator_msgs.ElevatorStatusPublisher) {
	targetFloor := call.Floor
	if targetFloor == currentFloor {
		// Nichts tun, wenn bereits im Zielgeschoss
		return
	}

	// Richtung festlegen
	goingUp = targetFloor > currentFloor
	goingDown = targetFloor < currentFloor

	// Aufzugsbewegung simulieren
	for currentFloor != targetFloor {
		if goingUp {
			currentFloor++
		} else if goingDown {
			currentFloor--
		}
		time.Sleep(1 * time.Second)
	}

	// Richtung zurücksetzen
	goingUp = false
	goingDown = false
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
