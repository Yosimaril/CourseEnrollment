import type { Student } from "~/models/student";
import type { StudentDto } from "~/dto/studentDto";
import { CoursePlanMapper } from "./coursePlanMapper";

export class StudentMapper {
    static toModel(dto: StudentDto): Student {
        return {
            id: dto.id,
            username: dto.username,
            email: dto.email,
            password: dto.password ?? undefined,
            nrp: dto.nrp,
            maxCredits: dto.max_credits,
            coursePlans: dto.course_plans ? dto.course_plans.map(CoursePlanMapper.toModel) : undefined,
            createdAt: new Date(dto.created_at),
            updatedAt: new Date(dto.updated_at),
            deletedAt: dto.deleted_at ? new Date(dto.deleted_at) : null,
        };
    }

    static toDto(model: Student): StudentDto {
        return {
            id: model.id,
            username: model.username,
            email: model.email,
            password: model.password ?? null,
            nrp: model.nrp,
            max_credits: model.maxCredits,
            course_plans: model.coursePlans ? model.coursePlans.map(CoursePlanMapper.toDto) : undefined,
            created_at: model.createdAt.toISOString(),
            updated_at: model.updatedAt.toISOString(),
            deleted_at: model.deletedAt ? model.deletedAt.toISOString() : null,
        };
    }
}
