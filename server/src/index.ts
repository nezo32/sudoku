import moduleAlias from "module-alias";
moduleAlias();

import { createSchema, createYoga } from "graphql-yoga";
import { typeDefs } from "@/schema/typeDefs.generated";
import { resolvers } from "@/schema/resolvers.generated";
import { createServer } from "node:http";
import { useJWT } from "@graphql-yoga/plugin-jwt";
import { useCookies } from "@whatwg-node/server-plugin-cookies";
import { jwtSecret } from "@/auth/jwt";
import { prisma } from "@/prisma";

const yoga = createYoga({
  schema: createSchema({
    typeDefs,
    resolvers,
  }),
  plugins: [
    useJWT({
      issuer: "http://graphql-yoga.com",
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

process.on("exit", async () => {
  await prisma.$disconnect();
});
