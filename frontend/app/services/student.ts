import type { Student, StudentCreate, StudentUpdate } from "~/models/student";

export const StudentService = {
    async getStudents(): Promise<Student[]> {
        const { $api } = useNuxtApp();
        return await $api("/admin/students?role=STUDENT");
    },

    async getStudent(id: number): Promise<Student> {
        const { $api } = useNuxtApp();
        return await $api(`/admin/students/${id}`);
    },

    async createStudent(student: StudentCreate): Promise<Student> {
        const { $api } = useNuxtApp();
        return await $api("/admin/students", {
            method: "POST",
            body: student,
        });
    },

    async updateStudent(id: number, student: StudentUpdate): Promise<Student> {
        const { $api } = useNuxtApp();
        return await $api(`/admin/students/${id}`, {
            method: "PUT",
            body: student,
        });
    },

    async deleteStudent(id: number): Promise<void> {
        const { $api } = useNuxtApp();
        return await $api(`/admin/students/${id}`, {
            method: "DELETE",
        });
    },
};
