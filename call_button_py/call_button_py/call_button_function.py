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

    def press_button_simple(self):
        msg = CallToFloor()
        msg.floor = int(self.floor)
        msg.going_up = False
        msg.going_down = False
        self.publisher_.publish(msg) # maybe publish two to three times every call
        self.get_logger().info("Call elevator to floor " + self.floor)

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

def main(args=None):
    rclpy.init(args=args)

    false_input = True
    floor = 0
    user_str = ''

    print("Version: 0.1.3")

    while false_input:
        print("Enter the floor of this panel: ")
        user_str = input()
        # Modify the regex to capture an optional minus sign and digits
        match = re.fullmatch(r"(-?\d+)", user_str)  # This will match an optional "-" followed by one or more digits
        if match:
            floor = str(int(match.group()))  # Convert the matched string to an integer
            false_input = False  # Set the flag to False to exit the loop
            print(f"You entered floor: {floor}")
        else:
            print("Invalid input. Please enter a valid floor number.") 

    call_button = CallButton(floor)
    going_up = True

    while(True):
        print("Floor " + call_button.floor)
        print("Press ENTER to call elevator")
        input()
        call_button.press_button_simple()


    # Destroy the node explicitly
    # (optional - otherwise it will be done automatically
    # when the garbage collector destroys the node object)
    call_button.destroy_node()
    rclpy.shutdown()


if __name__ == '__main__':
    main()
