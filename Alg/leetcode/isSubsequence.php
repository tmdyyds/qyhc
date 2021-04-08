<?php
// 392. 判断子序列
/*function isSubsequence($s, $t) {
    $n = strlen($s);
    $m = strlen($t);

    $i = 0;
    $j = 0;
    while ($i < $n && $j < $m) {
        if ($s[$i] == $t[$j]) {
            $i++;
        }

        $j++;
    }

    return $i == $n; //精髓
}
*/

function isSubsequence($s, $t) {
    if ($s == $t) {
        return true;
    }

    $sn = strlen($s);
    $j = -1;
    for ($i=0; $i < $sn; $i++) {
        $j = strpos($t, $s[$i], $j + 1); //关键点 每次查询$j都会变
        if ($j === false) {
            return false;
        }
    }

    return true;
}

