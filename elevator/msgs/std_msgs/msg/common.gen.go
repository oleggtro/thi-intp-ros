// Code generated by rclgo-gen. DO NOT EDIT.

package std_msgs_msg

/*
#cgo LDFLAGS: "-L/home/ole/intp-ros/install/greeting_msgs/lib" "-Wl,-rpath=/home/ole/intp-ros/install/greeting_msgs/lib"
#cgo LDFLAGS: "-L/home/ole/intp-ros/install/etagen_display/lib" "-Wl,-rpath=/home/ole/intp-ros/install/etagen_display/lib"
#cgo LDFLAGS: "-L/home/ole/intp-ros/install/display_rs/lib" "-Wl,-rpath=/home/ole/intp-ros/install/display_rs/lib"
#cgo LDFLAGS: "-L/home/ole/intp-ros/install/elevator_msgs/lib" "-Wl,-rpath=/home/ole/intp-ros/install/elevator_msgs/lib"
#cgo LDFLAGS: "-L/home/ole/intp-ros/install/call_button_py/lib" "-Wl,-rpath=/home/ole/intp-ros/install/call_button_py/lib"
#cgo LDFLAGS: "-L/home/ole/intp-ros/install/aufzug_call_panel/lib" "-Wl,-rpath=/home/ole/intp-ros/install/aufzug_call_panel/lib"
#cgo LDFLAGS: "-L/opt/ros/jazzy/lib" "-Wl,-rpath=/opt/ros/jazzy/lib"

#cgo LDFLAGS: -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c
#cgo LDFLAGS: -lbuiltin_interfaces__rosidl_typesupport_c -lbuiltin_interfaces__rosidl_generator_c

#cgo CFLAGS: "-I/home/ole/intp-ros/install/greeting_msgs/include/action_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/greeting_msgs/include/builtin_interfaces"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/greeting_msgs/include/example_interfaces"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/greeting_msgs/include/geometry_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/greeting_msgs/include/rcutils"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/greeting_msgs/include/rosidl_dynamic_typesupport"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/greeting_msgs/include/rosidl_runtime_c"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/greeting_msgs/include/rosidl_typesupport_interface"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/greeting_msgs/include/sensor_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/greeting_msgs/include/service_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/greeting_msgs/include/std_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/greeting_msgs/include/std_srvs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/greeting_msgs/include/test_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/greeting_msgs/include/type_description_interfaces"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/greeting_msgs/include/unique_identifier_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/greeting_msgs/include/builtin_interfaces"

#cgo CFLAGS: "-I/home/ole/intp-ros/install/greeting_msgs/include/std_msgs"

#cgo CFLAGS: "-I/home/ole/intp-ros/install/etagen_display/include/action_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/etagen_display/include/builtin_interfaces"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/etagen_display/include/example_interfaces"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/etagen_display/include/geometry_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/etagen_display/include/rcutils"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/etagen_display/include/rosidl_dynamic_typesupport"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/etagen_display/include/rosidl_runtime_c"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/etagen_display/include/rosidl_typesupport_interface"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/etagen_display/include/sensor_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/etagen_display/include/service_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/etagen_display/include/std_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/etagen_display/include/std_srvs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/etagen_display/include/test_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/etagen_display/include/type_description_interfaces"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/etagen_display/include/unique_identifier_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/etagen_display/include/builtin_interfaces"

#cgo CFLAGS: "-I/home/ole/intp-ros/install/etagen_display/include/std_msgs"

#cgo CFLAGS: "-I/home/ole/intp-ros/install/display_rs/include/action_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/display_rs/include/builtin_interfaces"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/display_rs/include/example_interfaces"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/display_rs/include/geometry_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/display_rs/include/rcutils"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/display_rs/include/rosidl_dynamic_typesupport"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/display_rs/include/rosidl_runtime_c"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/display_rs/include/rosidl_typesupport_interface"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/display_rs/include/sensor_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/display_rs/include/service_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/display_rs/include/std_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/display_rs/include/std_srvs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/display_rs/include/test_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/display_rs/include/type_description_interfaces"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/display_rs/include/unique_identifier_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/display_rs/include/builtin_interfaces"

#cgo CFLAGS: "-I/home/ole/intp-ros/install/display_rs/include/std_msgs"

#cgo CFLAGS: "-I/home/ole/intp-ros/install/elevator_msgs/include/action_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/elevator_msgs/include/builtin_interfaces"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/elevator_msgs/include/example_interfaces"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/elevator_msgs/include/geometry_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/elevator_msgs/include/rcutils"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/elevator_msgs/include/rosidl_dynamic_typesupport"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/elevator_msgs/include/rosidl_runtime_c"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/elevator_msgs/include/rosidl_typesupport_interface"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/elevator_msgs/include/sensor_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/elevator_msgs/include/service_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/elevator_msgs/include/std_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/elevator_msgs/include/std_srvs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/elevator_msgs/include/test_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/elevator_msgs/include/type_description_interfaces"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/elevator_msgs/include/unique_identifier_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/elevator_msgs/include/builtin_interfaces"

#cgo CFLAGS: "-I/home/ole/intp-ros/install/elevator_msgs/include/std_msgs"

#cgo CFLAGS: "-I/home/ole/intp-ros/install/call_button_py/include/action_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/call_button_py/include/builtin_interfaces"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/call_button_py/include/example_interfaces"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/call_button_py/include/geometry_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/call_button_py/include/rcutils"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/call_button_py/include/rosidl_dynamic_typesupport"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/call_button_py/include/rosidl_runtime_c"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/call_button_py/include/rosidl_typesupport_interface"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/call_button_py/include/sensor_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/call_button_py/include/service_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/call_button_py/include/std_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/call_button_py/include/std_srvs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/call_button_py/include/test_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/call_button_py/include/type_description_interfaces"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/call_button_py/include/unique_identifier_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/call_button_py/include/builtin_interfaces"

#cgo CFLAGS: "-I/home/ole/intp-ros/install/call_button_py/include/std_msgs"

#cgo CFLAGS: "-I/home/ole/intp-ros/install/aufzug_call_panel/include/action_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/aufzug_call_panel/include/builtin_interfaces"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/aufzug_call_panel/include/example_interfaces"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/aufzug_call_panel/include/geometry_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/aufzug_call_panel/include/rcutils"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/aufzug_call_panel/include/rosidl_dynamic_typesupport"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/aufzug_call_panel/include/rosidl_runtime_c"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/aufzug_call_panel/include/rosidl_typesupport_interface"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/aufzug_call_panel/include/sensor_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/aufzug_call_panel/include/service_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/aufzug_call_panel/include/std_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/aufzug_call_panel/include/std_srvs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/aufzug_call_panel/include/test_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/aufzug_call_panel/include/type_description_interfaces"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/aufzug_call_panel/include/unique_identifier_msgs"
#cgo CFLAGS: "-I/home/ole/intp-ros/install/aufzug_call_panel/include/builtin_interfaces"

#cgo CFLAGS: "-I/home/ole/intp-ros/install/aufzug_call_panel/include/std_msgs"

#cgo CFLAGS: "-I/opt/ros/jazzy/include/action_msgs"
#cgo CFLAGS: "-I/opt/ros/jazzy/include/builtin_interfaces"
#cgo CFLAGS: "-I/opt/ros/jazzy/include/example_interfaces"
#cgo CFLAGS: "-I/opt/ros/jazzy/include/geometry_msgs"
#cgo CFLAGS: "-I/opt/ros/jazzy/include/rcutils"
#cgo CFLAGS: "-I/opt/ros/jazzy/include/rosidl_dynamic_typesupport"
#cgo CFLAGS: "-I/opt/ros/jazzy/include/rosidl_runtime_c"
#cgo CFLAGS: "-I/opt/ros/jazzy/include/rosidl_typesupport_interface"
#cgo CFLAGS: "-I/opt/ros/jazzy/include/sensor_msgs"
#cgo CFLAGS: "-I/opt/ros/jazzy/include/service_msgs"
#cgo CFLAGS: "-I/opt/ros/jazzy/include/std_msgs"
#cgo CFLAGS: "-I/opt/ros/jazzy/include/std_srvs"
#cgo CFLAGS: "-I/opt/ros/jazzy/include/test_msgs"
#cgo CFLAGS: "-I/opt/ros/jazzy/include/type_description_interfaces"
#cgo CFLAGS: "-I/opt/ros/jazzy/include/unique_identifier_msgs"
#cgo CFLAGS: "-I/opt/ros/jazzy/include/builtin_interfaces"

#cgo CFLAGS: "-I/opt/ros/jazzy/include/std_msgs"
*/
import "C"
