import { useContext } from "react";
import { UserContext } from "../contexts/UserContext";

export const useUser = () => {
    const context = useContext(UserContext);
    if (!context) {
        console.error('useUser hook must be within the UserProvider');
    }
    return context?.user;
}