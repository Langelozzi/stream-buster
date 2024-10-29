import { useContext } from "react";
import { UserContext, UserContextType } from "../contexts/UserContext";


export const useUser = (): UserContextType => {
    const context = useContext(UserContext);
    if (!context) {
        console.error('useUser hook must be within the UserProvider');
    }
    return context!;
}
