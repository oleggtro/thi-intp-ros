import rclpy
import re
import std_msgs.msg
from elevator_msgs.msg import CallToFloor

from rclpy.node import Node

#from std_msgs.msg import String



class CallButton(Node):

    def __init__(self, floor):
        super().__init__('call_button')
        self.floor = floor
        self.publisher_ = self.create_publisher(CallToFloor, 'CallToFloor', 10) # Versendeter Wert, Topic, Warteschlange wenn Subscriber zu langsam ist

    def press_button(self, going_up):
        msg = CallToFloor()
        msg.floor = self.floor
        msg.going_up = going_up
        msg.going_down = not going_up
        self.publisher_.publish(msg)# maybe publish two to three times every call
        going = ""
        if going_up:
            going = "up"
        else:
            going = "down"
        self.get_logger().info("Call elevator to floor " + self.floor + " going " + going)


def main(args=None):
    rclpy.init(args=args)

    false_input = True
    floor = 0
    input1 = ''

    print("Version: 0.1.1")

    while(false_input):
        print("Enter the floor of this panel: ")
        input1 = '1' #Input()
        if re.search("[0-9]", input1):
            floor = re.search("[0-9]", input1).group()
            false_input = False

    call_button = CallButton(floor)
    going_up = True

    #while(True):
    print("Floor " + call_button.floor)
    input = 'up' #input("Input 'down' or 'up' to call elevator: ")
    if input == "down" or input == "up":
        if input == "down":
            going_up = False
        call_button.press_button(going_up)


    # Destroy the node explicitly
    # (optional - otherwise it will be done automatically
    # when the garbage collector destroys the node object)
    minimal_publisher.destroy_node()
    rclpy.shutdown()


if __name__ == '__main__':
    main()
