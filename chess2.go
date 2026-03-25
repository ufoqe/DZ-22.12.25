package main

import "fmt"

func main() {

	var c int

	fmt.Print("Введите размерность шахматной доски: ")
	fmt.Scan(&c)

	if c <= 0 {
		fmt.Println("Некорректный размер")
		return
	}
	var d, e string

	fmt.Println("Введите имя игрока за белых: ")
	fmt.Scan(&d)
	fmt.Println("Введите имя игрока за чёрных: ")
	fmt.Scan(&e)

	whitePieces := []string{"♜", "♞", "♝", "♛", "♚", "♝", "♞", "♜"}
	whitePawn := "♟"
	blackPieces := []string{"♖", "♘", "♗", "♕", "♔", "♗", "♘", "♖"}
	blackPawn := "♙"
	light := "⬜"
	dark := "⬛"

	stolb := make([]string, c)
	for j := 0; j < c; j++ {
		if j < 26 {
			stolb[j] = string(rune('A' + j))
		} else {
			stolb[j] = fmt.Sprintf("%c%c", 'a'+j/26-1, 'a'+j%26)
		}
	}

	fmt.Printf("\nЧёрные: %s\n", e)

	fmt.Print("   ")
	for j := 0; j < c; j++ {
		fmt.Printf(" %s", stolb[j])
	}
	fmt.Println()

	for i := 0; i < c; i++ {
		rowNum := c - i
		fmt.Printf("%2d ", rowNum)

		for j := 0; j < c; j++ {
			var cell string
			if (i+j)%2 == 0 {
				cell = light
			} else {
				cell = dark
			}
			var pp string
			switch {
			case i == c-1:
				pp = whitePieces[j%len(whitePieces)]
			case i == c-2:
				pp = whitePawn
			case i == 1:
				pp = blackPawn
			case i == 0:
				pp = blackPieces[j%len(blackPieces)]
			default:
				pp = " "
			}

			if pp != " " {
				fmt.Print(pp + " ")
			} else {
				fmt.Print(cell)
			}
		}
		fmt.Printf(" %d", rowNum)
		fmt.Println()
	}
	fmt.Print("   ")
	for j := 0; j < c; j++ {
		fmt.Printf(" %s", stolb[j])
	}
	fmt.Println()
	fmt.Printf("Белые : %s\n", d)

}
