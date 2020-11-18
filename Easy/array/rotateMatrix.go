package main

import "fmt"

// 给定一个n×n的二维矩阵表示一个图像。将图像顺时针旋转 90 度。
//说明:你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
// [
//  [1,2,3],
//  [4,5,6],
//  [7,8,9]
//],
//[
//  [7,4,1],
//  [8,5,2],
//  [9,6,3]
//]
// [
//  [ 5, 1, 9,11],
//  [ 2, 4, 8,10],
//  [13, 3, 6, 7],
//  [15,14,12,16]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//  [15,13, 2, 5],
//  [14, 3, 4, 1],
//  [12, 6, 8, 9],
//  [16, 7,10,11]
//]

// 首先对数组转置，然后每行翻转
func rotateMatrix1(matrix [][]int)  {
	// 转置
	for i := 0; i < len(matrix[0]); i++ {
		for j := i+1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 每行翻转
	for i := 0; i < len(matrix); i++ {
		reverseArr(matrix[i])
	}
}
// 翻转数组
func reverseArr(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] =  arr[len(arr)-i-1], arr[i]
	}
}



func main() {
	matrix := [][]int{{1,2,3}, {4,5,6}, {7,8,9}}
	rotateMatrix1(matrix)
	fmt.Println(matrix)
}
