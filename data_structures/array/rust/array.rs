fn main() {
    println!("arrays");

    let mut arr: [u8; 3] = [1, 2, 3];
    println!("{:?}", arr);

    println!("{} {}", arr[0], arr[arr.len() - 1]);

    arr[1] = 0;
    println!("{:?}", arr);
}
