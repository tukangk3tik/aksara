import { SelectOption, SelectOptionProps } from "./utils";

export interface Province extends SelectOption {}

export interface Regency extends SelectOption {}

export interface District extends SelectOption {}

// interface for select options
export interface LocationData {
  province: SelectOptionProps;
  regency: SelectOptionProps;
  district: SelectOptionProps;
}

export interface ListProvinceResponse {
  data: Province[];
  meta_data: {
    current_page: number;
    per_page: number;
    total_items: number;
  };
}

export interface ListRegencyResponse {
  data: Regency[];
  meta_data: {
    current_page: number;
    per_page: number;
    total_items: number;
  };
}

export interface ListDistrictResponse {
  data: District[];
  meta_data: {
    current_page: number;
    per_page: number;
    total_items: number;
  };
}