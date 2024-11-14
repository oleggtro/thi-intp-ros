#include "rclcpp/rclcpp.hpp"
#include "std_msgs/msg/u_int8.hpp"
#include "std_msgs/msg/bool.hpp"
#include "std_msgs/msg/string.hpp"

using namespace std::chrono_literals;

class ElevatorDisplayNode : public rclcpp::Node {
public:
    ElevatorDisplayNode()
    : Node("elevator_display_node"), current_floor_(0), direction_up_(true) {
        // Subscriber für das floor_current Topic
        floor_subscriber_ = this->create_subscription<std_msgs::msg::UInt8>(
            "floor_current/floor", 10, std::bind(&ElevatorDisplayNode::floor_callback, this, std::placeholders::_1));
        direction_subscriber_ = this->create_subscription<std_msgs::msg::Bool>(
            "floor_current/direction", 10, std::bind(&ElevatorDisplayNode::direction_callback, this, std::placeholders::_1));

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

        RCLCPP_INFO(this->get_logger(), "Aktuelle Etage: %d", current_floor_);
    }

    void direction_callback(const std_msgs::msg::Bool::SharedPtr msg) {
        direction_up_ = msg->data;
        last_update_time_ = this->now();  // Aktualisiere die Zeit der letzten Bewegung

        std::string direction_text = direction_up_ ? "aufwärts" : "abwärts";
        RCLCPP_INFO(this->get_logger(), "Richtung: %s", direction_text.c_str());
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

    // Subscriber und Publisher
    rclcpp::Subscription<std_msgs::msg::UInt8>::SharedPtr floor_subscriber_;
    rclcpp::Subscription<std_msgs::msg::Bool>::SharedPtr direction_subscriber_;
    rclcpp::Publisher<std_msgs::msg::String>::SharedPtr timeout_publisher_;
    rclcpp::TimerBase::SharedPtr timeout_timer_;

    // Variablen für die aktuelle Etage, Richtung und letzte Aktualisierungszeit
    uint8_t current_floor_;
    bool direction_up_;
    rclcpp::Time last_update_time_;
};

int main(int argc, char * argv[]) {
    rclcpp::init(argc, argv);
    auto node = std::make_shared<ElevatorDisplayNode>();
    rclcpp::spin(node);
    rclcpp::shutdown();
    return 0;
}
