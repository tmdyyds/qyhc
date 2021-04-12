//179. 最大数
//golang
func largestNumber(nums []int) string {
	ss := make([]string, len(nums))
	for i, num := range nums {
		ss[i] = strconv.Itoa(num)
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i]+ss[j] >= ss[j]+ss[i]
	})
	o := strings.Join(ss, "")
	if o[0] == '0' {
		return "0"
	}

	return o
}

//php
class Solution {

    /**
     * @param Integer[] $nums
     * @return String
     */
    function largestNumber($nums) {
        $len_n = count($nums);
        usort($nums, function($a, $b) {
            if (strcmp("".$a. $b, "". $b. $a) >= 0) {
                return -1;
            } else {
                return 1;
            }
        });

        return join("", $nums) == "0" ? "0" : join("", $nums);
    }
}


$n = new Solution;
$nums = [3,30,34,5,9];
var_dump($n->largestNumber($nums));