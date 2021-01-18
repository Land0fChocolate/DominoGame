package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type dominoGame struct {
	players   []player
	pieces    []dominoPiece
	grid      dominoGrid
	turnOrder []int
}

type player struct {
	playerNumber int
	ownedPieces  []dominoPiece
}

type dominoPiece struct {
	top int
	bot int
}

type dominoGrid struct {
	grid [][]string
}

func main() {
	fmt.Println("-------- Domino Game --------")

	//creating the game and players
	game := generateNewGame(pickPlayers())

	//start game
	fmt.Println("---- Game Start! ----")
	fmt.Println("game.turnOrder:", game.turnOrder)
	game.playGame()
	fmt.Println("---- Game End! ----")

	printDebug(game)
}

func printGrid(grid dominoGrid) {
	for _, v := range grid.grid {
		fmt.Println(v)
	}
}

func pickPlayers() int {
	var numPlayers int
	for {
		fmt.Printf("\nHow many players?\n")
		fmt.Scan(&numPlayers)
		if numPlayers >= 2 && numPlayers <= 4 {
			return numPlayers
		}
		fmt.Println("Invalid number of players. Please pick 2 to 4 players.")
	}
}

func generateNewGame(numPlayers int) dominoGame {
	var game dominoGame
	//build the player objects
	for i := 1; i < numPlayers+1; i++ {
		game.players = append(game.players, player{playerNumber: i})
	}
	//build the domino pieces
	k := 0
	for i := 0; i < 7; i++ {
		for j := k; j < 7; j++ {
			game.pieces = append(game.pieces, dominoPiece{top: i, bot: j})
		}
		k++
	}

	//build a 3x3 grid to start, this will expand as pieces get placed
	game.grid.grid = [][]string{{"X", "X", "X"}, {"X", "X", "X"}, {"X", "X", "X"}}

	//shuffling the pieces
	game, firstMove := assignPieces(game)
	game.turnOrder = generateTurnOrder(firstMove, game.players)

	return game
}

func assignPieces(gameRaw dominoGame) (dominoGame, int) {
	var firstTurn, highestDouble int
	//pieces will be reshuffled if nobody starts with a double
	for {
		game := gameRaw
		//assign domino pieces to players
		for k, player := range game.players {
			for i := 1; i <= 7; i++ {
				r := rand.Intn(len(game.pieces) - 1)
				player.ownedPieces = append(player.ownedPieces, game.pieces[r])
				firstTurn, highestDouble = firstMove(game.pieces[r], highestDouble, firstTurn, player.playerNumber)
				game.pieces = remove(game.pieces, r)
			}
			game.players[k] = player
		}
		if firstTurn != 0 {
			return game, firstTurn
		}
	}
}

//determining which player places the first piece
func firstMove(piece dominoPiece, highestDouble, firstTurn, playerNum int) (int, int) {
	if (piece.top == piece.bot) && (piece.top > highestDouble) {
		firstTurn = playerNum
		highestDouble = piece.top
	}
	return firstTurn, highestDouble
}

