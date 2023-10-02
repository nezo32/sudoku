import { databaseErrorResolver } from "@/db/errors";
import { createGame } from "@/db/game";
import { generateSudoku } from "@/sudoku/generator";
import { createGraphQLError } from "graphql-yoga";
import { MutationResolvers } from "@/schema/types.generated";

export const getNewSudoku: NonNullable<MutationResolvers["getNewSudoku"]> = async (_parent, _arg, _ctx) => {
  const board = generateSudoku(_arg.size, _arg.countOfBeginNumbers);
  const user = _ctx.user_id;

  if (!user) throw createGraphQLError("User not found");

  try {
    const stringifyBoard = JSON.stringify(board);
    await createGame(user, stringifyBoard);
  } catch (err) {
    throw databaseErrorResolver(err);
  }

  return board;
};
