/*here are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
Subscribe to see which companies asked this question*/

func minInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func getMedian(nums1 []int, nums2 []int, key int) int {
	if len(nums1) > len(nums2) {
		return getMedian(nums2, nums1, key)
	}

	if len(nums1) == 0 {
		return nums2[key-1]
	}

	if key == 1 {
		return minInt(nums1[0], nums2[0])
	}

	posA := minInt(key/2, len(nums1))
	posB := key - posA
	if nums1[posA-1] < nums2[posB-1] {
		return getMedian(nums1[posA:], nums2, key-posA)
	} else if nums1[posA-1] > nums2[posB-1] {
		return getMedian(nums1, nums2[posB:], key-posB)
	}
	return nums1[posA-1]

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total%2 != 0 {
		return float64(getMedian(nums1, nums2, total/2+1))
	} else {
		return (float64(getMedian(nums1, nums2, total/2+1)) + float64(getMedian(nums1, nums2, total/2))) / 2
	}
}