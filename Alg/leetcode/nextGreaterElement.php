<?php
//496. 下一个更大元素 I
class Solution {

    /**
     * @param Integer[] $nums1
     * @param Integer[] $nums2
     * @return Integer[]
     */
    function nextGreaterElement($nums1, $nums2) {
            $stack = new SplStack;

            $tmp = [];
            for ($i=0; $i < count($nums2); $i++) {
                while (!$stack->isEmpty() && $nums2[$i] > $stack->top()) {
                    //若存在第一个大于的值 则保存为临时数组
                    $tmp[$stack->pop()] = $nums2[$i];
                }

                //入栈
                $stack->push($nums2[$i]);
            }

            $ret = [];
            for ($i=0; $i < count($nums1); $i++) {
                if (isset($tmp[$nums1[$i]])) {
                    $ret[] = $tmp[$nums1[$i]];
                } else {
                    $ret[] = -1;
                }
            }


            return $ret;
    }
}