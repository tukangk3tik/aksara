import apiRequest from "./api";
import { CreateUpdateOffice, ListOfficeResponse, Office } from "../types/office";

export async function getOffices(page: number, perPage: number): Promise<ListOfficeResponse> {
  return apiRequest<ListOfficeResponse>(`/offices?page=${page}&perPage=${perPage}`, 'GET', undefined, true);
}

export async function createOffice(officeData: CreateUpdateOffice): Promise<{ message: string; data: Office }> {
  return apiRequest<{ message: string; data: Office }>('/offices', 'POST', officeData, true);
}

export async function updateOffice(id: number, officeData: CreateUpdateOffice): Promise<{ message: string; data: Office }> {
  return apiRequest<{ message: string; data: Office }>(`/offices/${id}`, 'PUT', officeData, true);
}

export async function deleteOffice(id: number): Promise<{ message: string }> {
  return apiRequest<{ message: string }>(`/offices/${id}`, 'DELETE', undefined, true);
}
