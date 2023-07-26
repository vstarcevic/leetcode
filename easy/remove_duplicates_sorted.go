// Remove Duplicates from Sorted Array
// Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

// Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

// Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
// Return k.

// Example 1:
// Input: nums = [1,1,2]
// Output: 2, nums = [1,2,_]
// Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
// It does not matter what you leave beyond the returned k (hence they are underscores).

// Constraints:
// 1 <= nums.length <= 3 * 104
// -100 <= nums[i] <= 100
// nums is sorted in non-decreasing order.

package easy

func removeDuplicates(nums []int) int {
	tmp := nums[0]
	k := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != tmp {
			nums[k] = nums[i]
			k++
		}
		tmp = nums[i]
	}
	return k
}
