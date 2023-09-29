import { isValidSudoku } from "../../../../utils/sudoku/validator";
import type { QueryResolvers } from "./../../../types.generated";
export const validateSudoku: NonNullable<QueryResolvers['validateSudoku']> = async (_parent, _arg, _ctx) => {
  return isValidSudoku(_arg.sudoku);
};
