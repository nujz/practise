pub struct Solution;

impl Solution {
    pub fn exist(board: Vec<Vec<char>>, word: String) -> bool {
        let mut points = Vec::new();
        points.resize(board.len(), vec![]);
        for i in 0..board.len() {
            points[i].resize(board[i].len(), false);
        }
        let word = word.chars().collect::<Vec<char>>();

        for i in 0..board.len() {
            for j in 0..board[i].len() {
                if Self::search(&board, &word, i, j, 0, &mut points) {
                    return true;
                }
            }
        }

        false
    }

    fn search(
        board: &Vec<Vec<char>>,
        word: &Vec<char>,
        i: usize,
        j: usize,
        k: usize,
        points: &mut Vec<Vec<bool>>,
    ) -> bool {
        if points[i][j] {
            return false;
        }
        if board[i][j] != word[k] {
            return false;
        }
        if k + 1 == word.len() {
            return true;
        }

        points[i][j] = true;
        if i >= 1 && Self::search(board, word, i - 1, j, k + 1, points) {
            return true;
        }
        if i + 1 < board.len() && Self::search(board, word, i + 1, j, k + 1, points) {
            return true;
        }

        if j >= 1 && Self::search(board, word, i, j - 1, k + 1, points) {
            return true;
        }
        if j + 1 < board[i].len() && Self::search(board, word, i, j + 1, k + 1, points) {
            return true;
        }
        points[i][j] = false;
        return false;
    }
}

fn main() {
    let arr = vec![
        vec!['A', 'B', 'C', 'E'],
        vec!['S', 'F', 'C', 'S'],
        vec!['A', 'D', 'E', 'E'],
    ];
    let b = Solution::exist(arr, "ABCCED".to_owned());
    println!("{:?}", b);
}
