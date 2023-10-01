import { YogaInitialContext } from "graphql-yoga";

export type RequestContextType = YogaInitialContext & { user_id?: string };
