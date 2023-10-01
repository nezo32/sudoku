import { PrismaClientKnownRequestError, PrismaClientUnknownRequestError } from "@prisma/client/runtime/library";
import { createGraphQLError } from "graphql-yoga";

export function databaseErrorResolver(err: any) {
  if (err instanceof PrismaClientKnownRequestError) {
    return createGraphQLError(`${err.code} — ${err.message}`);
  }
  if (err instanceof PrismaClientUnknownRequestError) {
    return createGraphQLError(err.message);
  } else {
    return createGraphQLError((err as Error).message);
  }
}
