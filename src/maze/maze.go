package main

import (
	"fmt"
	"os"
)

//从文件中读取迷宫数组
func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

//坐标，i是行，j是列
type point struct {
	i, j int
}

//方向，上左下右
var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

//检查某个点是否越界，并返回该点的值
func (p point) at(grid [][]int) (int, bool) {
	//检查行是否越界
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	//检查列是否越界
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

//广度优先搜索迷宫
func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	//作为队列使用，存放走得通的点,当某个点要被探索便把它出队列
	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}
		//沿着上左下右四个方向走一遍
		for _, dir := range dirs {
			next := cur.add(dir)
			val, ok := next.at(maze)
			//判断该点是否越界或者遇到墙
			if !ok || val == 1 {
				continue
			}
			//判断该点是否越界或者已经走过了
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			//不能往回走
			if next == start {
				continue
			}
			//当前的步骤数
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			//next作为下一个要被探索的点放进队列
			Q = append(Q, next)

		}
	}
	return steps
}

//查看到达某个点所走的步数
func stepNum(steps [][]int, end point) int {
	return steps[end.i][end.j]
}

func main() {
	maze := readMaze("src/maze/maze.in")
	//把左上角和右下角的点分别作为起点和终点
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
	//到达终点走的步数
	num := stepNum(steps, point{len(maze) - 1, len(maze[0]) - 1})
	fmt.Printf("arrival terminal point: %d", num)
}
