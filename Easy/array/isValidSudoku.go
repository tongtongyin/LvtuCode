package main

// 有效的数独
// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。

//  遍历一次即可 判断每行是否重复，每列是否重复，没个3x3方格是否重复
// 判断每行是否重复：创建一个9*9的数组row，row[i][num]表示当前数字num在第i行num的位置，如果不在给这个位置填上1，如果存在那么返回false
// 判断列是同样的道理
// 判断每个小方盒：一共有0-8  9个小方盒每个小方格索引的计算方法为 (i/3)*3 + j/3
func isValidSudoku(board [][]byte) bool {
	var row, col, box [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				if row[i][num] == 1 {
					return false
				} else {
					row[i][num] = 1
				}
				if col[j][num] == 1 {
					return false
				} else {
					col[j][num] = 1
				}
				index := (i/3)*3 + j/3
				if box[index][num] == 1 {
					return false
				} else {
					box[index][num] = 1
				}
			}

		}
	}
	return true
}

func main() {

}
