import apiRequest from "./api";
import { CreateUpdateSchool, ListSchoolResponse, School } from "../types/school";

export async function getSchools(page: number, perPage: number): Promise<ListSchoolResponse> {
  return apiRequest<ListSchoolResponse>(`/schools?page=${page}&limit=${perPage}`, 'GET', undefined, true);
}

export async function createSchool(schoolData: CreateUpdateSchool): Promise<{ message: string; data: School }> {
  return apiRequest<{ message: string; data: School }>('/schools', 'POST', schoolData, true);
}

export async function updateSchool(id: number, schoolData: CreateUpdateSchool): Promise<{ message: string; data: School }> {
  return apiRequest<{ message: string; data: School }>(`/schools/${id}`, 'PUT', schoolData, true);
}

export async function deleteSchool(id: number): Promise<{ message: string }> {
  return apiRequest<{ message: string }>(`/schools/${id}`, 'DELETE', undefined, true);
}
