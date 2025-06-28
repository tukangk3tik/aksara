import apiRequest from "./api";
import { CreateUpdateOffice, ListOfficeResponse, Office, SelectOptionOfficeResponse } from "../types/office";

export async function getOffices(page: number, perPage: number): Promise<ListOfficeResponse> {
  return apiRequest<ListOfficeResponse>(`/offices?page=${page}&limit=${perPage}`, 'GET', undefined, true);
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

export async function fetchOfficesSelectOption(searchQuery: string = ''): Promise<SelectOptionOfficeResponse> {
  return apiRequest<SelectOptionOfficeResponse>(`/offices/select-option?search_query=${searchQuery}`, 'GET', undefined, true);
}