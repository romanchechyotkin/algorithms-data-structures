fn main() {
    println!("{}", single_number(vec![1, 2, 3, 3, 2]));
    println!("{}", single_number(vec![1, 2, 0, 2, 1]));
    println!("{}", single_number(vec![3, 1, 0, 0, 1]));
}

fn single_number(nums: Vec<usize>) -> usize {
    let mut xor: usize = 0;

    for n in nums{
        xor ^= n;
    }

    xor
}