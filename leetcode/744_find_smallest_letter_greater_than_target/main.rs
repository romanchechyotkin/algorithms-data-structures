fn main() {
    println!("{}", next_greatest_letter(vec!['c', 'f', 'j'], 'a'));
    println!("{}", next_greatest_letter(vec!['c', 'f', 'j'], 'c'));
    println!("{}", next_greatest_letter(vec!['x', 'x', 'y', 'y'], 'z'));
}

fn next_greatest_letter(letters: Vec<char>, target: char) -> char {
    let mut low = 0;
    let mut high = letters.len() as i32 - 1;
    let mut mid = 0; 

    while low <= high {
        mid = low + (high - low)/2;

        if letters[mid as usize] <= target {
            low = mid + 1;
        } else {
            high = mid - 1;
        }
    } 

    if low >= letters.len() as i32 {
        return letters[0];
    }

    return letters[low as usize];
}