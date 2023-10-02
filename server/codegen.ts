import type { CodegenConfig } from "@graphql-codegen/cli";
import { defineConfig } from "@eddeee888/gcg-typescript-resolver-files";

const config: CodegenConfig = {
  schema: "**/.graphql",
  generates: {
    "src/schema": defineConfig({
      typesPluginsConfig: {
        contextType: "src/context/index#RequestContextType",
      },
    }),
  },
  emitLegacyCommonJSImports: false,
  hooks: {
    afterOneFileWrite: ["prettier --write"],
  },
};
export default config;
