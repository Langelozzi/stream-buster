import { createContext, useState, useEffect, ReactNode, useCallback } from 'react';
import { jwtDecode } from 'jwt-decode';
import { User } from '../models/user';
import { getCurrentUser } from '../api/services/user.service';
import { useNavigate } from 'react-router-dom';

export interface UserContextType {
    user: User | null;
    login: (user: User, token: string) => void;
    logout: () => void;
    validateToken: () => boolean;
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

    const navigate = useNavigate();

    // Helper function to get token from cookies
    const getTokenFromCookies = (name: string): string | undefined => {
        const cookies = document.cookie.split('; ');
        for (const cookie of cookies) {
            const [key, value] = cookie.split('=');
            if (key === name) {
                return decodeURIComponent(value);
            }
        }
    };

    // Helper function to decode JWT and check if it’s expired
    const decodeToken = (token: string): TokenClaims | null => {
        try {
            const decodedToken = jwtDecode<TokenClaims>(token);
            if (decodedToken.exp * 1000 < Date.now()) {
                return null;
            }
            return decodedToken;
        } catch (error) {
            console.error('Failed to decode JWT:', error);
            return null;
        }
    };

    // Login function to store the user and token
    const login = (newUser: User, token: string) => {
        setUser(newUser);
        localStorage.setItem('user', JSON.stringify(newUser));
        document.cookie = `token=${token}; path=/; secure; HttpOnly`;
    };

    // Logout function to clear the user and token
    const logout = () => {
        setUser(null);
        localStorage.removeItem('user');
        document.cookie = 'token=; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT';
        navigate('/login');
    };

    // Validate token and user state
    const validateToken = useCallback((): boolean => {
        const token = getTokenFromCookies('token');
        if (!token) {
            logout();
            return false;
        }

        const tokenClaims = decodeToken(token);
        if (!tokenClaims) {
            logout();
            return false;
        }

        return true;
    }, []);

    useEffect(() => {
        // On mount, validate the token and fetch the user if valid
        if (validateToken() && !user) {
            getCurrentUser(true)
                .then((fetchedUser) => {
                    setUser(fetchedUser);
                    localStorage.setItem('user', JSON.stringify(fetchedUser));
                })
                .catch((error) => {
                    console.error('Failed to fetch user:', error);
                    logout();
                });
        }
    }, [validateToken, user]);

    return (
        <UserContext.Provider value={{ user, login, logout, validateToken }}>
            {children}
        </UserContext.Provider>
    );
};
