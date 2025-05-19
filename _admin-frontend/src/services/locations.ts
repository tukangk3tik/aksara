import { ListProvinceResponse } from "../types/location";
import apiRequest from "./api";

export async function getProvinces(searchQuery: string = '', page: number = 1, perPage: number = 100): Promise<ListProvinceResponse> {
  return apiRequest<ListProvinceResponse>(`/loc/provinces?page=${page}&perPage=${perPage}&search_query=${searchQuery}`, 'GET', undefined, true);
}

export async function getRegenciesByProvince(provinceId: number, searchQuery: string = '', page: number = 1, perPage: number = 100): Promise<any> {
  return apiRequest<any>(`/loc/regencies?page=${page}&perPage=${perPage}&province_id=${provinceId}&search_query=${searchQuery}`, 'GET', undefined, true);
}

export async function getDistrictsByRegency(regencyId: number, searchQuery: string = '', page: number = 1, perPage: number = 100): Promise<any> {
  return apiRequest<any>(`/loc/districts?page=${page}&perPage=${perPage}&regency_id=${regencyId}&search_query=${searchQuery}`, 'GET', undefined, true);
}
