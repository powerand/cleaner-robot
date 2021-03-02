package main
import "fmt"

const X = 5

func main() {
	vRoom := [][]int{
		{X, X, X, X, X, X, X, X},
		{X, 0, 0, 0, 0, X, X, X},
		{X, 0, 0, 0, X, 0, X, X},
		{X, 0, 0, X, X, 0, X, X},
		{X, 0, 0, 0, 0, 0, 0, X},
		{X, 0, 0, X, X, 0, X, X},
		{X, 0, 0, 0, X, 0, X, X},
		{X, X, X, X, X, X, X, X},
	}

	row := 3
	col := 2

	// default direction - top
	dirX := 0
	dirY := -1

	success := true
	borderReached := false
	fails := 0
	impudence := 1

	PrintRoom(vRoom)

	for {
		success = Move(&col, &row, dirX, dirY, vRoom, impudence)
		if !borderReached {
			if success {
				continue
			} else {
				borderReached = true
				TurnRight(&dirX, &dirY)
				success = Move(&col, &row, dirX, dirY, vRoom, impudence)
			}
		}
		if success {
			fails = 0
			impudence = 1
			TurnLeft(&dirX, &dirY)
		} else {
			fails++
			TurnRight(&dirX, &dirY)
		}
		if fails > 4 {
			if impudence == X {
				vRoom[row][col]++
				PrintRoom(vRoom)
				break
			}
			impudence++
			fails = 0
		}
	}

	PrintRoom(vRoom)
}

func TurnLeft(dirX *int, dirY *int) {
	if *dirX == 0 {
		*dirX = *dirY
		*dirY = 0
		return
	}
	if *dirY == 0 {
		*dirY = *dirX * -1
		*dirX = 0
		return
	}
}

func TurnRight(dirX *int, dirY *int) {
	if *dirX == 0 {
		*dirX = *dirY * -1
		*dirY = 0
		return
	}
	if *dirY == 0 {
		*dirY = *dirX
		*dirX = 0
		return
	}
}

func Move(col *int, row *int, dirX int, dirY int, vRoom [][]int, impudence int) (success bool) {
	nCol := *col + dirX
	nRow := *row + dirY
	success = vRoom[nRow][nCol] < impudence
	if success {
		Clean(*col, *row, vRoom)
		vRoom[*row][*col]++
		*row = nRow
		*col = nCol
	}
	PrintRoom(vRoom)
	return success
}

func Clean(col int, row int, vRoom [][]int) {
}

func PrintRoom(room [][]int) {
	for _, rRow := range room {
		fmt.Printf("%v\n", rRow)
	}
	fmt.Printf("--------\n")
}
