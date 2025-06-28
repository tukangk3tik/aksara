import { MetaData } from "./pagination";

export interface School {
  id: number;
  code: string;
  name: string;
  office_id: number;
  office: string;
  is_public_school: boolean;
  province_id: number;
  regency_id: number;
  district_id: number;
  province: string;
  regency: string;
  district: string;
  email: string;
  phone: string;
  address: string;
  logo_url: string;
}

export interface CreateUpdateSchool {
  id?: number;
  code: string;
  name: string;
  province_id: number;
  regency_id: number;
  district_id: number;
  office_id: number;
  is_public_school: boolean;
  email: string;
  phone: string;
  address: string;
  logo_url: string;
}

export interface ListSchoolResponse {
  message: string;
  data: School[];
  meta_data: MetaData;
}
