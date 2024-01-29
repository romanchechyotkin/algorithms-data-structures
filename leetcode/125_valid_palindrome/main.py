def is_palindrome(s: str) -> bool:
    s = s.lower()
    
    left = 0
    right = len(s)-1

    while left < right:
        if not s[left].isalnum():
            left += 1
            continue
        if not s[right].isalnum():
            right -= 1
            continue
        
        if s[left] != s[right]:
            return False
        
        left += 1
        right -= 1

    return True


print(is_palindrome("roma"))    
print(is_palindrome("A man, a plan, a canal: Panama"))    
print(is_palindrome("race a car"))    
print(is_palindrome(""))    
print(is_palindrome(" "))    