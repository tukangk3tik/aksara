import { defineStore } from "pinia";
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL;

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem("access_token") || null,
    user: null,
  }),
  getters: {
    isAuthenticated: (state) => !!state.token,
  },
  actions: {
    async login(credentials) {
      try {
        const response = await fetch(`${API_BASE_URL}/users/login`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(credentials),
        });
        const payload = await response.json();

        if (response.ok) {
          const accessToken = payload.data.access_token;
          const refreshToken = payload.data.refresh_token;
          
          localStorage.setItem("access_token", accessToken);
          localStorage.setItem("refresh_token", refreshToken);
        } else {
          throw new Error(payload.message);
        }
      } catch (error) {
        console.log("Login error: " + error);
        throw error;
      }
    },
    async logout() {
      this.token = null;
      this.user = null;
      localStorage.removeItem("access_token");
      localStorage.removeItem("refresh_token");
    },
  },
});