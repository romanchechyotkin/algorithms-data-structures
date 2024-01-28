use std::collections::HashMap;

fn main() {
    println!("{}", is_valid("{}()[]".to_string()));
    println!("{}", is_valid("{}(]".to_string()));
    println!("{}", is_valid("({[]})".to_string()));
}

fn is_valid(s: String) -> bool {
    let mut hashmap = HashMap::new();
    hashmap.insert("}".to_string(), "{".to_string());
    hashmap.insert("]".to_string(), "[".to_string());
    hashmap.insert(")".to_string(), "(".to_string());
    
    let mut stack = Vec::new(); 

    for ch in s.chars() {
        if let Some(v) = hashmap.get(&ch.to_string()) {
            if stack.len() > 0 && stack[stack.len()-1] == *v {
                stack.pop();
                continue;
            } else {
                return false;    
            }
            
        } else {
            stack.push(ch.to_string());
        }

    } 

    stack.len() == 0
}