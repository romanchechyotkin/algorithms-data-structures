fn main() {
    println!("{:?}", search(&vec![1, 2, 3], &1));
    println!("{:?}", search(&vec![1, 2, 3], &4));
}

fn search(nums: &Vec<u8>, val: &u8) -> bool {
    for num in nums {
        if num == val {
            return true;
        } else {
            continue;
        }
    }

    false
}