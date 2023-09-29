import type { DocumentNode } from 'graphql';
  export const typeDefs = {"kind":"Document","definitions":[{"kind":"ObjectTypeDefinition","name":{"kind":"Name","value":"Query","loc":{"start":5,"end":10}},"interfaces":[],"directives":[],"fields":[],"loc":{"start":0,"end":10}},{"kind":"ObjectTypeExtension","name":{"kind":"Name","value":"Query","loc":{"start":23,"end":28}},"interfaces":[],"directives":[],"fields":[{"kind":"FieldDefinition","name":{"kind":"Name","value":"getNewSudoku","loc":{"start":33,"end":45}},"arguments":[{"kind":"InputValueDefinition","name":{"kind":"Name","value":"size","loc":{"start":46,"end":50}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"Int","loc":{"start":52,"end":55}},"loc":{"start":52,"end":55}},"defaultValue":{"kind":"IntValue","value":"9","loc":{"start":58,"end":59}},"directives":[],"loc":{"start":46,"end":59}},{"kind":"InputValueDefinition","name":{"kind":"Name","value":"countOfBeginNumbers","loc":{"start":61,"end":80}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"Int","loc":{"start":82,"end":85}},"loc":{"start":82,"end":85}},"defaultValue":{"kind":"IntValue","value":"45","loc":{"start":88,"end":90}},"directives":[],"loc":{"start":61,"end":90}}],"type":{"kind":"NonNullType","type":{"kind":"ListType","type":{"kind":"NonNullType","type":{"kind":"ListType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String","loc":{"start":95,"end":101}},"loc":{"start":95,"end":101}},"loc":{"start":94,"end":102}},"loc":{"start":94,"end":103}},"loc":{"start":93,"end":104}},"loc":{"start":93,"end":105}},"directives":[],"loc":{"start":33,"end":105}},{"kind":"FieldDefinition","name":{"kind":"Name","value":"validateSudoku","loc":{"start":108,"end":122}},"arguments":[{"kind":"InputValueDefinition","name":{"kind":"Name","value":"sudoku","loc":{"start":123,"end":129}},"type":{"kind":"NonNullType","type":{"kind":"ListType","type":{"kind":"NonNullType","type":{"kind":"ListType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String","loc":{"start":133,"end":139}},"loc":{"start":133,"end":139}},"loc":{"start":132,"end":140}},"loc":{"start":132,"end":141}},"loc":{"start":131,"end":142}},"loc":{"start":131,"end":143}},"directives":[],"loc":{"start":123,"end":143}}],"type":{"kind":"NamedType","name":{"kind":"Name","value":"Boolean","loc":{"start":146,"end":153}},"loc":{"start":146,"end":153}},"directives":[],"loc":{"start":108,"end":153}}],"loc":{"start":11,"end":155}}],"loc":{"start":0,"end":156}} as unknown as DocumentNode