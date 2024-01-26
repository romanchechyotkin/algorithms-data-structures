def search(nums, val):
    for num in nums:
        if num == val:
            return True

    return False    

print(search([1, 2, 3, 4], 5))
print(search([1, 2, 3, 4], 3))