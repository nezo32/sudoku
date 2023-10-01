import jwt from "jsonwebtoken";
import { UserCredentials } from "@/types/auth";
import { readFileSync } from "fs";
import dotenv from "dotenv";

dotenv.config();

const privateKeyPath = process.env.JWT_SECRET;

if (!privateKeyPath) {
  throw new Error("Path to JWT secret is not provided");
}

export const jwtSecret = readFileSync(privateKeyPath);

export function jwtSign(user: UserCredentials, expiresIn: string) {
  const { username, id } = user;
  return jwt.sign({ username }, jwtSecret, {
    issuer: process.env.JWT_ISSUER ?? "noone",
    subject: id,
    algorithm: "RS256",
    expiresIn,
  });
}

export function jwtVerify(token: string) {
  return jwt.verify(token, jwtSecret, {
    algorithms: ["RS256"],
    issuer: process.env.JWT_ISSUER ?? "noone",
  });
}
