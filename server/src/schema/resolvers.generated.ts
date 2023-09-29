/* This file was automatically generated. DO NOT UPDATE MANUALLY. */
    import type   { Resolvers } from './types.generated';
    import    { getNewSudoku as Query_getNewSudoku } from './sudoku/resolvers/Query/getNewSudoku';
import    { validateSudoku as Query_validateSudoku } from './sudoku/resolvers/Query/validateSudoku';
    export const resolvers: Resolvers = {
      Query: { getNewSudoku: Query_getNewSudoku,validateSudoku: Query_validateSudoku },
      
      
      
    }