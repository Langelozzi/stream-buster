import { createContext, useState, useEffect, ReactNode, useMemo } from 'react';
import { jwtDecode } from 'jwt-decode';
import { User } from '../models/user';
import { getCurrentUser } from '../api/services/user.service';

export interface UserContextType {
    user: User | null;
    loading: boolean;
}

interface TokenClaims {
    id: number;
    email: string;
    firstName: string;
    lastName: string;
    iss: string;
    exp: number;
    iat: number;
}

export const UserContext = createContext<UserContextType | undefined>(undefined);

export const UserProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const [user, setUser] = useState<User | null>(() => {
        const storedUser = localStorage.getItem('user');
        return storedUser ? JSON.parse(storedUser) : null;
    });
    const [loading, setLoading] = useState(true);

    const getTokenFromCookies = (name: string): string | undefined => {
        const cookies = document.cookie.split('; ');
        for (const cookie of cookies) {
            const [key, value] = cookie.split('=');
            if (key === name) {
                return decodeURIComponent(value);
            }
        }
    };

    const decodeToken = (token: string): TokenClaims => {
        const decodedToken: TokenClaims = jwtDecode<TokenClaims>(token);
        return decodedToken;
    };

    useEffect(() => {
        const fetchCurrentUserFull = async () => {
            try {
                const fetchedUser: User = await getCurrentUser(true);
                setUser(fetchedUser);
                localStorage.setItem('user', JSON.stringify(fetchedUser)); // Store in localStorage
            } catch (error) {
                console.error('Error fetching current user:', error);
                setUser(null);
                localStorage.removeItem('user'); // Remove if fetch fails
            } finally {
                setLoading(false);
            }
        };

        const token = getTokenFromCookies("token");

        if (token) {
            try {
                // const tokenClaims = decodeToken(token);
                if (!user) {
                    fetchCurrentUserFull();
                }
            } catch (e) {
                console.error("Invalid JWT", e);
            }
        } else {
            setUser(null);
            localStorage.removeItem('user');
        }
        setLoading(false);
    }, []);

    // Memoize context value to avoid unnecessary re-renders
    const value = useMemo(() => ({ user, loading }), [user, loading]);

    return (
        <UserContext.Provider value={value}>
            {children}
        </UserContext.Provider>
    );
};
