export function getBoxId(x: number, y: number) {
  return 3 * Math.floor(x / 3) + Math.floor(y / 3);
}

export function createEmptyArray(n: number) {
  return new Array(n).fill(undefined);
}

export function createArrayOfSets(n: number) {
  return createEmptyArray(n).map(() => new Set<string>());
}
