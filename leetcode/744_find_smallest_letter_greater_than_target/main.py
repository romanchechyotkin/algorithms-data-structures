def next_greatest_letter(letters: list[str], target: str) -> str:
    low = 0
    high = len(letters) -1

    while low <= high:
        mid = low + (high-low)//2
        guess = letters[mid]

        if guess <= target:
            low = mid + 1 
        else:
            high = mid - 1    

    if low >= len(letters):
        return letters[0]

    return letters[low]

print(next_greatest_letter(["c","f","j"], "a"))
print(next_greatest_letter(["c","f","j"], "c"))
print(next_greatest_letter(["x","x","y","y"], "z"))
