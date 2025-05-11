import { LoginResponse } from '../types/auth';
import { RequestHeaders } from '../types/header';

const API_BASE_URL = import.meta.env.VITE_API_URL || '/api';

// Helper function to get auth header
const getAuthHeader = () => {
  const token = localStorage.getItem('access_token');
  return token ? { Authorization: `Bearer ${token}` } : {};
};

// Generic API request function with authentication
async function apiRequest<T>(
  endpoint: string,
  method: string = 'GET',
  data?: any,
  requiresAuth: boolean = true
): Promise<T> {
  const url = `${API_BASE_URL}${endpoint}`;
  
  const headers: RequestHeaders = {
    'Content-Type': 'application/json',
    ...(requiresAuth ? getAuthHeader() : {})
  };

  const config: RequestInit = {
    method: method,
    headers: headers as HeadersInit,
    body: data ? JSON.stringify(data) : undefined,
  };

  const response = await fetch(url, config);

  if (!response.ok) {
    console.error('API request failed:', response);
    
    if (endpoint === '/auth/login' && response.status === 401) {
      throw new Error('Email atau password salah. Silahkan coba lagi.');
    }

    // Handle token expiration
    if (response.status === 401) {
      // Try to refresh token or logout
      // const refreshResult = await refreshToken();
      // if (refreshResult) {
        // Retry the original request
        // return apiRequest<T>(endpoint, method, data, requiresAuth);
      // } else {
        // If refresh fails, logout
        //logout();
        // throw new Error('Session expired. Please login again.');
      // }
    }
    
    const errorData = await response.json();
    // For login failures, provide a more user-friendly message
    throw new Error(errorData.message || 'API request failed');
  }

  const responseData = await response.json();
  return responseData.data;
}

// Authentication functions
export async function loginUser(email: string, password: string): Promise<LoginResponse> {
  return apiRequest<LoginResponse>('/auth/login', 'POST', { email, password }, false);
}

export async function refreshToken(): Promise<boolean> {
  try {
    const refreshToken = localStorage.getItem('refresh_token');
    if (!refreshToken) return false;

    const response = await fetch(`${API_BASE_URL}/refresh-token`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ refresh_token: refreshToken }),
    });

    if (!response.ok) return false;

    const data = await response.json();
    localStorage.setItem('access_token', data.data.access_token);
    return true;
  } catch (error) {
    return false;
  }
}

export function logout() {
  localStorage.removeItem('access_token');
  localStorage.removeItem('refresh_token');
  window.location.href = '/login';
}

// Export the generic request function for other API calls
export default apiRequest;
