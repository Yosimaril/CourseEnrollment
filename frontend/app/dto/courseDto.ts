import { z } from "zod";
import { BaseIdDto } from "./baseIdDto.ts";
import { BaseTimestampDto } from "./baseTimestampDto.ts";

export const CourseDto = BaseIdDto
    .extend(BaseTimestampDto)
    .extend({
        code: z.string(),
        name: z.string(),
        credits: z.number(),
    });

export type CourseDtoType = z.infer<typeof CourseDto>;