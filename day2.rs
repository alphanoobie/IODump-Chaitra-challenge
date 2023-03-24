use std::io;

fn main() {
    let stdin = io::stdin();
    let mut user_input = String::new();
    stdin.read_line(&mut user_input).unwrap();
    let char = user_input.as_bytes()[0] as char;
    let vowels: &str = "AEIOUaeiou";
    if vowels.contains(char){
        println!("vowel");
    } else {
        println!("consonant");
    }   
}