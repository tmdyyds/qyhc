<?php
//8. 字符串转换整数 (atoi)
function myAtoi($str) {
    $str = trim($str);
    $max = pow(2, 31) - 1;
    $min = pow(-2, 31);

    $prefix = '';
    if ($str[0] == '-') {
        $prefix = '-';
        $str = substr($str, 1);
    }

    $ret = '';
    for ($i=0; $i < strlen($str); $i++) {
        if ((ord($str[$i]) >= 97 && ord($str[$i]) <= 122) || (ord($str[$i]) >= 65 && ord($str[$i]) <= 90) || ord($str[$i]) == 45) {
            break;
        }

        $ret .= $str[$i];
    }

    if ($prefix && $ret > 0) {
        $num = $prefix.$ret < $min ? $min : $prefix.$ret;
    } else {
        $num = $ret > $max ? $max : $ret;
    }

    return (int)$num;
    }
}