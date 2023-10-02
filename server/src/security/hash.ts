import bcrypt from "bcrypt";

export async function generatePasswordHash(password: string) {
  const salt = await bcrypt.genSalt();
  const hash = await bcrypt.hash(password, salt);

  return { hash, salt };
}

export async function getHashedPassword(password: string, salt: string) {
  const hash = await bcrypt.hash(password, salt);
  return hash;
}
