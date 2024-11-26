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
        self.publisher_ = self.create_publisher(CallToFloor, 'calltofloor', 10) # Versendeter Wert, Topic, Warteschlange wenn Subscriber zu langsam ist

    def press_button(self, going_up):
        msg = CallToFloor()
        msg.floor = int(self.floor)
        msg.going_up = going_up
        msg.going_down = not going_up
        self.publisher_.publish(msg) # maybe publish two to three times every call
        going = ""
        if going_up:
            going = "up"
        else:
            going = "down"
        self.get_logger().info("Call elevator to floor " + self.floor + " going " + going)


    
def press_button_simple(self):
        msg = CallToFloor()
        msg.floor = int(self.floor)
        msg.going_up = false
        msg.going_down = false
        self.publisher_.publish(msg) # maybe publish two to three times every call
        self.get_logger().info("Call elevator to floor " + self.floor)


def main(args=None):
    rclpy.init(args=args)

    false_input = True
    floor = 0
    user_str = ''

    print("Version: 0.1.3")

    while(false_input):
        print("Enter the floor of this panel: ")
        user_str = input()
        if re.search("[0-9]", user_str):
            floor = re.search("[0-9]", user_str).group()
            false_input = False

    call_button = CallButton(floor)
    going_up = True

    while(True):
        print("Floor " + call_button.floor)
        #print("Input 'down' or 'up' to call elevator: ")
        #user_str = input().lower()
        #if user_str == "down" or user_str == "up":
        #    if user_str == "down":
        #        going_up = False
        #    call_button.press_button(going_up)
        print("Press ENTER to call elevator")


    # Destroy the node explicitly
    # (optional - otherwise it will be done automatically
    # when the garbage collector destroys the node object)
    call_button.destroy_node()
    rclpy.shutdown()


if __name__ == '__main__':
    main()
