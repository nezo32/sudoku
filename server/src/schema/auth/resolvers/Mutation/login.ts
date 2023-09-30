import { getUser } from "@/db/user";
import type { MutationResolvers } from "@/schema/types.generated";
import { getHashedPassword } from "@/security/hash";
import { jwtSign } from "@/security/jwt";
import { createGraphQLError } from "graphql-yoga";
export const login: NonNullable<MutationResolvers["login"]> = async (_parent, _arg, _ctx) => {
  const { username, password } = _arg.data;

  const dbResult = await getUser(username);
  if (!dbResult) {
    throw createGraphQLError("User not found");
  }

  const incomingPasswordHash = await getHashedPassword(password, dbResult.salt);
  if (incomingPasswordHash !== dbResult.password) {
    throw createGraphQLError("Invalid password");
  }

  const token = jwtSign({ id: dbResult.id, username: dbResult.username });

  await _ctx.request.cookieStore?.set({
    name: "authorization",
    value: token,
    expires: new Date(new Date().getTime() * 1000 * 60 * 60 * 24),
    domain: "localhost",
    secure: true,
    sameSite: "strict",
  });

  return true;
};
