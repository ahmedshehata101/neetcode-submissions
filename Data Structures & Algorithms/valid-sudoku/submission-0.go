func isValidSudoku(board [][]byte) bool {
    var rows, cols, boxes [9][9]bool

    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            val := board[r][c]
            if val == '.' {
                continue
            }
            d := val - '1' // 0-8 index for digit 1-9
            b := (r/3)*3 + c/3

            if rows[r][d] || cols[c][d] || boxes[b][d] {
                return false
            }
            rows[r][d] = true
            cols[c][d] = true
            boxes[b][d] = true
        }
    }
    return true
}