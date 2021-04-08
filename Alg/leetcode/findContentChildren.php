<?php
// 455. 分发饼干
function findContentChildren($g, $s) {
    sort($g);
    sort($s);

    $cnt = 0;
    $j = 0;
    for ($i=0; $i < count($s); $i++) {
        if ($j < count($g) && $s[$i] >= $g[$j]) {
            $cnt++;
            $j++;
        }
    }

    return $cnt;
}


$g = [1,2]; $s = [1,2,3];

var_dump(findContentChildren($g, $s));