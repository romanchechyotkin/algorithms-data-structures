fn main() {
    println!("{:?}", binary_search(&vec![1, 3, 5, 6, 11, 45, 34], &1));
    println!("{:?}", binary_search(&vec![1, 3, 5, 6, 11, 45, 34], &45));
    println!("{:?}", binary_search(&vec![1, 3, 5, 6, 11, 45, 34], &46));
}

fn binary_search(nums: &Vec<u8>, val: &u8) -> bool {
    let mut low = 0;
    let mut high = nums.len() - 1;
    
    let mut mid = 0; 
    let mut guess: &u8 = &0; 

    while low <= high {
        mid = low + (high - low) / 2;
        guess = &nums[mid];

        if val == guess {
            return true;
        } else if guess < val {
            low = mid + 1;
        } else {
            high = mid - 1;
        }
    }

    return false;
}