use std::collections::HashMap;

fn main() {
    println!("{}", contains_duplicate(&vec![1, 2, 3]));
    println!("{}", contains_duplicate(&vec![1, 2, 3, 1]));
    println!("{}", contains_duplicate(&vec![1, 1, 2, 3]));
}

fn contains_duplicate(nums: &Vec<u8>) -> bool {
    let mut hashmap: HashMap<u8, u8> = HashMap::new();
    
    for num in nums {
        match hashmap.get(num) {
            Some(_) => return true,
            None => {
                hashmap.insert(*num, 1)
            },
        };
    };

    false
}