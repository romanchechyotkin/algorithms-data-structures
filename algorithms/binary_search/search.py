def binary_search(nums, val) -> bool:
    low = 0
    high = len(nums) - 1 

    while low <= high:
        mid = low + (high-low)//2
        guess = nums[mid]
        if guess == val:
            return True
        elif guess < val:
            low = mid + 1
        else: 
            high = mid - 1

    return False    

print(binary_search([1, 2, 4, 6, 7, 11, 32, 44], 5))
print(binary_search([1, 2, 4, 6, 7, 11, 32, 44], 1))
print(binary_search([1, 2, 4, 6, 7, 11, 32, 44], 7))
print(binary_search([1, 2, 4, 6, 7, 11, 32, 44], 44))