import axios from "axios";
import { authApi } from "../api/api";
import type { LoginRequest, LoginResponse, RegisterRequest, RegisterResponse, CheckSessionResponse } from "../types/auth";


const handleError = (error: unknown): never => {
  if (axios.isAxiosError(error)) {
    const status = error.response?.status;
    switch (status) {
      case 400:
        throw new Error('Неверные данные. Проверьте введённые поля.');
      case 401:
        throw new Error('Неправильный логин или пароль.');
      case 409:
        throw new Error('Пользователь с таким e-mail уже зарегистрирован.');
      case 500:
        throw new Error('Ошибка сервера. Попробуйте позже.');
      default:
        throw new Error(`Неизвестная ошибка (${status}).`);
    }
  } else {
    throw error;
  }
};

export const authService = {
    login: async (req:  LoginRequest): Promise<LoginResponse> => {
        try {
            const res = await authApi.post<LoginResponse>('/login', req);
            const accessToken = res.data.access_token;
            const user = res.data.user;

            if (accessToken && user && user.email) {
                localStorage.setItem('access_token', accessToken);
                return res.data;
            } else {
                localStorage.removeItem('access_token');
                throw new Error('Failed to login');
            }
        } catch (error) {
            localStorage.removeItem('access_token');
            return handleError(error);
        }
    },

    register: async (req: RegisterRequest): Promise<RegisterResponse> => {
        try {
           const res = await authApi.post<RegisterResponse>('/register', req);
           const accessToken = res.data.access_token;
           const user = res.data.user;

            if (accessToken && user && user.email) {
                localStorage.setItem('access_token', accessToken);
                return res.data;
            } else {
                localStorage.removeItem('access_token');
                throw new Error('Failed to register');
            }
        } catch (error) {
            localStorage.removeItem('access_token');
            return handleError(error);
        }
    },

    logout: async (): Promise<{success: boolean}> => {
        try {
            localStorage.removeItem('access_token');
            await authApi.post('/logout');
            return {success: true};
        } catch (error) {
            return handleError(error);
        }
    },
    
    checkSession: async(): Promise<CheckSessionResponse> => {
        try {
            const res = await authApi.get<CheckSessionResponse>("/check-session", {validateStatus: () => true});
            return res.data;
        } catch (error) {
            return handleError(error);
        }
    }
}
