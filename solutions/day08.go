package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type Day08 struct{}

func (Day08) Name() string {
	return "day08"
}

type forest struct {
	height      int
	width       int
	treeHeights [][]int
}

func forestFromInput(input string) *forest {
	lines := strings.Split(input, "\n")
	f := forest{height: len(lines), width: len(lines[0])}
	for _, line := range lines {
		var trees []int
		for _, tree := range line {
			height, _ := strconv.Atoi(string(tree))
			trees = append(trees, height)
		}
		f.treeHeights = append(f.treeHeights, trees)
	}
	return &f
}

func (f *forest) String() string {
	var sb strings.Builder
	for _, line := range f.treeHeights {
		for _, treeHeight := range line {
			sb.WriteString(fmt.Sprint(treeHeight))
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (f *forest) treeScore(row, col int) (int, bool) {
	treeHeight := f.treeHeights[row][col]
	visible := false
	score := 1
	distance := 1
	for x := col + 1; x < f.width && treeHeight > f.treeHeights[row][x]; x++ {
		if x == f.width-1 {
			visible = true
		} else {
			distance++
		}
	}
	score *= distance
	distance = 1
	for x := col - 1; x >= 0 && treeHeight > f.treeHeights[row][x]; x-- {
		if x == 0 {
			visible = true
		} else {
			distance++
		}
	}
	score *= distance
	distance = 1
	for y := row + 1; y < f.height && treeHeight > f.treeHeights[y][col]; y++ {
		if y == f.height-1 {
			visible = true
		} else {
			distance++
		}
	}
	score *= distance
	distance = 1
	for y := row - 1; y >= 0 && treeHeight > f.treeHeights[y][col]; y-- {
		if y == 0 {
			visible = true
		} else {
			distance++
		}
	}
	score *= distance
	return score, visible
}

func (Day08) Solution(input string, part2 bool) string {
	f := forestFromInput(input)
	if !part2 {
		total := 2 * (f.height + f.width - 2)
		for row := 1; row < f.height-1; row++ {
			for col := 1; col < f.width-1; col++ {
				if _, visible := f.treeScore(row, col); visible {
					total++
				}
			}
		}
		return fmt.Sprint(total)
	} else {
		max := 0
		for row := 1; row < f.height-1; row++ {
			for col := 1; col < f.width-1; col++ {
				if score, _ := f.treeScore(row, col); score > max {
					max = score
				}
			}
		}
		return fmt.Sprint(max)
	}
}
