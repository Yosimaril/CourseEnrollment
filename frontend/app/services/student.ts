import type { Student, StudentCreate, StudentUpdate } from "~/models/student";

type StudentApiResponse = {
    id: number;
    username?: string;
    Username?: string;
    email: string;
    role: Student["role"];
    nrp?: string | null;
    max_credits?: number | null;
    maxCredits?: number | null;
    MaxCredits?: number | null;
    created_at?: string;
    createdAt?: string;
    updated_at?: string;
    updatedAt?: string;
};

function normalizeStudent(student: StudentApiResponse): Student {
    return {
        id: student.id,
        username: student.username ?? student.Username ?? "",
        email: student.email,
        role: student.role,
        nrp: student.nrp ?? undefined,
        maxCredits: student.maxCredits ?? student.max_credits ?? student.MaxCredits ?? undefined,
        createdAt: student.createdAt ?? student.created_at ?? "",
        updatedAt: student.updatedAt ?? student.updated_at ?? "",
    };
}

function normalizeStudents(students: StudentApiResponse[]): Student[] {
    return students.map((student) => normalizeStudent(student));
}

export const StudentService = {
    async getStudents(): Promise<Student[]> {
        const { $api } = useNuxtApp();
        const response = await $api<StudentApiResponse[]>("/admin/students?role=STUDENT");
        return normalizeStudents(response);
    },

    async getStudent(id: number): Promise<Student> {
        const { $api } = useNuxtApp();
        const response = await $api<StudentApiResponse>(`/admin/students/${id}`);
        return normalizeStudent(response);
    },

    async createStudent(student: StudentCreate): Promise<Student> {
        const { $api } = useNuxtApp();
        const response = await $api<StudentApiResponse>("/admin/students", {
            method: "POST",
            body: student,
        });
        return normalizeStudent(response);
    },

    async updateStudent(id: number, student: StudentUpdate): Promise<Student> {
        const { $api } = useNuxtApp();
        const response = await $api<StudentApiResponse>(`/admin/students/${id}`, {
            method: "PUT",
            body: student,
        });
        return normalizeStudent(response);
    },

    async deleteStudent(id: number): Promise<void> {
        const { $api } = useNuxtApp();
        return await $api(`/admin/students/${id}`, {
            method: "DELETE",
        });
    },
};
