import { defineStore } from "pinia";
import { DataTableDto } from "@/utils/dtos/datatableDto";
import { authorizedReq } from "@/utils/authorizedReq";
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL;

export const useOfficeStore = defineStore('office', {
  state: () => ({   
    offices: [],
    totalItems: 0,
    limit: 10,
    page: 1,
    sort: '',
    search: '',
    isLoading: true,
  }),
  actions: {
    async fetchOffices() {
      this.isLoading = true;
      
      setTimeout(async () => {
        this.isLoading = false;
        try {
          const dataTableDto = new DataTableDto({page: this.page, limit: this.limit, sort: this.sort, search: this.search});
          const response = await authorizedReq(`${API_BASE_URL}/offices?${dataTableDto.toQueryParams()}`);
          const payload = await response.json();
          if (response.ok) {
            this.offices = payload.data;
            this.totalItems = payload.meta_data.total_items;
          } else {
            throw new Error(payload.message);
          }
        } catch (error) {
          console.log("Fetch offices error: " + error);
          throw error;
        } finally {
          this.isLoading = false;
        }
      }, 500);
    },
  },
});