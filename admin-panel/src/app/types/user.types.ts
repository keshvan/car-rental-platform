export interface User {
    id: number;
    email: string;
    role: 'user' | 'admin';
    balance?: number;
    created_at?: string;
}

export interface UsersResponse {
    users: User[];
}

export interface UpdateUserRolePayload {
    role: 'admin';
}