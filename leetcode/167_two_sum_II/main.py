def twoSum(numbers: list[int], target: int) -> list[int]:
        left = 0
        right = len(numbers) - 1

        while True:
            if numbers[left] + numbers[right] == target:
                return [left+1, right+1] 
            elif numbers[left] + numbers[right] > target:
                right -= 1
            else: 
                left += 1



print(twoSum([1, 2, 4], 6))
print(twoSum([1, 2], 3))
print(twoSum([-1, 0, 1], 0))