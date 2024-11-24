import { createContext, useState, useEffect, ReactNode } from 'react';
import { User } from '../models/user';
import { getCurrentUser } from '../api/services/user.service';
import { useNavigate } from 'react-router-dom';
import { routes } from '../router/Routes';
import { logoutUser } from '../api/services/auth.service';

export interface UserContextType {
    user: User | null;
    login: (user: User, token: string) => void;
    logout: () => void;
}

export const UserContext = createContext<UserContextType | undefined>(undefined);

export const UserProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const [user, setUser] = useState<User | null>(() => {
        const storedUser = localStorage.getItem('user');
        return storedUser ? JSON.parse(storedUser) : null;
    });

    const navigate = useNavigate();

    // Login function to store the user and token
    const login = (newUser: User) => {
        setUser(newUser);
        localStorage.setItem('user', JSON.stringify(newUser));
    };

    // Logout function to clear the user and token
    const logout = () => {
        setUser(null);
        logoutUser()
        localStorage.removeItem('user');
        navigate(routes.login);
    };
    useEffect(() => {
        // On mount, validate the token and fetch the user if valid
        getCurrentUser(true)
            .then((fetchedUser) => {
                setUser(fetchedUser);
                localStorage.setItem('user', JSON.stringify(fetchedUser));
            })
            .catch((error) => {
                console.error('Failed to fetch user:', error);
            });
    }, []);

    return (
        <UserContext.Provider value={{ user, login, logout }}>
            {children}
        </UserContext.Provider>
    );
};
