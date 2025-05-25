import  { createContext, useContext, useEffect, useState } from "react";
import type { ReactNode } from "react";
import type { User } from "../types/auth";
import { authService } from "../services/authService";

interface AuthContextType {
    user: User | null;
    loading: boolean;
    logout: () => Promise<void>;
    setUser: React.Dispatch<React.SetStateAction<User | null>>;
}


export const AuthContext = createContext<AuthContextType | undefined>(undefined);

export default function AuthProvider({ children }: { children: ReactNode}) {
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const checkSession = async () => {
            const res = await authService.checkSession();
            if (res.is_active && res.user) {
                setUser(res.user);
            }
            setLoading(false);
        }
        checkSession();
        console.log(user);
    }, []);

    const logout = async () => {
        await authService.logout();
        setUser(null);
    };

    return (
        <AuthContext.Provider value={{ user, loading, logout, setUser }}>
            {children}
        </AuthContext.Provider>
    );
}

export const useAuth = () => {
    const ctx = useContext(AuthContext);
    if (!ctx) throw new Error("useAuth must be used within an AuthProvider");
    return ctx;
};