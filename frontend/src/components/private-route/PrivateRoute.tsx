import { Navigate } from "react-router-dom";
import { useUser } from "../../hooks/useUser";

interface PrivateRouteProps {
	children: JSX.Element;
}

export const PrivateRoute = ({ children }: PrivateRouteProps) => {
	const { user, loading } = useUser()

	if (loading) {
		// Render nothing or a loading indicator while waiting
		return <div>Loading...</div>;
	}

	return user ? children : <Navigate to="/login" replace />;
};
