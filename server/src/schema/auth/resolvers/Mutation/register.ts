import { databaseErrorResolver } from "@/db/errors";
import { createUser } from "@/db/user";
import { generatePasswordHash } from "@/security/hash";
import type { MutationResolvers } from "@/schema/types.generated";

export const register: NonNullable<MutationResolvers["register"]> = async (_parent, _arg, _ctx) => {
  const { username, password } = _arg.data;
  const passHash = await generatePasswordHash(password);

  try {
    await createUser(username, passHash.hash, passHash.salt);
  } catch (err) {
    throw databaseErrorResolver(err);
  }

  return true;
};
