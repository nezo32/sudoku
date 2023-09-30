import type { DocumentNode } from "graphql";
export const typeDefs = {
  kind: "Document",
  definitions: [
    {
      kind: "ObjectTypeDefinition",
      name: { kind: "Name", value: "Query", loc: { start: 5, end: 10 } },
      interfaces: [],
      directives: [],
      fields: [],
      loc: { start: 0, end: 10 },
    },
    {
      kind: "ObjectTypeDefinition",
      name: { kind: "Name", value: "Mutation", loc: { start: 17, end: 25 } },
      interfaces: [],
      directives: [],
      fields: [],
      loc: { start: 12, end: 25 },
    },
    {
      kind: "ScalarTypeDefinition",
      name: { kind: "Name", value: "DateTime", loc: { start: 34, end: 42 } },
      directives: [],
      loc: { start: 27, end: 42 },
    },
    {
      kind: "InputObjectTypeDefinition",
      name: { kind: "Name", value: "AuthData", loc: { start: 49, end: 57 } },
      directives: [],
      fields: [
        {
          kind: "InputValueDefinition",
          name: { kind: "Name", value: "username", loc: { start: 62, end: 70 } },
          type: {
            kind: "NonNullType",
            type: {
              kind: "NamedType",
              name: { kind: "Name", value: "String", loc: { start: 72, end: 78 } },
              loc: { start: 72, end: 78 },
            },
            loc: { start: 72, end: 79 },
          },
          directives: [],
          loc: { start: 62, end: 79 },
        },
        {
          kind: "InputValueDefinition",
          name: { kind: "Name", value: "password", loc: { start: 82, end: 90 } },
          type: {
            kind: "NonNullType",
            type: {
              kind: "NamedType",
              name: { kind: "Name", value: "String", loc: { start: 92, end: 98 } },
              loc: { start: 92, end: 98 },
            },
            loc: { start: 92, end: 99 },
          },
          directives: [],
          loc: { start: 82, end: 99 },
        },
      ],
      loc: { start: 43, end: 101 },
    },
    {
      kind: "ObjectTypeExtension",
      name: { kind: "Name", value: "Mutation", loc: { start: 115, end: 123 } },
      interfaces: [],
      directives: [],
      fields: [
        {
          kind: "FieldDefinition",
          name: { kind: "Name", value: "login", loc: { start: 128, end: 133 } },
          arguments: [
            {
              kind: "InputValueDefinition",
              name: { kind: "Name", value: "data", loc: { start: 134, end: 138 } },
              type: {
                kind: "NonNullType",
                type: {
                  kind: "NamedType",
                  name: { kind: "Name", value: "AuthData", loc: { start: 140, end: 148 } },
                  loc: { start: 140, end: 148 },
                },
                loc: { start: 140, end: 149 },
              },
              directives: [],
              loc: { start: 134, end: 149 },
            },
          ],
          type: {
            kind: "NamedType",
            name: { kind: "Name", value: "Boolean", loc: { start: 152, end: 159 } },
            loc: { start: 152, end: 159 },
          },
          directives: [],
          loc: { start: 128, end: 159 },
        },
        {
          kind: "FieldDefinition",
          name: { kind: "Name", value: "register", loc: { start: 162, end: 170 } },
          arguments: [
            {
              kind: "InputValueDefinition",
              name: { kind: "Name", value: "data", loc: { start: 171, end: 175 } },
              type: {
                kind: "NonNullType",
                type: {
                  kind: "NamedType",
                  name: { kind: "Name", value: "AuthData", loc: { start: 177, end: 185 } },
                  loc: { start: 177, end: 185 },
                },
                loc: { start: 177, end: 186 },
              },
              directives: [],
              loc: { start: 171, end: 186 },
            },
          ],
          type: {
            kind: "NamedType",
            name: { kind: "Name", value: "Boolean", loc: { start: 189, end: 196 } },
            loc: { start: 189, end: 196 },
          },
          directives: [],
          loc: { start: 162, end: 196 },
        },
      ],
      loc: { start: 103, end: 198 },
    },
    {
      kind: "ObjectTypeExtension",
      name: { kind: "Name", value: "Query", loc: { start: 211, end: 216 } },
      interfaces: [],
      directives: [],
      fields: [
        {
          kind: "FieldDefinition",
          name: { kind: "Name", value: "getNewSudoku", loc: { start: 221, end: 233 } },
          arguments: [
            {
              kind: "InputValueDefinition",
              name: { kind: "Name", value: "size", loc: { start: 234, end: 238 } },
              type: {
                kind: "NamedType",
                name: { kind: "Name", value: "Int", loc: { start: 240, end: 243 } },
                loc: { start: 240, end: 243 },
              },
              defaultValue: { kind: "IntValue", value: "9", loc: { start: 246, end: 247 } },
              directives: [],
              loc: { start: 234, end: 247 },
            },
            {
              kind: "InputValueDefinition",
              name: { kind: "Name", value: "countOfBeginNumbers", loc: { start: 249, end: 268 } },
              type: {
                kind: "NamedType",
                name: { kind: "Name", value: "Int", loc: { start: 270, end: 273 } },
                loc: { start: 270, end: 273 },
              },
              defaultValue: { kind: "IntValue", value: "45", loc: { start: 276, end: 278 } },
              directives: [],
              loc: { start: 249, end: 278 },
            },
          ],
          type: {
            kind: "NonNullType",
            type: {
              kind: "ListType",
              type: {
                kind: "NonNullType",
                type: {
                  kind: "ListType",
                  type: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "String", loc: { start: 283, end: 289 } },
                    loc: { start: 283, end: 289 },
                  },
                  loc: { start: 282, end: 290 },
                },
                loc: { start: 282, end: 291 },
              },
              loc: { start: 281, end: 292 },
            },
            loc: { start: 281, end: 293 },
          },
          directives: [],
          loc: { start: 221, end: 293 },
        },
        {
          kind: "FieldDefinition",
          name: { kind: "Name", value: "validateSudoku", loc: { start: 296, end: 310 } },
          arguments: [
            {
              kind: "InputValueDefinition",
              name: { kind: "Name", value: "sudoku", loc: { start: 311, end: 317 } },
              type: {
                kind: "NonNullType",
                type: {
                  kind: "ListType",
                  type: {
                    kind: "NonNullType",
                    type: {
                      kind: "ListType",
                      type: {
                        kind: "NamedType",
                        name: { kind: "Name", value: "String", loc: { start: 321, end: 327 } },
                        loc: { start: 321, end: 327 },
                      },
                      loc: { start: 320, end: 328 },
                    },
                    loc: { start: 320, end: 329 },
                  },
                  loc: { start: 319, end: 330 },
                },
                loc: { start: 319, end: 331 },
              },
              directives: [],
              loc: { start: 311, end: 331 },
            },
          ],
          type: {
            kind: "NamedType",
            name: { kind: "Name", value: "Boolean", loc: { start: 334, end: 341 } },
            loc: { start: 334, end: 341 },
          },
          directives: [],
          loc: { start: 296, end: 341 },
        },
      ],
      loc: { start: 199, end: 343 },
    },
  ],
  loc: { start: 0, end: 344 },
} as unknown as DocumentNode;
