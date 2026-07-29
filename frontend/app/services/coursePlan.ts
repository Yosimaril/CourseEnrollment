import type { CoursePlanItemStatusEnum } from "~/constants/coursePlanItemStatusEnum";
import type { CoursePlan } from "~/models/coursePlan";
import type { CoursePlanItem } from "~/models/coursePlanItem";
import type { Course } from "~/models/course";
import type { User } from "~/models/user";
import type { RawCourse } from "~/mapper/courseMapper";
import type { RawCoursePlan } from "~/mapper/coursePlanMapper";
import type { RawCoursePlanItem } from "~/mapper/coursePlanItemMapper";
import type { RawUser } from "~/mapper/userMapper";

function normalizeUser(user?: RawUser): User | undefined {
    if (!user) {
        return undefined;
    }

    return {
        id: user.id ?? 0,
        username: user.username ?? user.Username ?? "",
        email: user.email ?? "",
        role: (user.role ?? "STUDENT") as User["role"],
        nrp: user.nrp,
        maxCredits: user.maxCredits ?? user.max_credits,
        createdAt: user.createdAt ?? user.created_at ?? "",
        updatedAt: user.updatedAt ?? user.updated_at ?? "",
    };
}

function normalizeCourse(course?: RawCourse): Course | undefined {
    if (!course) {
        return undefined;
    }

    return {
        id: course.id ?? 0,
        code: course.code ?? course.Code ?? "",
        name: course.name ?? course.Name ?? "",
        credits: course.credits ?? course.Credits ?? 0,
        createdAt: course.createdAt ?? course.created_at ?? "",
        updatedAt: course.updatedAt ?? course.updated_at ?? "",
    };
}

function normalizeCoursePlanItem(item: RawCoursePlanItem): CoursePlanItem {
    return {
        coursePlanId: item.coursePlanId ?? item.course_plan_id ?? item.CoursePlanID ?? 0,
        courseId: item.courseId ?? item.course_id ?? item.CourseID ?? 0,
        status: (item.status ?? "PENDING") as CoursePlanItemStatusEnum,
        remarks: item.remarks ?? undefined,
        course: normalizeCourse(item.course ?? item.Course),
        createdAt: item.createdAt ?? item.created_at ?? "",
        updatedAt: item.updatedAt ?? item.updated_at ?? "",
    };
}

function normalizeCoursePlan(plan: RawCoursePlan): CoursePlan {
    const items = plan.items ?? plan.Items ?? [];

    return {
        id: plan.id ?? 0,
        studentId: plan.studentId ?? plan.student_id ?? 0,
        status: (plan.status ?? "DRAFT") as CoursePlan["status"],
        student: normalizeUser(plan.student ?? plan.Student),
        items: items.map(normalizeCoursePlanItem),
        createdAt: plan.createdAt ?? plan.created_at ?? "",
        updatedAt: plan.updatedAt ?? plan.updated_at ?? "",
    };
}

export const CoursePlanService = {
    async getMyCoursePlan(): Promise<CoursePlan> {
        const { $api } = useNuxtApp();
        const response = await $api<RawCoursePlan>("/student/course-plan");
        return normalizeCoursePlan(response);
    },

    async getMyCoursePlans(): Promise<CoursePlan[]> {
        const { $api } = useNuxtApp();
        const response = await $api<RawCoursePlan[]>("/student/course-plans");
        return Array.isArray(response) ? response.map(normalizeCoursePlan) : [];
    },

    async submitMyCoursePlan(): Promise<CoursePlan> {
        const { $api } = useNuxtApp();
        const response = await $api<RawCoursePlan>("/student/course-plan/submit", {
            method: "POST",
        });

        return normalizeCoursePlan(response);
    },

    async cancelMyCoursePlan(id: number) {
        const { $api } = useNuxtApp();
        return await $api(`/student/course-plans/${id}`, {
            method: "DELETE",
        });
    },

    async getPendingCoursePlans(): Promise<CoursePlan[]> {
        const { $api } = useNuxtApp();
        const response = await $api<RawCoursePlan[]>("/admin/course-plans?status=PENDING_REVIEW");
        return Array.isArray(response) ? response.map(normalizeCoursePlan) : [];
    },

    async reviewCoursePlan(
        id: number,
        payload: {
            item_status: CoursePlanItemStatusEnum;
            course_ids?: number[];
        },
    ) {
        const { $api } = useNuxtApp();
        const response = await $api<RawCoursePlan>(`/admin/course-plans/${id}/review`, {
            method: "PUT",
            body: payload,
        });

        return normalizeCoursePlan(response);
    },
};
