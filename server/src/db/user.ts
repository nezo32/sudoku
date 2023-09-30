import { prisma } from "@/prisma";

export async function getUser(username: string) {
  return await prisma.users.findFirst({
    where: {
      username,
    },
  });
}
