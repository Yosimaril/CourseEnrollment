import { z } from "zod";

export const BaseTimestampDto = z.object({
    created_at: z.string(),
    updated_at: z.string(),
    deleted_at: z.string().nullable().optional(),
});