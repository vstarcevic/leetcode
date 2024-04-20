package hard

// Median of Two Sorted Arrays
// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
// The overall run time complexity should be O(log (m+n)).

// Example 1:
// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.

// Example 2:
// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalEl := len(nums1) + len(nums2)
	medianCounter := 0
	needTwo := totalEl%2 == 0

	var result float64

	counter1 := 0
	counter2 := 0
	stopAt := 0
	if needTwo {
		stopAt = totalEl/2 - 1
	} else {
		stopAt = totalEl / 2
	}

	for i := 0; i < totalEl; i++ {
		if len(nums1)-1 >= counter1 && len(nums2)-1 >= counter2 {

			if nums1[counter1] <= nums2[counter2] {
				if i >= stopAt {
					result += float64(nums1[counter1])
					medianCounter++
				}
				counter1++
			} else {
				if i >= stopAt {
					result += float64(nums2[counter2])
					medianCounter++
				}
				counter1++
			}

		} else if len(nums1)-1 >= counter1 {
			if i >= stopAt {
				result += float64(nums1[counter1])
				medianCounter++
			}
			counter1++
		} else {
			if i >= stopAt {
				result += float64(nums2[counter2])
				medianCounter++
			}
			counter2++
		}

		if medianCounter == 2 && needTwo {
			return result / 2
		}

		if medianCounter == 1 && !needTwo {
			return result
		}
	}

	return 0.0
}
