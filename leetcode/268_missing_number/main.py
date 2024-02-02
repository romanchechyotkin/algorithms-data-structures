def missing_num(nums: list[int]) -> int:
    n = len(nums)

    sum = 0
    for i in range(n+1):
        sum += i

    for num in nums:
        sum -= num

    return sum     


print(missing_num([0, 1, 3, 4]))
print(missing_num([0, 4, 1, 3]))
print(missing_num([2, 4, 1, 3]))