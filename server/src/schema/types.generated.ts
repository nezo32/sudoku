import { GraphQLResolveInfo, GraphQLScalarType, GraphQLScalarTypeConfig } from "graphql";
import { RequestContextType } from "src/context/index";
export type Maybe<T> = T | null | undefined;
export type InputMaybe<T> = T | null | undefined;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends " $fragmentName" | "__typename" ? T[P] : never };
export type RequireFields<T, K extends keyof T> = Omit<T, K> & { [P in K]-?: NonNullable<T[P]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string | number };
  String: { input: string; output: string };
  Boolean: { input: boolean; output: boolean };
  Int: { input: number; output: number };
  Float: { input: number; output: number };
  DateTime: { input: Date | string; output: Date | string };
};

export type AuthData = {
  password: Scalars["String"]["input"];
  username: Scalars["String"]["input"];
};

export type Game = {
  __typename?: "Game";
  begin_date?: Maybe<Scalars["DateTime"]["output"]>;
  board: Array<Array<Maybe<Scalars["String"]["output"]>>>;
  id?: Maybe<Scalars["ID"]["output"]>;
  player_id?: Maybe<Scalars["String"]["output"]>;
};

export type Mutation = {
  __typename?: "Mutation";
  getNewSudoku: Array<Array<Maybe<Scalars["String"]["output"]>>>;
  login?: Maybe<Scalars["Boolean"]["output"]>;
  logout?: Maybe<Scalars["Boolean"]["output"]>;
  register?: Maybe<Scalars["Boolean"]["output"]>;
};

export type MutationgetNewSudokuArgs = {
  countOfBeginNumbers?: InputMaybe<Scalars["Int"]["input"]>;
  size?: InputMaybe<Scalars["Int"]["input"]>;
};

export type MutationloginArgs = {
  data: AuthData;
};

export type MutationregisterArgs = {
  data: AuthData;
};

export type Query = {
  __typename?: "Query";
  getUserSudoku: Array<Maybe<Game>>;
  validateSudoku?: Maybe<Scalars["Boolean"]["output"]>;
};

export type QuerygetUserSudokuArgs = {
  user_id?: InputMaybe<Scalars["String"]["input"]>;
};

export type QueryvalidateSudokuArgs = {
  sudoku: Array<Array<InputMaybe<Scalars["String"]["input"]>>>;
};

export type ResolverTypeWrapper<T> = Promise<T> | T;

export type ResolverWithResolve<TResult, TParent, TContext, TArgs> = {
  resolve: ResolverFn<TResult, TParent, TContext, TArgs>;
};
export type Resolver<TResult, TParent = {}, TContext = {}, TArgs = {}> =
  | ResolverFn<TResult, TParent, TContext, TArgs>
  | ResolverWithResolve<TResult, TParent, TContext, TArgs>;

export type ResolverFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo,
) => Promise<TResult> | TResult;

export type SubscriptionSubscribeFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo,
) => AsyncIterable<TResult> | Promise<AsyncIterable<TResult>>;

export type SubscriptionResolveFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo,
) => TResult | Promise<TResult>;

export interface SubscriptionSubscriberObject<TResult, TKey extends string, TParent, TContext, TArgs> {
  subscribe: SubscriptionSubscribeFn<{ [key in TKey]: TResult }, TParent, TContext, TArgs>;
  resolve?: SubscriptionResolveFn<TResult, { [key in TKey]: TResult }, TContext, TArgs>;
}

export interface SubscriptionResolverObject<TResult, TParent, TContext, TArgs> {
  subscribe: SubscriptionSubscribeFn<any, TParent, TContext, TArgs>;
  resolve: SubscriptionResolveFn<TResult, any, TContext, TArgs>;
}

export type SubscriptionObject<TResult, TKey extends string, TParent, TContext, TArgs> =
  | SubscriptionSubscriberObject<TResult, TKey, TParent, TContext, TArgs>
  | SubscriptionResolverObject<TResult, TParent, TContext, TArgs>;

export type SubscriptionResolver<TResult, TKey extends string, TParent = {}, TContext = {}, TArgs = {}> =
  | ((...args: any[]) => SubscriptionObject<TResult, TKey, TParent, TContext, TArgs>)
  | SubscriptionObject<TResult, TKey, TParent, TContext, TArgs>;

