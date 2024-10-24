import { createContext, useState, useEffect, ReactNode } from 'react';
import { jwtDecode } from 'jwt-decode'
import { User } from '../models/user'

interface UserContextType {
    user: User | null;
}

interface TokenClaims {
    user: User;
    iss: string;
    exp: number;
    iat: number;
}

export const UserContext = createContext<UserContextType | undefined>(undefined);

export const UserProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const [user, setUser] = useState<User | null>(null);

    const getTokenFromCookies = (): string => {
        const token: string = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjk3NDMxMTAsImlhdCI6MTcyOTczOTUxMCwiaXNzIjoiYXV0aC1zZXJ2aWNlIiwic3ViIjoidXNlcm5hbWUifQ.9VDmUhEQeGOQshe5p7uZh5vvgn97_9j6U-GKt2OeQX4'
        return token;
    }

    const decodeToken = (token: string): TokenClaims => {
        const decodedToken: TokenClaims = jwtDecode<TokenClaims>(token);
        return decodedToken;
    }

    useEffect(() => {
        const token = getTokenFromCookies();

        if (token) {
            try {
                const tokenClaims = decodeToken(token);
                setUser(tokenClaims.user);
            } catch (e) {
                console.log("Invalid jwt", e);
            }
        }
    }, [])

    return (
        <UserContext.Provider value={{ user }}>
            {children}
        </UserContext.Provider>
    )
}