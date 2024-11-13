import rclpy
import re

from rclpy.node import Node

from std_msgs.msg import String



class CallButton(Node):

    def __init__(self, floor):
        super().__init__('call_button')
        self.floor = floor
        self.publisher_ = self.create_publisher(int, 'call_elevator', 10) # Versendeter Wert, Topic, Warteschlange wenn Subscriber zu langsam ist

    def press_button(self):
        self.publisher_.publish(self.floor)# maybe publish two to three times every call
        self.get_logger().info("Call to elevator to floor " + floor)


def main(args=None):
    rclpy.init(args=args)

    false_input = True
    floor = 0

    while(false_input):
        floor = input("Enter the floor of this panel: ")
        if re.search("[0-20]", floor):
            floor = re.search("[0-20]", floor).group()
            false_input = False

    call_button = CallButton(floor)

    while(True):
        print("Floor " + call_button.floor)
        input("Input ENTER to call the elevator: ")
        call_button.press_button()


    # Destroy the node explicitly
    # (optional - otherwise it will be done automatically
    # when the garbage collector destroys the node object)
    minimal_publisher.destroy_node()
    rclpy.shutdown()


if __name__ == '__main__':
    main()
