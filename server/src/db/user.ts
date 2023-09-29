import { prisma } from "@/prisma";

export async function getUser(username: string, password: string) {
  return await prisma.users.findFirst({
    where: {
      username,
      password,
    },
  });
}
