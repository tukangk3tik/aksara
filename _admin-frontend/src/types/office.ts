import { District, Province, Regency } from "./location";
import { MetaData } from "./pagination";
import { SelectOption } from "./utils";

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

export interface SelectOptionOffice extends SelectOption {
  additional_data: {
    province: Province;
    regency: Regency;
    district: District;
  };
}

export interface SelectOptionOfficeResponse {
  message?: string;
  data: SelectOptionOffice[];
}
