import { Navigate } from "react-router-dom";
import { useUser } from "../../hooks/useUser";

interface PrivateRouteProps {
	children: JSX.Element;
}

export const PrivateRoute: React.FC<PrivateRouteProps> = ({ children }) => {
	const { validateToken } = useUser();

	if (!validateToken()) {
		return <Navigate to="/login" />;
	}

	return children;
};