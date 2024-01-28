use std::collections::HashMap;

fn main() {
    println!("{}", is_anagram("catt".to_string(), "ratt".to_string()));
    println!("{}", is_anagram("ccac".to_string(), "aacc".to_string()));
    println!("{}", is_anagram("anagram".to_string(), "aaangrm".to_string()));
}

fn is_anagram(s: String, t: String) -> bool {
    if s.len() != t.len() {
        return false;
    }
    
    let mut hash_map = HashMap::new();

    for i in s.chars() {
        if let Some(x) = hash_map.get_mut(&i) {
            *x += 1;
            continue;
        }
        hash_map.insert(i, 1);
    } 

    for i in t.chars() {
        if let Some(x) = hash_map.get_mut(&i) {
            *x -= 1;
            continue;
        }
        return false;
    } 

    for (_, v) in hash_map.iter() {
        if *v != 0 {
            return false;
        } else {
            continue;
        }
    } 

    true
}   