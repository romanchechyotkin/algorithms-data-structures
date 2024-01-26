def containsDuplicate(nums: list[int]) -> bool:
    dct = dict()

    for num in nums:
        if dct.get(num) == 1:
            return True
        else:
            dct[num] = 1

    return False



print(containsDuplicate([1, 2, 3, 4, 1]))
print(containsDuplicate([1, 2, 3, 4]))
print(containsDuplicate([1,1,1,1,3,3,4,3,2,4,2])) 