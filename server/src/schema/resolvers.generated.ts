/* This file was automatically generated. DO NOT UPDATE MANUALLY. */
import type { Resolvers } from "./types.generated";
import { Game } from "./sudoku/resolvers/Game";
import { getNewSudoku as Mutation_getNewSudoku } from "./sudoku/resolvers/Mutation/getNewSudoku";
import { login as Mutation_login } from "./auth/resolvers/Mutation/login";
import { logout as Mutation_logout } from "./auth/resolvers/Mutation/logout";
import { register as Mutation_register } from "./auth/resolvers/Mutation/register";
import { getUserSudoku as Query_getUserSudoku } from "./sudoku/resolvers/Query/getUserSudoku";
import { validateSudoku as Query_validateSudoku } from "./sudoku/resolvers/Query/validateSudoku";
import { DateTimeResolver } from "graphql-scalars";
export const resolvers: Resolvers = {
  Query: { getUserSudoku: Query_getUserSudoku, validateSudoku: Query_validateSudoku },
  Mutation: {
    getNewSudoku: Mutation_getNewSudoku,
    login: Mutation_login,
    logout: Mutation_logout,
    register: Mutation_register,
  },

  Game: Game,
  DateTime: DateTimeResolver,
};
