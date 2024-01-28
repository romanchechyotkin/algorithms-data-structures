def valid_anagram(s: str, t: str) -> bool:
    if len(s) != len(t):
        return False

    s_count, t_count = {}, {}

    for i in range(len(s)):
        s_count[s[i]] = s_count.get(s[i], 0) + 1
        t_count[t[i]] = t_count.get(t[i], 0) + 1

    for c in s_count:
        if s_count[c] != t_count.get(c, 0):
            return False

    return True    

print(valid_anagram("anagram", "nagaram"))    
print(valid_anagram("cat", "rat"))    
print(valid_anagram("aacc", "ccac"))    