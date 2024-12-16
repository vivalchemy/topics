use std::io::{self, Read};

fn main() {
    let mut input = String::new();

    // Read all input from stdin (which could be piped)
    io::stdin().read_to_string(&mut input).unwrap();

    // Print the received piped input
    println!("Piped input:\n{}", input);
}
