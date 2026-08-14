import type { CoursePlan } from "~/models/coursePlan";
import type { CoursePlanDto } from "~/dto/coursePlanDto";
import { UserMapper } from "./userMapper";
import { CoursePlanItemMapper } from "./coursePlanItemMapper";

export class CoursePlanMapper {
    static toModel(dto: CoursePlanDto): CoursePlan {
        return {
            id: dto.id,
            studentId: dto.student_id,
            status: dto.status,
            student: dto.student ? UserMapper.toModel(dto.student) : undefined,
            items: dto.course_plan_items.map(CoursePlanItemMapper.toModel),
            createdAt: new Date(dto.created_at),
            updatedAt: new Date(dto.updated_at),
            deletedAt: dto.deleted_at ? new Date(dto.deleted_at) : null,
        };
    }

    static toDto(model: CoursePlan): CoursePlanDto {
        return {
            id: model.id,
            student_id: model.studentId,
            status: model.status,
            student: model.student ? UserMapper.toDto(model.student) : undefined,
            course_plan_items: model.items.map(CoursePlanItemMapper.toDto),
            created_at: model.createdAt.toISOString(),
            updated_at: model.updatedAt.toISOString(),
            deleted_at: model.deletedAt ? model.deletedAt.toISOString() : null,
        };
    }
}
