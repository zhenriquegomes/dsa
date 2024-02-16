def binary_search(nums: list[int], target: int) -> int:
    begin = 0
    end = len(nums) - 1
    while begin <= end:
        mid = (begin + end) // 2
        if nums[mid] > target:
            end = mid - 1
        elif nums[mid] < target:
            begin = mid + 1
        else:
            return mid
    return -1