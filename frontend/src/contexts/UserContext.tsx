
import { createContext, useState, useEffect, ReactNode } from 'react';
import { jwtDecode } from 'jwt-decode';
import { User } from '../models/user';

interface UserContextType {
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
    const [user, setUser] = useState<User | null>(null);
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
        const token = getTokenFromCookies("token");

        if (token) {
            try {
                const tokenClaims = decodeToken(token);
                setUser({
                    email: tokenClaims.email,
                    firstName: tokenClaims.firstName,
                    lastName: tokenClaims.lastName,
                    id: tokenClaims.id,
                });
            } catch (e) {
                console.error("Invalid JWT", e);
            }
        }

        // Set loading to false once token is processed
        setLoading(false);
    }, []);

    return (
        <UserContext.Provider value={{ user, loading }}>
            {children}
        </UserContext.Provider>
    );
};
