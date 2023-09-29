import type { MutationResolvers } from "@/schema/types.generated";
export const register: NonNullable<MutationResolvers["register"]> = async (_parent, _arg, _ctx) => {
  return true;
};
