def bubble_sort(nums: list[int]) -> list[int]:
    for i in range(len(nums)):
        for j in range(len(nums)-(i+1)):
            left = nums[j]
            right = nums[j+1]

            if left > right:
                nums[j] = right
                nums[j + 1] = left
    
    return nums

print(bubble_sort([1, 2, 3, 7, 8, 9, 23, 4, 5]))
print(bubble_sort([7, 8, 1, 23, 9, 23, 4, 5]))
print(bubble_sort([9, 8, 7, 6, 4, 3, 1]))