import type { BaseIdDto } from "~/dto/baseIdDto";
import type { BaseTimestampDto } from "~/dto/baseTimestampDto";
import type { CoursePlanItemDto } from "../dto/coursePlanItemDto";
import type { StudentDto } from "../dto/studentDto";

export type RawCoursePlan = BaseIdDto &
    BaseTimestampDto & {
        id: number;
        student_id: number;
        status: string;
        items: CoursePlanItemDto[];
        student: StudentDto;
    };

export class CoursePlanMapper {
    static fromJson(json: unknown): RawCoursePlan {
        return json as RawCoursePlan;
    }

    static toJson(dto: RawCoursePlan): object {
        return {
            id: dto.id,
            student_id: dto.student_id,
            status: dto.status,
            student: dto.student,
            items: dto.items,
            created_at: dto.created_at,
            updated_at: dto.updated_at,
            deleted_at: dto.deleted_at,
        };
    }
}