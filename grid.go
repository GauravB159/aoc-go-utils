package aocutils

import "strconv"

type Grid struct {
	Data map[Key]int
	Rows int
	Cols int
}

type Key struct {
	Row int
	Col int
}

func CreateGrid(lines []string) Grid {
	grid := Grid{
		Data: make(map[Key]int, len(lines)*len(lines[0])),
		Rows: len(lines),
		Cols: len(lines[0]),
	}
	for row, line := range lines {
		for col, char := range line {
			value, _ := strconv.Atoi(string(char))
			grid.Data[Key{Row: row, Col: col}] = value + 1
		}
	}
	return grid
}

func (grid *Grid) GetNumRows() int {
	return grid.Rows
}

func (grid *Grid) GetNumCols() int {
	return grid.Cols
}

func (grid *Grid) GetData() map[Key]int {
	return grid.Data
}

func (grid *Grid) GetValue(row int, col int) (int, bool) {
	key := Key{Row: row, Col: col}
	value, exists := grid.Data[key]
	return value, exists
}

func (grid *Grid) SetValue(row int, col int, value int) {
	key := Key{Row: row, Col: col}
	grid.Data[key] = value
}
