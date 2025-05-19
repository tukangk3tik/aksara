import { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { loginUser as apiLoginUser, logout as apiLogout } from '../services/api';
import { AuthState, LoginResponse } from '../types/auth';

interface AuthContextType extends AuthState {
  login: (email: string, password: string) => Promise<void>;
  logout: () => void;
  loading: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}

interface AuthProviderProps {
  children: ReactNode;
}

export function AuthProvider({ children }: AuthProviderProps) {
  const [authState, setAuthState] = useState<AuthState>({
    isAuthenticated: false,
    user: null,
    accessToken: null,
    refreshToken: null,
  });
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    // Check if user is already logged in
    const storedAccessToken = localStorage.getItem('access_token');
    const storedRefreshToken = localStorage.getItem('refresh_token');
    
    if (storedAccessToken && storedRefreshToken) {
      setAuthState({
        isAuthenticated: true,
        user: null, // We'll need to fetch user info in a real app
        accessToken: storedAccessToken,
        refreshToken: storedRefreshToken,
      });
    }
    
    setLoading(false);
  }, []);

  const login = async (email: string, password: string) => {
    setLoading(true);

    try {
      const response = await apiLoginUser<LoginResponse>(email, password);
        
      // Store tokens
      localStorage.setItem('access_token', response.data.access_token);
      localStorage.setItem('refresh_token', response.data.refresh_token);
        
      // In a real app, you would fetch user info here
      // const userInfo = await fetchUserInfo();
        
      setAuthState({
        isAuthenticated: true,
        user: null, // Replace with userInfo in a real app
        accessToken: response.data.access_token,
        refreshToken: response.data.refresh_token,
      });
    } catch (error) {
      throw error;
    } finally {
      setLoading(false);
    }
  };
    
  const logout = () => {
    apiLogout();
    setAuthState({
      isAuthenticated: false,
      user: null,
      accessToken: null,
      refreshToken: null,
    });
  };

  const value = {
    ...authState,
    login,
    logout,
    loading
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}
