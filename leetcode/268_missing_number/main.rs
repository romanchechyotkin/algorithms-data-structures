fn main() {
    println!("{}", missing_number(vec![1, 2, 3]));
    println!("{}", missing_number(vec![1, 2, 0]));
    println!("{}", missing_number(vec![3, 1, 0]));
}

fn missing_number(nums: Vec<usize>) -> usize {
    let n = nums.len() + 1;
    let mut sum: usize = 0;

    for i in 0..n {
        sum += i;
    }

    for i in nums {
        sum -= i;
    }

    sum
}