def single_num(nums: list[int]) -> int:
    xor = 0

    for n in nums:
        xor ^= n

    return xor

print(single_num([1, 2, 3, 2, 3]))
print(single_num([1, 1, 2]))