/* eslint-disable */
import type { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  /** A date-time string at UTC, such as 2007-12-03T10:15:30Z, compliant with the `date-time` format outlined in section 5.6 of the RFC 3339 profile of the ISO 8601 standard for representation of dates and times using the Gregorian calendar. */
  DateTime: { input: any; output: any; }
};

export type AuthData = {
  password: Scalars['String']['input'];
  username: Scalars['String']['input'];
};

export type Game = {
  __typename?: 'Game';
  begin_date?: Maybe<Scalars['DateTime']['output']>;
  board: Array<Array<Maybe<Scalars['String']['output']>>>;
  id?: Maybe<Scalars['ID']['output']>;
  player_id?: Maybe<Scalars['String']['output']>;
};

export type Mutation = {
  __typename?: 'Mutation';
  getNewSudoku: Array<Array<Maybe<Scalars['String']['output']>>>;
  login?: Maybe<Scalars['Boolean']['output']>;
  logout?: Maybe<Scalars['Boolean']['output']>;
  register?: Maybe<Scalars['Boolean']['output']>;
};


export type MutationGetNewSudokuArgs = {
  countOfBeginNumbers?: InputMaybe<Scalars['Int']['input']>;
  size?: InputMaybe<Scalars['Int']['input']>;
};


export type MutationLoginArgs = {
  data: AuthData;
};


export type MutationRegisterArgs = {
  data: AuthData;
};

export type Query = {
  __typename?: 'Query';
  getUserSudoku: Array<Maybe<Game>>;
  validateSudoku?: Maybe<Scalars['Boolean']['output']>;
};


export type QueryGetUserSudokuArgs = {
  user_id?: InputMaybe<Scalars['String']['input']>;
};


export type QueryValidateSudokuArgs = {
  sudoku: Array<Array<InputMaybe<Scalars['String']['input']>>>;
};

export type LoginMutationVariables = Exact<{
  data: AuthData;
}>;


export type LoginMutation = { __typename?: 'Mutation', login?: boolean | null };

export type RegisterMutationVariables = Exact<{
  data: AuthData;
}>;


export type RegisterMutation = { __typename?: 'Mutation', register?: boolean | null };


export const LoginDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"Login"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"data"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"AuthData"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"login"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"data"},"value":{"kind":"Variable","name":{"kind":"Name","value":"data"}}}]}]}}]} as unknown as DocumentNode<LoginMutation, LoginMutationVariables>;
export const RegisterDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"Register"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"data"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"AuthData"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"register"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"data"},"value":{"kind":"Variable","name":{"kind":"Name","value":"data"}}}]}]}}]} as unknown as DocumentNode<RegisterMutation, RegisterMutationVariables>;