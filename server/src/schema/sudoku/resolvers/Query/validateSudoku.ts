import type { QueryResolvers } from "@/schema/types.generated";
import { isValidSudoku } from "@/sudoku/validator";
export const validateSudoku: NonNullable<QueryResolvers["validateSudoku"]> = async (_parent, _arg, _ctx) => {
  return isValidSudoku(_arg.sudoku);
};
