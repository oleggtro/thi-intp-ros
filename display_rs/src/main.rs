use chrono::{TimeDelta, Utc};
use futures::{executor::LocalPool, future, stream::StreamExt, task::LocalSpawnExt};
use r2r::elevator_msgs;
use r2r::QosProfile;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let ctx = r2r::Context::create()?;
    let mut node = r2r::Node::create(ctx, "display_rs", "")?;
    // let mut sub =
    //     node.subscribe::<elevator_msgs::msg::ElevatorStatus>("/topic", QosProfile::default())?;
    // let p = node
    //     .create_publisher::<elevator_msgs::msg::ElevatorStatus>("/topic2", QosProfile::default())?;

    let mut pool = LocalPool::new();
    let spawner = pool.spawner();

    // task that every other time forwards message to topic2
    /* spawner.spawn_local(async move {
        let mut x: i32 = 0;
        loop {
            match sub.next().await {
                Some(msg) => {
                    if x % 2 == 0 {
                        p.publish(&r2r::elevator_msgs::msg::ElevatorStatus {
                            going_up: msg.going_up,
                            going_down: msg.going_down,
                            current_floor: msg.current_floor,
                        })
                        .unwrap();
                    }
                }
                None => break,
            }
            x += 1;
        }
    })?; */

    let last_update_received = Utc::now();
    let mut worker = node.subscribe::<elevator_msgs::msg::ElevatorStatus>(
        "/elevatorstatus",
        QosProfile::default(),
    )?;

    spawner.spawn_local(async move {
        let mut elevator_paused = true;

        let mut paused_counter = 0;
        let mut print_msg = true;

        loop {
            match worker.next().await {
                Some(msg) => {
                    let going_up = msg.going_up;
                    let going_down = msg.going_down;

                    let dir = {
                        let rs = if going_up && !going_down {
                            print_msg = true;
                            elevator_paused = false;
                            paused_counter = 0;
                            "Hoch"
                        } else if !going_up && going_down {
                            print_msg = true;
                            elevator_paused = false;
                            paused_counter = 0;
                            "Runter"
                        } else if !going_up && !going_down {
                            print_msg = false;
                            elevator_paused = true;
                            paused_counter += 1;
                            "Stehend"
                        } else {
                            print_msg = true;
                            "Fehler"
                        };
                        rs
                    };

                    if print_msg == true {
                        println!("Aktuelle Etage: {}, Richtung: {}", msg.current_floor, dir);
                    } else if paused_counter > 10 {
                        println!("Aktuelle Etage: {}, Aufzug steht.", msg.current_floor);
                        paused_counter = 0;
                    }
                    //println!("topic2: new msg: {}", msg.expect("deserialization error"));
                }
                None => {
                    if Utc::now()
                        .signed_duration_since(&last_update_received)
                        .gt(&TimeDelta::seconds(10))
                    {
                        // check if elevator currently isn't paused => last we checked it ran
                        if elevator_paused == false {
                            println!("last update is more than 30s old");
                            elevator_paused = true;
                        }
                        println!("elevator is paused, but im outside of the paused check");
                    }
                }
            }
        }
    })?;

    /* // for sub2 we just print the data
    let sub2 = node.subscribe_untyped(
        "/elevatorstatus",
        "elevator_msgs/msg/ElevatorStatus",
        QosProfile::default(),
    )?;
    spawner.spawn_local(async move {
        sub2.for_each(|msg| {
            //let going_up = msg.expect("deserialization error").going_up;
            //let going_down = msg.expect("deserialization error").going_down;
            //let current_floor = msg.expect("deserialization error").current_floor;

            //let dir = {
            //    if going_up && !going_down {
            //        "Hoch"
            //    } else if !going_up && going_down {
            //        "Runter"
            //    } else if !going_up && !going_down {
            //        "Stehend"
            //    } else {
            //        "Fehler"
            //    }
            //
            //};
            //println!("Aktuelle Etage: {}, Richtung: {}", msg.current_floor, dir);
            //println!("topic2: new msg: {}", msg.expect("deserialization error"));
            future::ready(())
        })
        .await
    })?; */

    loop {
        node.spin_once(std::time::Duration::from_millis(100));
        pool.run_until_stalled();
    }
}
