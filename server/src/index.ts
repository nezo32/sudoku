import moduleAlias from "module-alias";
moduleAlias();

import { createSchema, createYoga } from "graphql-yoga";
import { typeDefs } from "@/schema/typeDefs.generated";
import { resolvers } from "@/schema/resolvers.generated";
import { createServer } from "node:http";
import { useJWT } from "@graphql-yoga/plugin-jwt";
import { useCookies } from "@whatwg-node/server-plugin-cookies";
import { jwtSecret } from "@/security/jwt";
import { prisma } from "@/prisma";
import { applyMiddleware } from "graphql-middleware";
import { jwtCheck } from "@/middleware/jwt";

process.on("exit", async () => {
  await prisma.$disconnect();
});

async function main() {
  await prisma.$connect();

  const middlewareSchema = applyMiddleware(
    createSchema({
      typeDefs,
      resolvers,
    }),
    jwtCheck,
  );

  const yoga = createYoga({
    schema: middlewareSchema,
    plugins: [
      useJWT({
        issuer: process.env.JWT_ISSUER ?? "noone",
        signingKey: jwtSecret.toString(),
        algorithms: ["RS256"],
        async getToken(params) {
          return (await params.request.cookieStore?.get("authorization"))?.value;
        },
      }),
      useCookies(),
    ],
  });

  const server = createServer(yoga);

  server.listen(4000, () => {
    console.info("Server is running on http://localhost:4000/graphql");
  });
}

main().then();
