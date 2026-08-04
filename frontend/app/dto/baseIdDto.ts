import { z } from "zod";

export const BaseIdDto = z.object({
    id: z.number()
});
