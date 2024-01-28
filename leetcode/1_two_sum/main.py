def two_sum(nums: list[int], target: int) -> list[int]:
    dct = dict()

    for i in range(len(nums)):
        cur = nums[i]
        guess = dct.get(cur)
        if guess is None:
            dct[target-cur] = i
        else: 
            return [guess, i]

    return []

print(two_sum([2, 7, 9, 0], 9))
print(two_sum([3, 2, 4], 6))
print(two_sum([3, 3], 6))