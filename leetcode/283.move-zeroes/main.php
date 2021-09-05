<?php

class Solution
{
    function moveZeroes(&$nums)
    {
        $p = 0;
        for ($i = 0; $i < count($nums); $i++) {
            if ($nums[$i] != 0) {
                if ($p != $i) {
                    $tmp = $nums[$i];
                    $nums[$i] = $nums[$p];
                    $nums[$p] = $tmp;
                }
                $p++;
            }
        }
    }
}

$nums = [0, 1, 0, 2, 0, 0, -1];
(new Solution())->moveZeroes($nums);
print_r($nums);
