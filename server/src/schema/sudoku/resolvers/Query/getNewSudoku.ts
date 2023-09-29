import type { QueryResolvers } from "@/schema/types.generated";
import { generateSudoku } from "@/sudoku/generator";
export const getNewSudoku: NonNullable<QueryResolvers["getNewSudoku"]> = async (_parent, _arg, _ctx) => {
  return generateSudoku(_arg.size, _arg.countOfBeginNumbers);
};
