import java.util.ArrayList;
import java.util.List;

class Solution {
    public List<String> generateParenthesis(int n) {
        List<String> arr = new ArrayList<String>();
        generate(arr, 0, 0, n, new StringBuilder());
        return arr;
    }

    public void generate(List<String> arr, int l, int r, int n, StringBuilder pre) {
        if (l >= n && r >= n) {
            arr.add(pre.toString());
            return;
        }
        if (l <= r) {
            pre.append("(");
            generate(arr, l + 1, r, n, pre);
            pre.deleteCharAt(l + r);
        } else if (l < n) {
            pre.append("(");
            generate(arr, l + 1, r, n, pre);
            pre.deleteCharAt(l + r);

            pre.append(")");
            generate(arr, l, r + 1, n, pre);
            pre.deleteCharAt(l + r);
        } else {
            pre.append(")");
            generate(arr, l, r + 1, n, pre);
            pre.deleteCharAt(l + r);
        }
    }

    public static void main(String[] args) {
        System.out.println(new Solution().generateParenthesis(3));
    }
}