import { MetaData } from "./pagination";

export interface Office {
  id: number;
  code: string;
  name: string;
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

export interface CreateUpdateOffice {
  id?: number;
  code: string;
  name: string;
  province_id: number;
  regency_id: number;
  district_id: number;
  email: string;
  phone: string;
  address: string;
  logo_url: string;
}

export interface ListOfficeResponse {
  message: string;
  data: Office[];
  meta_data: MetaData;
}
