<?php
//4. 寻找两个正序数组的中位数
class Solution {

    /**
     * @param Integer[] $nums1
     * @param Integer[] $nums2
     * @return Float
     */
    function findMedianSortedArrays($nums1, $nums2) {
            $mergeArr = array_merge($nums1, $nums2);
    sort($mergeArr);

    $center = intval(count($mergeArr)/2);
    //偶数
    if (count($mergeArr)%2 == 0) {
        $ret = ($mergeArr[$center-1] + $mergeArr[$center]) / 2;
    } else {
        $ret = $mergeArr[$center];
    }

    return $ret;
    }
}