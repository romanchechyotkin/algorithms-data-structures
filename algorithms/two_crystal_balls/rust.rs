fn main() {
    println!("{}", two_crystal_balls(&vec![false, false, true]));
    println!("{}", two_crystal_balls(&vec![false, false, false, true]));
    println!("{}", two_crystal_balls(&vec![false, false, false, false, false, false, false, false, true]));
    println!("{}", two_crystal_balls(&vec![false, false, false]));
    println!("{}", two_crystal_balls(&vec![true, true, true]));
}

fn two_crystal_balls(arr: &Vec<bool>) -> u8 {
    let jmpAmount = (arr.len() as f64).sqrt().round();
    let jmpAmount = jmpAmount as u8;
    
    let mut i = &jmpAmount;

    while *i < arr.len() as u8 {
        if arr[i] {
            break;
        } else {
            *i += jmpAmount 
        }
    }

// todo 
    jmpAmount
}