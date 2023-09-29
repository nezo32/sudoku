/* This file was automatically generated. DO NOT UPDATE MANUALLY. */
import type { Resolvers } from "./types.generated";
import { login as Mutation_login } from "./auth/resolvers/Mutation/login";
import { register as Mutation_register } from "./auth/resolvers/Mutation/register";
import { getNewSudoku as Query_getNewSudoku } from "./sudoku/resolvers/Query/getNewSudoku";
import { validateSudoku as Query_validateSudoku } from "./sudoku/resolvers/Query/validateSudoku";
import { DateTimeResolver } from "graphql-scalars";
export const resolvers: Resolvers = {
  Query: { getNewSudoku: Query_getNewSudoku, validateSudoku: Query_validateSudoku },
  Mutation: { login: Mutation_login, register: Mutation_register },

  DateTime: DateTimeResolver,
};
