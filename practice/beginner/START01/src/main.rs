fn main() {
    use std::io::{stdin};
    let mut s = String::new();
    stdin().read_line(&mut s).expect("No input");
    print!("{}", s);
}
