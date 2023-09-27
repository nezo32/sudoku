import { generateSudoku } from "../../../../utils/sudoku";
import type { QueryResolvers } from "./../../../types.generated";
export const getNewSudoku: NonNullable<QueryResolvers["getNewSudoku"]> = async (_parent, _arg, _ctx) => {
  const tmp = generateSudoku(_arg.size, _arg.countOfBeginNumbers);
  return tmp;
};
