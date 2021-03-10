package main

func CheckForWinner(board [9]string, n int) (bool, string) {
	boolean := false
	var winner string

	boolean, winner = CheckHorizontal(board)
	if boolean == true {
		return boolean, winner
	}
	boolean, winner = CheckVertical(board)
	if boolean == true {
		return boolean, winner
	}
	boolean, winner = CheckLeftDiagonal(board)
	if boolean == true {
		return boolean, winner
	}
	boolean, winner = CheckRightDiagonal(board)
	if boolean == true {
		return boolean, winner
	} else {
		return CheckDraw(n)
	}
}

func CheckHorizontal(board [9]string) (bool, string) {
	i := 0
	test := false

	for i < 9 {
		test = board[i] == board[i+1] &&
			board[i+1] == board[i+2] &&
			board[i] != ""
		if !test {
			i += 3
		} else {
			return true, board[i]
		}
	}
	return false, ""
}

func CheckVertical(board [9]string) (bool, string) {
	test := false
	i := 0

	for i < 3 {
		test = board[i] == board[i+3] &&
			board[i+3] == board[i+6] &&
			board[i] != ""
		if !test {
			i += 1
		} else {
			return true, board[i]
		}
	}
	return false, ""
}

func CheckLeftDiagonal(board [9]string) (bool, string) {
	if board[0] == board[4] && board[4] == board[8] && board[0] != "" {
		return true, board[0]
	}
	return false, ""
}

func CheckRightDiagonal(board [9]string) (bool, string) {
	if board[2] == board[4] && board[4] == board[6] && board[2] != "" {
		return true, board[0]
	}
	return false, ""
}

func CheckDraw(turnNumber int) (bool, string) {
	if turnNumber == 9 {
		return true, ""
	}
	return false, ""
}
