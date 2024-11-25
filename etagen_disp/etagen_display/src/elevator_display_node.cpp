#include "rclcpp/rclcpp.hpp"
#include "std_msgs/msg/u_int8.hpp"
#include "std_msgs/msg/bool.hpp"
#include "std_msgs/msg/string.hpp"
#include "etagen_display/msg/elevator_status.hpp"

using namespace std::chrono_literals;

class ElevatorDisplayNode : public rclcpp::Node {
public:
    ElevatorDisplayNode()
    : Node("elevator_display_node"), current_floor_(0), direction_up_(false), direction_down_(false) {
        // Subscriber für das ElevatorStatus-Topic
        elevator_status_subscriber_ = this->create_subscription<std_msgs::msg::UInt8>(
            "ElevatorStatus/current_floor", 10, std::bind(&ElevatorDisplayNode::floor_callback, this, std::placeholders::_1));
        direction_up_subscriber_ = this->create_subscription<std_msgs::msg::Bool>(
            "ElevatorStatus/going_up", 10, std::bind(&ElevatorDisplayNode::direction_up_callback, this, std::placeholders::_1));
        direction_down_subscriber_ = this->create_subscription<std_msgs::msg::Bool>(
            "ElevatorStatus/going_down", 10, std::bind(&ElevatorDisplayNode::direction_down_callback, this, std::placeholders::_1));

        status_msg_subscriber_ = this->create_subscription<elevator_display::msg::ElevatorStatus>(
            "elevatorstatus", 10, std::bind(&ElevatorDisplayNode::status_callback, this, std::placeholders::_1)
        );



        // Publisher für Timeout-Nachricht
        timeout_publisher_ = this->create_publisher<std_msgs::msg::String>("timeout", 10);

        // Timer für Timeout-Überprüfung
        timeout_timer_ = this->create_wall_timer(1s, std::bind(&ElevatorDisplayNode::check_timeout, this));

        // Initialisiere den Zeitpunkt der letzten Aktualisierung
        last_update_time_ = this->now();
    }

private:
    void floor_callback(const std_msgs::msg::UInt8::SharedPtr msg) {
        current_floor_ = msg->data;
        last_update_time_ = this->now();  // Aktualisiere die Zeit der letzten Bewegung

	std::string direction_text = get_direction_text(); //holt die Richtung
        RCLCPP_INFO(this->get_logger(), "Aktuelle Etage: %d, Richtung: %s", current_floor_, direction_text.c_str());
    }

    void direction_up_callback(const std_msgs::msg::Bool::SharedPtr msg) {
        direction_up_ = msg->data;
        last_update_time_ = this->now();  // Aktualisiere die Zeit der letzten Bewegung
    }

    void direction_down_callback(const std_msgs::msg::Bool::SharedPtr msg) {
        direction_down_ = msg->data;
        last_update_time_ = this->now();  // Aktualisiere die Zeit der letzten Bewegung
    }

    void status_callback(const etagen_display::msg::ElevatorStatus::SharedPtr msg) {
        
    }

    void check_timeout() {
        // Zeit seit der letzten Aktualisierung berechnen
        auto time_since_last_update = (this->now() - last_update_time_).seconds();

        // Wenn 30 Sekunden ohne Aktualisierung vergangen sind, eine Timeout-Nachricht veröffentlichen
        if (time_since_last_update >= 30.0) {
            auto timeout_msg = std_msgs::msg::String();
            timeout_msg.data = "Aufzug steht seit 30 Sekunden still";
            timeout_publisher_->publish(timeout_msg);

            RCLCPP_WARN(this->get_logger(), "%s", timeout_msg.data.c_str());

            // Setze last_update_time_ zurück, um die Nachricht nur einmal auszugeben
            last_update_time_ = this->now();
        }
    }

    std::string get_direction_text() {
        if (direction_up_) {
            return "aufwärts";
        } else if (direction_down_) {
            return "abwärts";
        } else {
            return "stillstehend";
        }
    }

    // Subscriber und Publisher
    rclcpp::Subscription<std_msgs::msg::UInt8>::SharedPtr elevator_status_subscriber_;
    rclcpp::Subscription<std_msgs::msg::Bool>::SharedPtr direction_up_subscriber_;
    rclcpp::Subscription<std_msgs::msg::Bool>::SharedPtr direction_down_subscriber_;
    rclcpp::Publisher<std_msgs::msg::String>::SharedPtr timeout_publisher_;
    rclcpp::TimerBase::SharedPtr timeout_timer_;

    // Variablen für die aktuelle Etage, Richtung und letzte Aktualisierungszeit
    uint8_t current_floor_;
    bool direction_up_;
    bool direction_down_;
    rclcpp::Time last_update_time_;
};

int main(int argc, char * argv[]) {
    rclcpp::init(argc, argv);
    auto node = std::make_shared<ElevatorDisplayNode>();
    rclcpp::spin(node);
    rclcpp::shutdown();
    return 0;
}
