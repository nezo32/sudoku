import { RequestContextType } from "@/context";
import { jwtVerify } from "@/security/jwt";
import type { IMiddleware } from "graphql-middleware";
import { createGraphQLError } from "graphql-yoga";

const ENDPOINT_SPIP = ["login", "register", "logout"] as const;

export const jwtCheck: IMiddleware<any, RequestContextType> = async (resolver, parent, args, context, info) => {
  const endpoint = info.path.key;
  let user_id: string | undefined;
  if (!ENDPOINT_SPIP.some((el) => el === endpoint)) {
    const cookie = await context.request.cookieStore?.get("authorization");

    if (!cookie) throw createGraphQLError("Unauthorized");

    try {
      const jwt = jwtVerify(cookie.value);
      if (typeof jwt === "object") {
        user_id = jwt.sub;
      }
    } catch (err) {
      throw createGraphQLError((err as Error).message);
    }
  }

  const result = resolver(parent, args, { ...context, user_id }, info);
  return result;
};
