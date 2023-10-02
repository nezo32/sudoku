import type { MutationResolvers } from "@/schema/types.generated";

export const logout: NonNullable<MutationResolvers["logout"]> = async (_parent, _arg, _ctx) => {
  await _ctx.request.cookieStore?.delete("authorization");
  return true;
};
