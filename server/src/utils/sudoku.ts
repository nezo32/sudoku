export function generateSudoku(size: number, countOfBeginNumbers: number): (string | null)[][] {
  if (size % 3 !== 0 || size === 0) return [];

  const board: (string | null)[][] = createEmptyArray(size).map(() => createEmptyArray(size).map(() => null));
  const rows = createArrayOfSets(size);
  const columns = createArrayOfSets(size);
  const boxes = createArrayOfSets(size);

  function backtrack(row = 0, column = 0) {
    if (row >= board.length || column >= board.length) return true;

    if (board[row][column] == null) {
      const possibleNums = new Array(9)
        .fill(0)
        .map((_, i) => i + 1)
        .sort(() => Math.random() - 0.5);
      for (let i = 0; i < possibleNums.length; i++) {
        let randIndex = Math.floor(Math.random() * possibleNums.length);
        if (possibleNums.length == 1) {
          randIndex = 0;
        }
        const value = possibleNums[randIndex].toString();
        if (isValidPlacement(row, column, value)) {
          setNewValue(row, column, value);

          const [rNext, cNext] = getNextRowColumn(row, column);
          if (backtrack(rNext, cNext)) {
            board[row][column] = value;
            return true;
          }
          removeValue(row, column, value);
        }
      }
    } else {
      const [rNext, cNext] = getNextRowColumn(row, column);
      return backtrack(rNext, cNext);
    }

    return false;
  }
  backtrack();

  const countOfCells = size ** 2;
  let done = 0;
  while (done < countOfCells - countOfBeginNumbers) {
    const xRand = Math.floor(Math.random() * size);
    const yRand = Math.floor(Math.random() * size);
    if (board[xRand][yRand]) {
      board[xRand][yRand] = null;
      done++;
    }
  }

  return board;

  function removeValue(row: number, column: number, value: string) {
    rows[row].delete(value);
    columns[column].delete(value);
    boxes[getBoxId(row, column)].delete(value);
  }
  function getNextRowColumn(row: number, column: number): [number, number] {
    if (column === board.length - 1) {
      return [row + 1, 0];
    } else {
      return [row, column + 1];
    }
  }
  function isValidPlacement(row: number, column: number, value: string): boolean {
    return !rows[row].has(value) && !columns[column].has(value) && !boxes[getBoxId(row, column)].has(value);
  }
  function setNewValue(row: number, column: number, value: string) {
    rows[row].add(value);
    columns[column].add(value);
    boxes[getBoxId(row, column)].add(value);
  }
}

function getBoxId(x: number, y: number) {
  return 3 * Math.floor(x / 3) + Math.floor(y / 3);
}

function createEmptyArray(n: number) {
  return new Array(n).fill(undefined);
}
function createArrayOfSets(n: number) {
  return createEmptyArray(n).map(() => new Set<string>());
}
