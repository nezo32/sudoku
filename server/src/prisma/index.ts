import { PrismaClient } from "@prisma/client";
import dotenv from "dotenv";
import { readFileSync } from "fs";

dotenv.config();

const path = process.env.DB_URL_FILE;

if (!path) {
  throw new Error("DATABASE_URL_FILE env is not provided");
}

const databaseURL = readFileSync(path);

export const prisma = new PrismaClient({
  datasourceUrl: databaseURL.toString(),
});
