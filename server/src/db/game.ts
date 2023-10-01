import { prisma } from "@/prisma";

export async function createGame(user: string, board: string) {
  return await prisma.games.create({
    data: {
      board,
      users: {
        connect: {
          id: user,
        },
      },
    },
  });
}

export async function getGame(id: string, player_id: string) {
  return await prisma.games.findFirstOrThrow({
    where: {
      id,
      player_id,
    },
  });
}

export async function getGames(player_id: string) {
  return await prisma.games.findMany({
    where: {
      player_id,
    },
  });
}
