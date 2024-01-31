fn main() {
    println!("{:?}", two_sum(vec![1, 2], 3));
    println!("{:?}", two_sum(vec![-11, 2, 11], 0));
    println!("{:?}", two_sum(vec![-11, 2, 11, 22], 0));
}

fn two_sum(nums: Vec<i32>, target: i32) -> Vec<usize> {
    let mut left = 0;
    let mut right = nums.len() - 1;
    
    loop {
        if nums[left] + nums[right] == target {
            return vec![left+1, right+1];
        } else if nums[left] + nums[right] > target {
            right -= 1;
        } else {
            left += 1;
        }
    }

}