import { databaseErrorResolver } from "@/db/errors";
import { getGames } from "@/db/game";
import type { QueryResolvers } from "@/schema/types.generated";
import { createGraphQLError } from "graphql-yoga";

type BoardType = (string | null)[][];

export const getUserSudoku: NonNullable<QueryResolvers["getUserSudoku"]> = async (_parent, _arg, _ctx) => {
  let user = _ctx.user_id;
  if (_arg.user_id) user = _arg.user_id;
  if (!user) throw createGraphQLError("User not found");

  let results: Awaited<ReturnType<typeof getGames>> = [];
  let parsedBoards: BoardType[] = [];

  try {
    results = await getGames(user);
  } catch (err) {
    throw databaseErrorResolver(err);
  }

  try {
    parsedBoards = results.map((el) => JSON.parse(el.board) as unknown) as BoardType[];
  } catch (err) {
    throw createGraphQLError((err as Error).message);
  }

  return results.map((el, i) => ({
    ...el,
    board: parsedBoards[i],
  }));
};
