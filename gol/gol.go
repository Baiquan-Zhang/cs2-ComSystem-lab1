package main

func calculateNextState(p golParams, world [][]byte) [][]byte {
	var changeCells []cell
	var temp cell

	// var newWorld [][]byte

	// for i, column := range world {
	// 	var newColumn []byte
	// 	newWorld = append(newWorld, newColumn)
	// }

	for i, column := range world {
		for j, pixel := range column {
			if findNextState(pixel, i, j, p.imageHeight, p.imageWidth, world) {
				temp.x = j
				temp.y = i
				changeCells = append(changeCells, temp)
			}
		}
	}

	for _, diff := range changeCells {
		changeCell(diff, world)
	}

	return world
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	aliveCells := []cell{}
	var newCell cell
	for i := 0; i < p.imageHeight; i++ {
		for j := 0; j < p.imageWidth; j++ {
			if world[i][j] == 255 {
				newCell.x = j
				newCell.y = i
				aliveCells = append(aliveCells, newCell)
			}
		}
	}
	return aliveCells
}

func findNextState(pixel byte, i, j, height, width int, world [][]byte) bool {
	num := findLiveNeigh(i, j, height, width, world)
	if pixel == 255 {
		if num < 2 || num > 3 {
			return true
		}
	} else {
		if num == 3 {
			return true
		}
	}
	return false
}

func findLiveNeigh(i, j, height, width int, world [][]byte) int {

	var preRow, preCol, nextRow, nextCol int

	if i == 0 {
		preRow = height - 1
		nextRow = i + 1
	} else if i == height-1 {
		preRow = i - 1
		nextRow = 0
	} else {
		preRow = i - 1
		nextRow = i + 1
	}

	if j == 0 {
		preCol = width - 1
		nextCol = j + 1
	} else if j == width-1 {
		preCol = j - 1
		nextCol = 0
	} else {
		preCol = j - 1
		nextCol = j + 1
	}

	liveNeigh := 0
	Neigh := []byte{
		world[preRow][preCol], world[preRow][j], world[preRow][nextCol],
		world[i][preCol], world[i][nextCol],
		world[nextRow][preCol], world[nextRow][j], world[nextRow][nextCol]}
	for _, pixel := range Neigh {
		if pixel == 255 {
			liveNeigh++
		}
	}

	return liveNeigh
}

func changeCell(diff cell, world [][]byte) {
	if world[diff.y][diff.x] == 255 {
		world[diff.y][diff.x] = 0
	} else {
		world[diff.y][diff.x] = 255
	}
}
