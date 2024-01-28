fn main() {
    println!("{:?}", bubble_sort(vec![1, 2, 3, 4, 5, 6, 0]));
    println!("{:?}", bubble_sort(vec![21, 3, 4, 5, 0, 1, 11, 22]));
    println!("{:?}", bubble_sort(vec![9, 8, 7, 6, 5, 4, 2, 1]));
}

fn bubble_sort(mut nums: Vec<u32>) -> Vec<u32> {
    for i in 0..nums.len() {
        for j in 0..(nums.len()-(i+1)) {
            let left = nums[j]; 
            let right = nums[j+1];

            if left > right {
                nums[j] = right;
                nums[j+1] = left;
            }
        }
    }
    
    nums
}