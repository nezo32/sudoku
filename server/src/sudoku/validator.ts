import { getBoxId, createArrayOfSets } from "./general";

export function isValidSudoku(board: (string | null | undefined)[][]): boolean {
  let result = true;

  const columns = createArrayOfSets(board.length);
  const boxes = createArrayOfSets(board.length);
  const rows = createArrayOfSets(board.length);

  for (let i = 0; i < board.length; i++) {
    if (!result) break;
    for (let j = 0; j < board[i].length; j++) {
      const value = board[i][j];

      if (!value) return false;

      const isInvalid = rows[i].has(value) || columns[j].has(value) || boxes[getBoxId(i, j)].has(value);
      if (isInvalid) {
        return false;
      } else {
        rows[i].add(value);
        columns[j].add(value);
        boxes[getBoxId(i, j)].add(value);
      }
    }
  }

  return result;
}
