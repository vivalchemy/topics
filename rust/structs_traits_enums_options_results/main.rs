#[allow(dead_code)]
trait Shape {
    fn area(&self) -> f64;
}

#[allow(dead_code)]
struct Rectangle {
    width: u32,
    height: u32,
}

#[allow(dead_code)]
impl Shape for Rectangle {
    fn area(&self) -> f64 {
        (self.width * self.height) as f64
    }
}

#[allow(dead_code)]
struct Circle {
    radius: u32,
}

#[allow(dead_code)]
impl Shape for Circle {
    fn area(&self) -> f64 {
        std::f64::consts::PI * (self.radius * self.radius) as f64
    }
}

#[allow(dead_code)]
enum ShapeEnum {
    Circle(Circle),
    Rectangle(Rectangle),
}

#[allow(dead_code)]
enum NormalShapeEnum {
    Circle(usize),
    Cir { radius: usize },
    Rectangle(usize, usize),
    Rect { height: usize, width: usize },
}

impl Shape for ShapeEnum {
    fn area(&self) -> f64 {
        match self {
            ShapeEnum::Circle(c) => c.area(),
            ShapeEnum::Rectangle(r) => r.area(),
        }
    }
}

#[allow(dead_code)]
fn calculate_area(s: &impl Shape) -> f64 {
    return s.area();
}

#[allow(dead_code)]
fn find_a_in_string(s: &str) -> Result<String, String> {
    for (index, char) in s.chars().enumerate() {
        if char == 'a' {
            return Ok(format!("The character a exists in the string at {}", index));
        }
    }

    return Err(String::from("The character a was not found"));
}

#[allow(dead_code)]
fn test_optional_enums(s: &str) -> Option<String> {
    if s == "Some" {
        return Some(String::from("The string is equal to 'Some'"));
    }
    return None;
}

#[allow(dead_code)]
struct CustomResultType<T>{
    status: u16,
    message: T
}

#[allow(dead_code)]
enum CustomResult<T> {
    Ok(CustomResultType<T>),
    Err(CustomResultType<T>),
}


fn match_string(s: String) -> CustomResult<String>{
    if s != "This is a matching string" {
        return CustomResult::Err(CustomResultType{
            status: 404,
            message: String::from("The strings are not matching")
        })
    }
    return CustomResult::Err(CustomResultType{
        status: 200,
        message: String::from("The strings are matching")
    })

}

/*
 * MAIN
*/



fn main() {
    let rect = Rectangle {
        width: 70,
        height: 50,
    };
    println!("{} ", calculate_area(&rect)); // this is the trait way
    println!("{}", ShapeEnum::Rectangle(rect).area()); // this one way for enum

    // Another way for enums is
    let circ = NormalShapeEnum::Cir { radius: 7 };

    let area = match circ {
        NormalShapeEnum::Circle(c) => std::f64::consts::PI * (c.pow(2) as f64),
        NormalShapeEnum::Cir { radius } => std::f64::consts::PI * (radius.pow(2) as f64),
        NormalShapeEnum::Rectangle(w, h) => (w * h) as f64,
        NormalShapeEnum::Rect { height, width } => (width * height) as f64,
    };

    println!("{}", area);

    let returned_string = find_a_in_string("This is a string of text");

    match returned_string {
        Ok(s) => println!("Success: {}", s),
        Err(e) => println!("Error: {}", e),
    }

    let returned_option = test_optional_enums("Some");

    match returned_option {
        Some(val) => println!("Value returned as: {}", val),
        None => println!("No value was returned"),
    }

    let is_string_matched = match_string(String::from("This is a matching string"));

    match is_string_matched{
        CustomResult::Ok(s) => println!("The status is: {}.\nThe message is: {}", s.status, s.message),
        CustomResult::Err(s) => println!("The status is: {}.\nThe message is: {}", s.status, s.message),
    }

    println!("Functions can be declared after the main function {}",add(4,5));
}

#[allow(dead_code)]
fn add(a: usize, b: usize) -> usize {
    a + b
}

// NOTE: The enum values for result and option can be declared without their
// scope unless using a custom solution. For e.g. Option::Some(val) can also be
// used as Some(val) similarly for Result::Ok(val) can also be declared as Ok(val)
