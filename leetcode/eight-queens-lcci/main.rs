pub struct Solution;

impl Solution {
    fn solve(out: &mut Vec<Vec<String>>, queens: &mut Vec<Vec<u8>>, row: usize) {
        if row == queens.len() {
            out.push(
                queens
                    .iter_mut()
                    .map(|a| String::from_utf8_lossy(a).into())
                    .collect(),
            );
            return;
        }

        for col in 0..queens[row].len() {
            if Self::valid(queens, row, col) {
                queens[row][col] = b'Q';
                Self::solve(out, queens, row + 1);
                queens[row][col] = b'.';
            }
        }
    }

    fn valid(queens: &Vec<Vec<u8>>, row: usize, col: usize) -> bool {
        for i in 0..row {
            if queens[i][col] == b'Q' {
                return false;
            }
        }

        let clen = queens[row].len();

        let mut r = row + 1;
        let mut c = col;
        while r > 0 && c < clen {
            if queens[r - 1][c] == b'Q' {
                return false;
            }
            r -= 1;
            c += 1;
        }

        let mut r = row + 1;
        let mut c = col + 1;
        while r > 0 && c > 0 {
            if queens[r - 1][c - 1] == b'Q' {
                return false;
            }
            r -= 1;
            c -= 1;
        }

        true
    }

    pub fn solve_n_queens(n: i32) -> Vec<Vec<String>> {
        let mut queens = vec![vec![b'.'; n as _]; n as _];
        let mut out = vec![];
        Self::solve(&mut out, &mut queens, 0);
        out
    }
}

fn main() {
    let a = Solution::solve_n_queens(4);
    println!("{:?}", a);
}
