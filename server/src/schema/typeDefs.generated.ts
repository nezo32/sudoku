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
        {
          kind: "FieldDefinition",
          name: { kind: "Name", value: "logout", loc: { start: 199, end: 205 } },
          arguments: [],
          type: {
            kind: "NamedType",
            name: { kind: "Name", value: "Boolean", loc: { start: 207, end: 214 } },
            loc: { start: 207, end: 214 },
          },
          directives: [],
          loc: { start: 199, end: 214 },
        },
      ],
      loc: { start: 103, end: 216 },
    },
    {
      kind: "ObjectTypeDefinition",
      name: { kind: "Name", value: "Game", loc: { start: 222, end: 226 } },
      interfaces: [],
      directives: [],
      fields: [
        {
          kind: "FieldDefinition",
          name: { kind: "Name", value: "id", loc: { start: 231, end: 233 } },
          arguments: [],
          type: {
            kind: "NamedType",
            name: { kind: "Name", value: "ID", loc: { start: 235, end: 237 } },
            loc: { start: 235, end: 237 },
          },
          directives: [],
          loc: { start: 231, end: 237 },
        },
        {
          kind: "FieldDefinition",
          name: { kind: "Name", value: "board", loc: { start: 240, end: 245 } },
          arguments: [],
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
                    name: { kind: "Name", value: "String", loc: { start: 249, end: 255 } },
                    loc: { start: 249, end: 255 },
                  },
                  loc: { start: 248, end: 256 },
                },
                loc: { start: 248, end: 257 },
              },
              loc: { start: 247, end: 258 },
            },
            loc: { start: 247, end: 259 },
          },
          directives: [],
          loc: { start: 240, end: 259 },
        },
        {
          kind: "FieldDefinition",
          name: { kind: "Name", value: "player_id", loc: { start: 262, end: 271 } },
          arguments: [],
          type: {
            kind: "NamedType",
            name: { kind: "Name", value: "String", loc: { start: 273, end: 279 } },
            loc: { start: 273, end: 279 },
          },
          directives: [],
          loc: { start: 262, end: 279 },
        },
        {
          kind: "FieldDefinition",
          name: { kind: "Name", value: "begin_date", loc: { start: 282, end: 292 } },
          arguments: [],
          type: {
            kind: "NamedType",
            name: { kind: "Name", value: "DateTime", loc: { start: 294, end: 302 } },
            loc: { start: 294, end: 302 },
          },
          directives: [],
          loc: { start: 282, end: 302 },
        },
      ],
      loc: { start: 217, end: 304 },
    },
    {
      kind: "ObjectTypeExtension",
      name: { kind: "Name", value: "Query", loc: { start: 318, end: 323 } },
      interfaces: [],
      directives: [],
      fields: [
        {
          kind: "FieldDefinition",
          name: { kind: "Name", value: "validateSudoku", loc: { start: 328, end: 342 } },
          arguments: [
            {
              kind: "InputValueDefinition",
              name: { kind: "Name", value: "sudoku", loc: { start: 343, end: 349 } },
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
                        name: { kind: "Name", value: "String", loc: { start: 353, end: 359 } },
                        loc: { start: 353, end: 359 },
                      },
                      loc: { start: 352, end: 360 },
                    },
                    loc: { start: 352, end: 361 },
                  },
                  loc: { start: 351, end: 362 },
                },
                loc: { start: 351, end: 363 },
              },
              directives: [],
              loc: { start: 343, end: 363 },
            },
          ],
          type: {
            kind: "NamedType",
            name: { kind: "Name", value: "Boolean", loc: { start: 366, end: 373 } },
            loc: { start: 366, end: 373 },
          },
          directives: [],
          loc: { start: 328, end: 373 },
        },
        {
          kind: "FieldDefinition",
          name: { kind: "Name", value: "getUserSudoku", loc: { start: 376, end: 389 } },
          arguments: [
            {
              kind: "InputValueDefinition",
              name: { kind: "Name", value: "user_id", loc: { start: 390, end: 397 } },
              type: {
                kind: "NamedType",
                name: { kind: "Name", value: "String", loc: { start: 399, end: 405 } },
                loc: { start: 399, end: 405 },
              },
              directives: [],
              loc: { start: 390, end: 405 },
            },
          ],
          type: {
            kind: "NonNullType",
            type: {
              kind: "ListType",
              type: {
                kind: "NamedType",
                name: { kind: "Name", value: "Game", loc: { start: 409, end: 413 } },
                loc: { start: 409, end: 413 },
              },
              loc: { start: 408, end: 414 },
            },
            loc: { start: 408, end: 415 },
          },
          directives: [],
          loc: { start: 376, end: 415 },
        },
      ],
      loc: { start: 306, end: 417 },
    },
    {
      kind: "ObjectTypeExtension",
      name: { kind: "Name", value: "Mutation", loc: { start: 431, end: 439 } },
      interfaces: [],
      directives: [],
      fields: [
        {
          kind: "FieldDefinition",
          name: { kind: "Name", value: "getNewSudoku", loc: { start: 444, end: 456 } },
          arguments: [
            {
              kind: "InputValueDefinition",
              name: { kind: "Name", value: "size", loc: { start: 457, end: 461 } },
              type: {
                kind: "NamedType",
                name: { kind: "Name", value: "Int", loc: { start: 463, end: 466 } },
                loc: { start: 463, end: 466 },
              },
              defaultValue: { kind: "IntValue", value: "9", loc: { start: 469, end: 470 } },
              directives: [],
              loc: { start: 457, end: 470 },
            },
            {
              kind: "InputValueDefinition",
              name: { kind: "Name", value: "countOfBeginNumbers", loc: { start: 472, end: 491 } },
              type: {
                kind: "NamedType",
                name: { kind: "Name", value: "Int", loc: { start: 493, end: 496 } },
                loc: { start: 493, end: 496 },
              },
              defaultValue: { kind: "IntValue", value: "45", loc: { start: 499, end: 501 } },
              directives: [],
              loc: { start: 472, end: 501 },
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
                    name: { kind: "Name", value: "String", loc: { start: 506, end: 512 } },
                    loc: { start: 506, end: 512 },
                  },
                  loc: { start: 505, end: 513 },
                },
                loc: { start: 505, end: 514 },
              },
              loc: { start: 504, end: 515 },
            },
            loc: { start: 504, end: 516 },
          },
          directives: [],
          loc: { start: 444, end: 516 },
        },
      ],
      loc: { start: 419, end: 518 },
    },
  ],
  loc: { start: 0, end: 519 },
} as unknown as DocumentNode;
