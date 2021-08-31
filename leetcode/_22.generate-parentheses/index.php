<?php

class Solution
{
    /**
     * @param Integer $n
     * @return String[]
     */
    function generateParenthesis($n)
    {
        $out = [];
        $this->generate($out, 0, 0, $n, "");
        return $out;
    }

    private function generate(&$out, $l, $r, $n, $pre)
    {
        if ($l >= $n && $r >= $n) {
            $out[] = $pre;
            return;
        }

        if ($l <= $r) {
            $this->generate($out, $l + 1, $r, $n, $pre . "(");
        } else if ($l < $n) {
            $this->generate($out, $l + 1, $r, $n, $pre . "(");
            $this->generate($out, $l, $r + 1, $n, $pre . ")");
        } else {
            $this->generate($out, $l, $r + 1, $n, $pre . ")");
        }
    }
}

$arr = (new Solution())->generateParenthesis(3);
var_dump($arr);
