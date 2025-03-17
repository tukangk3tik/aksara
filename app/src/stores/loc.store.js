import { defineStore } from 'pinia'
import { authorizedReq } from "@/utils/authorizedReq";

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL;

export const useLocStore = defineStore('loc', {
  state: () => ({
    provinces: [],
    regencies: [],
    districts: [],
    selectedProvince: null,
    selectedRegency: null,
    selectedDistrict: null,
    searchQueryProvince: '',
    searchQueryRegency: '',
    searchQueryDistrict: '',
    isLoading: true,
  }),
  actions: {
    async fetchProvinces() {
      this.isLoading = true;
      
      setTimeout(async () => {
        try {
          const response = await authorizedReq(`${API_BASE_URL}/loc/provinces?search_query=${this.searchQueryProvince}`);
          const payload = await response.json();
          if (response.ok) {
            this.provinces = payload.data;
          } else {
            throw new Error(payload.message);
          }
        } catch (error) {
          console.log("Fetch provinces error: " + error);
          throw error;
        } finally {
          this.isLoading = false;
        }
      }, 500);
    },
    async fetchRegenciesByProvince() {
      this.isLoading = true;

      setTimeout(async () => {
        try {
          const response = await authorizedReq(`${API_BASE_URL}/loc/regencies?province_id=${this.selectedProvince}&search_query=${this.searchQueryRegency}`);
          const payload = await response.json();
          if (response.ok) {
            this.regencies = payload.data;
          } else {
            throw new Error(payload.message);
          }
        } catch (error) {
          console.log("Fetch regencies error: " + error);
          throw error;
        } finally {
          this.isLoading = false;
        }
      }, 500);
    },
    async fetchDistrictsByRegency() {
      this.isLoading = true;
      
      setTimeout(async () => {
        try {
          const response = await authorizedReq(`${API_BASE_URL}/loc/districts?regency_id=${this.selectedRegency}&search_query=${this.searchQueryDistrict}`);
          const payload = await response.json();
          if (response.ok) {
            this.districts = payload.data;
          } else {
            throw new Error(payload.message);
          }
        } catch (error) {
          console.log("Fetch districts error: " + error);
          throw error;
        } finally {
          this.isLoading = false;
        }
      }, 500);
    },
  },
})
