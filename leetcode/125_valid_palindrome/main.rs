fn main() {
    println!("{}", is_palindrome("".chars().collect()));
    println!("{}", is_palindrome(" ".chars().collect()));
    println!("{}", is_palindrome("A man, a plan, a canal: Panama".chars().collect()));
    println!("{}", is_palindrome("race a car".chars().collect()));
    println!("{}", is_palindrome("ronma".chars().collect()));
}

fn is_palindrome(s: Vec<char>) -> bool {
    if s.len() == 0 {
        return true;
    }

    let mut left = 0;
    let mut right = s.len() - 1;

    while left < right {

        if !s[left].is_alphanumeric() {
            left += 1;
            continue;
        }

        if !s[right].is_alphanumeric() {
            right -= 1;
            continue;
        }

        if s[right].to_lowercase().to_string() != s[left].to_lowercase().to_string() {
            return false;
        }
        
        left += 1;
        right -= 1;
    }

    true
}