func remove(s []dominoPiece, i int) []dominoPiece {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func generateTurnOrder(firstMove int, players []player) (turnOrder []int) {
	turnOrder = append(turnOrder, firstMove)
	for _, player := range players {
		if player.playerNumber != firstMove {
			turnOrder = append(turnOrder, player.playerNumber)
		}
	}
	return
}

func (game *dominoGame) playGame() {
	firstTurn := true
	var pickedPiece int
	var newOwnedPieces []dominoPiece
	var newGrid dominoGrid
	for {
		//players place their pieces down in specific turns
		for _, playerNum := range game.turnOrder {
			printGrid(game.grid)
			if firstTurn {
				//have to place the highest doubles piece for the first turn
				highestDouble := getHighestDouble(game.players)
				fmt.Println("Player ", playerNum, " starts first with their highest double.")
				game.players[playerNum-1].ownedPieces, game.grid = placePiece(highestDouble, game.players[playerNum-1].ownedPieces, game.grid, true)
				if game.players[playerNum-1].ownedPieces == nil {
					fmt.Println("Error placing piece. This line of code should never be reached.")
					continue
				}
				firstTurn = false
			} else {
				//does the player have any viable pieces?
				viablePiece := false
				for _, piece := range game.players[playerNum-1].ownedPieces {
					if checkPiece(piece, game.grid) {
						viablePiece = true
						break
					}
				}
				//take from boneyard if no viable piece in player's hand
				if !viablePiece && len(game.pieces) > 0 {
					fmt.Println("No viable piece in player ", playerNum, "'s hand. Select piece from the boneyard.")
					for {
						for k2, piece := range game.pieces {
							fmt.Println(k2, " - ", piece)
						}
						fmt.Scan(&pickedPiece)
						if pickedPiece < 0 || pickedPiece > len(game.pieces)-1 {
							fmt.Println("Invalid selection. Pick a number from 0 to ", len(game.pieces)-1)
							continue
						}
						break
					}
				}
				if !viablePiece && len(game.pieces) == 0 {
					fmt.Println("Player ", playerNum, " cannot make a move this turn as they have no viable pieces and the boneyard is empty.")
					continue
				}
				for {
					fmt.Println("Player ", playerNum, " select a piece.")
					//print out pieces in a list, select from 1-numPieces
					for k2, piece := range game.players[playerNum-1].ownedPieces {
						fmt.Println(k2, " - ", piece)
					}
					fmt.Scan(&pickedPiece)
					if pickedPiece < 0 || pickedPiece > len(game.players[playerNum-1].ownedPieces)-1 {
						fmt.Println("Invalid selection. Pick a number from 0 to ", len(game.players[playerNum-1].ownedPieces)-1)
						continue
					}
					//proposed piece placement
					newOwnedPieces, newGrid = placePiece(game.players[playerNum-1].ownedPieces[pickedPiece], game.players[playerNum-1].ownedPieces, game.grid, false)
					if newOwnedPieces == nil {
						fmt.Println("Selected piece not valid. Pick a different piece.")
						continue
					}
					break
				}
				//place piece on the grid
				game.grid = newGrid
				game.players[playerNum-1].ownedPieces = newOwnedPieces
			}
			//check win conditions
			if len(game.players[playerNum-1].ownedPieces) == 0 && len(game.pieces) == 0 {
				fmt.Println("Player ", playerNum, " wins!")
				return
			}
		}
	}
}

func getHighestDouble(players []player) dominoPiece {
	var max dominoPiece
	for _, player := range players {
		for _, piece := range player.ownedPieces {
			if piece.top == piece.bot {
				if piece.top > max.top {
					max = piece
				}
			}
		}
	}
	return max
}

func placePiece(piece dominoPiece, playerPieces []dominoPiece, grid dominoGrid, firstTurn bool) ([]dominoPiece, dominoGrid) {
	var x, y, ori int
	var end2 string
	var newGrid dominoGrid
	//check viability of piece selected
	if !checkPiece(piece, grid) && !firstTurn {
		return nil, newGrid
	}
	//select which end of piece to place first
	end := selectPieceEnd(piece)

	//select square for end to go
	for {
		newGrid = grid
		printGrid(newGrid)
		//get x axis of grid
		for {
			fmt.Println("Type x-axis for end to go. (starting from 1 at top left)")
			fmt.Scan(&x)
			if x > len(newGrid.grid[0]) {
				fmt.Println("x too large. Grid is currently ", len(newGrid.grid[0]), " squares long.")
				continue
			}
			if x < 1 {
				fmt.Println("x too small. Start from 1.")
				continue
			}
			break
		}
		//get y axis of grid
		for {
			fmt.Println("Type y-axis for end to go. (starting from 1 at top left)")
			fmt.Scan(&y)
			if y > len(newGrid.grid) {
				fmt.Println("y too large. Grid is currently ", len(newGrid.grid), " squares long.")
				continue
			}
			if y < 1 {
				fmt.Println("y too small. Start from 1.")
				continue
			}
			break
		}
		//check if space already occupied
		if isSpaceAlreadyOccupied(newGrid, x, y) {
			continue
		}

		//check if space is next to equivalent end
		if !isSpaceNextToEquivalentEnd(newGrid, y, x, end) && !firstTurn {
			continue
		}

		//place end
		newGrid.grid[y-1][x-1] = end
		//if end coordinates are on the edge of the grid, expand grid
		fmt.Println("=== y: ", y, "x: ", x)
		if y == 1 {
			expandGrid("top", newGrid)
		}
		if x == len(newGrid.grid) {
			expandGrid("right", newGrid)
		}
		if y == len(newGrid.grid[0]) {
			expandGrid("bot", newGrid)
		}
		if x == 1 {
			expandGrid("left", newGrid)
		}
		//get the other end of the domino piece
		endInt, _ := strconv.Atoi(end)
		if piece.top == endInt {
			end2 = strconv.Itoa(piece.bot)
		} else {
			end2 = strconv.Itoa(piece.top)
		}
		//get orientation, expand grid if end2 touches the edge of grid
		for {
			printGrid(newGrid)
			fmt.Println("Select orientation. 1-up, 2-right, 3-down, 4-left.")
			fmt.Scan(&ori)
			switch ori {
			case 1:
				if newGrid.grid[y-2][x-1] != "X" {
					fmt.Println("Space already occupied. Select an empty space.")
					continue
				}
				//place end2
				newGrid.grid[y-2][x-1] = end2
				if y == 2 {
					expandGrid("top", newGrid)
				}
				break
			case 2:
				if newGrid.grid[y-1][x] != "X" {
					fmt.Println("Space already occupied. Select an empty space.")
					continue
				}
				//place end2
				newGrid.grid[y-1][x] = end2
				if x == len(newGrid.grid)-1 {
					expandGrid("right", newGrid)
				}
				break
			case 3:
				if newGrid.grid[y][x-1] != "X" {
					fmt.Println("Space already occupied. Select an empty space.")
					continue
				}
				//place end2
				newGrid.grid[y][x-1] = end2
				if y == len(newGrid.grid[0])-1 {
					expandGrid("bot", newGrid)
				}
				break
			case 4:
				if newGrid.grid[y-1][x-2] != "X" {
					fmt.Println("Space already occupied. Select an empty space.")
					continue
				}
				//place end2
				newGrid.grid[y-1][x-2] = end2
				if x == 2 {
					expandGrid("left", newGrid)
				}
				break
			default:
				fmt.Println("Invalid orientation. Select one of the numbers for each side.")
			}
			break
		}
		break
	}
	//overwrite with the new grid
	grid = newGrid
	printGrid(grid)

	//remove piece from owned player pieces
	for k, playerPiece := range playerPieces {
		if playerPiece == piece {
			playerPieces = remove(playerPieces, k)
			break
		}
	}
	return playerPieces, grid
}

func checkPiece(piece dominoPiece, grid dominoGrid) bool {
	viable := false
	for y := 1; y <= len(grid.grid)-2; y++ {
		for x := 1; x <= len(grid.grid[0])-2; x++ {
			//check if it could be matched with any domino on the board
			if grid.grid[y][x] == strconv.Itoa(piece.top) || grid.grid[y][x] == strconv.Itoa(piece.bot) {
				//check if there is room to place
				if grid.grid[y+1][x] == "X" || grid.grid[y-1][x] == "X" || grid.grid[y][x+1] == "X" || grid.grid[y][x-1] == "X" {
					viable = true
				}
			}
		}
	}
	return viable
}

func selectPieceEnd(piece dominoPiece) (end string) {
	for {
		fmt.Println("Piece ", piece, " selected. Select end: Top -", piece.top, " Bot -", piece.bot)
		fmt.Scan(&end)
		endInt, err := strconv.Atoi(end)
		if err != nil {
			fmt.Println("Invalid input. Type in a number.")
			continue
		}
		if endInt != piece.top && endInt != piece.bot {
			fmt.Println("Invalid end. Select ", piece.top, " or ", piece.bot, ".")
			continue
		}
		break
	}
	return
}

func isSpaceAlreadyOccupied(newGrid dominoGrid, x, y int) (occupied bool) {
	if newGrid.grid[y-1][x-1] != "X" {
		fmt.Println("There is already a piece here. Choose a free set of coordinates.")
		printGrid(newGrid)
		return true
	}
	return false
}

func isSpaceNextToEquivalentEnd(newGrid dominoGrid, y, x int, end string) (spaceViable bool) {
	fmt.Println("x: ", x, " y: ", y, " end: ", end)
	//check space above if possible
	if y != 1 {
		if newGrid.grid[y-2][x-1] == end {
			return true
		}
	}
	//check space to the right if possible
	if x != len(newGrid.grid[0]) {
		if newGrid.grid[y-1][x] == end {
			return true
		}
	}
	//check space below if possible
	if y != len(newGrid.grid) {
		if newGrid.grid[y][x-1] == end {
			return true
		}
	}
	//check space to the left if possible
	if x != 1 {
		if newGrid.grid[y-1][x-2] == end {
			return true
		}
	}

	fmt.Println("There is no equivalent end next to this space. Select different coordinates.")
	return false
}

func expandGrid(edge string, grid dominoGrid) dominoGrid {
	fmt.Println("== expand grid ", edge)
	printGrid(grid)
	switch edge {
	case "top":
		//add a row and shift everything down by 1.
		grid.grid = append([][]string{grid.grid[len(grid.grid)-1]}, grid.grid...)
	case "right":
		//add a column
		for k, row := range grid.grid {
			row = append(row, "X")
			grid.grid[k] = row
		}
	case "bot":
		//add a row
		grid.grid = append(grid.grid, grid.grid[0])
	case "left":
		//add a column and shift everything right by 1
		for k, _ := range grid.grid {
			grid.grid[k] = append([]string{"X"}, grid.grid[k]...)
		}
	}
	fmt.Println()
	printGrid(grid)
	return grid
}

func printDebug(game dominoGame) {
	fmt.Println("--- game debug ---")
	//printing data for DEBUG
	fmt.Println("players: ", game.players)
	fmt.Println("pieces: ", game.pieces)
	fmt.Println("--- ---------- ---")
	//printGrid(game.grid)
}
