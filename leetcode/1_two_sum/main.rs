use std::collections::HashMap;
use std::convert::TryInto;


fn main() {
    println!("{:?}", two_sum(vec![2, 7, 9], 9));
    println!("{:?}", two_sum(vec![2, 7, 4], 6));
    println!("{:?}", two_sum(vec![3, 3], 6));
}

fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut h_map = HashMap::new();

    for (i, y) in nums.iter().enumerate() {
        let x = target - y;

        if let Some(&j) = h_map.get(&x) {
            return vec![j as i32, i as i32];
        }

        h_map.insert(y, i);
    }
    
    vec![0, 0]
}