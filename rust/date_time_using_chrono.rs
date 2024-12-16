use chrono::{Local, Utc};

fn main() {
    let global_now = Utc::now();
    println!(
        "Normal: {}\nFormatted: {}",
        global_now,
        global_now.format("%d/%m/%y %H:%M:%S")
    );

    let local_now = Local::now();
    println!(
        "Normal: {}\nFormatted: {}",
        local_now,
        local_now.format("%d/%m/%y %H:%M:%S")
    )
}
