import { prisma } from "@/prisma";

export async function getUser(username: string) {
  return await prisma.users.findFirst({
    where: {
      username,
    },
  });
}

export async function createUser(username: string, password: string, salt: string) {
  return await prisma.users.create({
    data: {
      username,
      password,
      salt,
    },
  });
}
