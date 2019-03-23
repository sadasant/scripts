let { expect } = require('chai')

let board_vertial_win = [
  [0, 0, 0, 1],
  [0, 0, 0, 0],
  [0, 0, 0, 1],
  [0, 0, 0, 1]
]

// moves:
// [ [player1 0 0] [player2 1 1] [player1 0 1] [player2 2 2] [player1 0 3] [player2 2 1] [player1 0 1] ]
let board_horizontal_win = [
  [1, 0, 1, 1],
  [0, 2, 0, 0],
  [0, 0, 2, 0],
  [0, 2, 0, 0]
]

let board_diagonal_win = [
  [2, 0, 0, 1],
  [0, 2, 1, 0],
  [0, 0, 2, 0],
  [1, 2, 0, 0]
]

// O(n)
let checkArrayWin = (array, value) => array.filter(x => x === value).length === array.length - 1 

// CORE STUFF
// TODO (because OCD): Move to a different file

// O(n)
let checkHorizontalWin = (board, move, player) => {
  let rowNumber = move[0]
  let row = board[rowNumber]
  return checkArrayWin(row, player)
}
expect(checkHorizontalWin(board_horizontal_win, [0, 1], 1)).to.equal(true)

// O(n)
let checkVerticalWin = (board, move, player) => {
  let colNumber = move[1]
  let col = board.reduce((a, row) => [...a, row[colNumber]], [])
  return checkArrayWin(col, player)
}
expect(checkVerticalWin(board_vertial_win, [0, 3], 1)).to.equal(true)

// O(n)
let checkDiagonal1Win = (board, move, player) => {
  let diagonal = board.reduce((a, row) => [...a, row[board.length - (a.length + 1)]], [])
  return checkArrayWin(diagonal, player)
}
expect(checkDiagonal1Win(board_diagonal_win, [2, 1], 1)).to.equal(true)

// O(n)
let checkDiagonal2Win = (board, move, player) => {
  let diagonal = board.reduce((a, row) => [...a, row[a.length]], [])
  return checkArrayWin(diagonal, player)
}
expect(checkDiagonal2Win(board_diagonal_win, [3, 3], 2)).to.equal(true)

// O(n)
let checkWinningMove = (...params) => {
  return (
    // TODO: CHECK IF THE MOVE IS VALID
    // TODO: pick some of them based on corners
    checkHorizontalWin(...params) ||
    checkVerticalWin(...params) ||
    checkDiagonal1Win(...params) ||
    checkDiagonal2Win(...params)
  )
}

// Solution of the excercise

// For any given board n by n, determine wether the last move is a winning move
expect(checkWinningMove(board_horizontal_win, [0, 1], 1)).to.equal(true)
expect(checkWinningMove(board_horizontal_win, [1, 0], 1)).to.equal(false)
