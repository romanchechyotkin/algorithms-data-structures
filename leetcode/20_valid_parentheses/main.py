def isValid(s: str) -> bool:
        closed = {
            "}": "{",
            ")": "(",
            "]": "[",
        }
        stack = []

        for ch in s:
            if ch in closed:
                if stack and stack[-1] == closed[ch]:
                    stack.pop()    
                else: 
                    return False
            else:    
                stack.append(ch)

        return len(stack) == 0  


print(isValid("({[]})"))
print(isValid("({])"))
print(isValid("{{}}()"))