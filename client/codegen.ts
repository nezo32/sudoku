import type { CodegenConfig } from '@graphql-codegen/cli'
import dotenv from 'dotenv'

dotenv.config()

const config: CodegenConfig = {
  schema: process.env.VITE_GRAPHQL_URL,
  documents: ['src/**/*.vue'],
  ignoreNoDocuments: true,
  generates: {
    './src/generated/': {
      preset: 'client',

      config: {
        useTypeImports: true
      },
      presetConfig: {
        gqlTagName: 'gql'
      }
    }
  }
}

export default config
