from math import sqrt, floor

def two_crystal_balls(arr) -> int:
    jmpAmount = floor(sqrt(len(arr)))

    i = jmpAmount
    for j in range(i, len(arr), jmpAmount):
        if arr[j]:
            break
        i += jmpAmount

    i -= jmpAmount

    for j in range(i, len(arr)):
        if arr[j]:
            return j

    return -1

print(two_crystal_balls([False, False, True]))
print(two_crystal_balls([True, True, True]))
print(two_crystal_balls([False, False, False, False]))
print(two_crystal_balls([False, False, False, True]))
