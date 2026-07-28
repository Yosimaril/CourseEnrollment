import type { Student, StudentCreate, StudentUpdate } from "~/models/student"

const { $api } = useNuxtApp()

export const StudentService = {
  async getStudents(): Promise<Student[]> {
    return await $api("/admin/students?role=STUDENT")
  },

  async getStudent(id: number): Promise<Student> {
    return await $api(`/admin/students/${id}`)
  },

  async createStudent(student: StudentCreate): Promise<Student> {
    return await $api("/admin/students", {
      method: "POST",
      body: student,
    })
  },

  async updateStudent(id: number, student: StudentUpdate): Promise<Student> {
    return await $api(`/admin/students/${id}`, {
      method: "PUT",
      body: student,
    })
  },

  async deleteStudent(id: number): Promise<void> {
    return await $api(`/admin/students/${id}`, {
      method: "DELETE",
    })
  },
}
