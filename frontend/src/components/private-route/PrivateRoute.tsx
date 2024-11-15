import { useUser } from "../../hooks/useUser";

interface PrivateRouteProps {
	children: JSX.Element;
}

export const PrivateRoute: React.FC<PrivateRouteProps> = ({ children }) => {
	const { validateToken, logout } = useUser();

	if (!validateToken()) {
		logout();
	}

	return children;
};