export type TypeResolveFn<TTypes, TParent = {}, TContext = {}> = (
  parent: TParent,
  context: TContext,
  info: GraphQLResolveInfo,
) => Maybe<TTypes> | Promise<Maybe<TTypes>>;

export type IsTypeOfResolverFn<T = {}, TContext = {}> = (
  obj: T,
  context: TContext,
  info: GraphQLResolveInfo,
) => boolean | Promise<boolean>;

export type NextResolverFn<T> = () => Promise<T>;

export type DirectiveResolverFn<TResult = {}, TParent = {}, TContext = {}, TArgs = {}> = (
  next: NextResolverFn<TResult>,
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo,
) => TResult | Promise<TResult>;

/** Mapping between all available schema types and the resolvers types */
export type ResolversTypes = {
  AuthData: AuthData;
  String: ResolverTypeWrapper<Scalars["String"]["output"]>;
  DateTime: ResolverTypeWrapper<Scalars["DateTime"]["output"]>;
  Game: ResolverTypeWrapper<Game>;
  ID: ResolverTypeWrapper<Scalars["ID"]["output"]>;
  Mutation: ResolverTypeWrapper<{}>;
  Int: ResolverTypeWrapper<Scalars["Int"]["output"]>;
  Boolean: ResolverTypeWrapper<Scalars["Boolean"]["output"]>;
  Query: ResolverTypeWrapper<{}>;
};

/** Mapping between all available schema types and the resolvers parents */
export type ResolversParentTypes = {
  AuthData: AuthData;
  String: Scalars["String"]["output"];
  DateTime: Scalars["DateTime"]["output"];
  Game: Game;
  ID: Scalars["ID"]["output"];
  Mutation: {};
  Int: Scalars["Int"]["output"];
  Boolean: Scalars["Boolean"]["output"];
  Query: {};
};

export interface DateTimeScalarConfig extends GraphQLScalarTypeConfig<ResolversTypes["DateTime"], any> {
  name: "DateTime";
}

export type GameResolvers<
  ContextType = RequestContextType,
  ParentType extends ResolversParentTypes["Game"] = ResolversParentTypes["Game"],
> = {
  begin_date?: Resolver<Maybe<ResolversTypes["DateTime"]>, ParentType, ContextType>;
  board?: Resolver<Array<Array<Maybe<ResolversTypes["String"]>>>, ParentType, ContextType>;
  id?: Resolver<Maybe<ResolversTypes["ID"]>, ParentType, ContextType>;
  player_id?: Resolver<Maybe<ResolversTypes["String"]>, ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
};

export type MutationResolvers<
  ContextType = RequestContextType,
  ParentType extends ResolversParentTypes["Mutation"] = ResolversParentTypes["Mutation"],
> = {
  getNewSudoku?: Resolver<
    Array<Array<Maybe<ResolversTypes["String"]>>>,
    ParentType,
    ContextType,
    RequireFields<MutationgetNewSudokuArgs, "countOfBeginNumbers" | "size">
  >;
  login?: Resolver<Maybe<ResolversTypes["Boolean"]>, ParentType, ContextType, RequireFields<MutationloginArgs, "data">>;
  logout?: Resolver<Maybe<ResolversTypes["Boolean"]>, ParentType, ContextType>;
  register?: Resolver<
    Maybe<ResolversTypes["Boolean"]>,
    ParentType,
    ContextType,
    RequireFields<MutationregisterArgs, "data">
  >;
};

export type QueryResolvers<
  ContextType = RequestContextType,
  ParentType extends ResolversParentTypes["Query"] = ResolversParentTypes["Query"],
> = {
  getUserSudoku?: Resolver<
    Array<Maybe<ResolversTypes["Game"]>>,
    ParentType,
    ContextType,
    Partial<QuerygetUserSudokuArgs>
  >;
  validateSudoku?: Resolver<
    Maybe<ResolversTypes["Boolean"]>,
    ParentType,
    ContextType,
    RequireFields<QueryvalidateSudokuArgs, "sudoku">
  >;
};

export type Resolvers<ContextType = RequestContextType> = {
  DateTime?: GraphQLScalarType;
  Game?: GameResolvers<ContextType>;
  Mutation?: MutationResolvers<ContextType>;
  Query?: QueryResolvers<ContextType>;
};
