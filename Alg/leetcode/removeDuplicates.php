<?php
//26. 删除有序数组中的重复项
function removeDuplicates(&$nums) {
    $cnt = count($nums);
    if ($cnt <= 1) {
        return $cnt;
    }

    //双指针
    $i = 0; //慢指针
    for ($j=1; $j < $cnt; $j++) {
        if ($nums[$i] != $nums[$j]) {
            $i++; //慢前期
            $nums[$i] = $nums[$j];
        }
    }

    return $i + 1;
}