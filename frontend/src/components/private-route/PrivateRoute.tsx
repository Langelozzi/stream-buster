
import { Navigate } from "react-router-dom";
import React, { useState, useEffect } from "react";
import instance from "../../api/axios";

interface PrivateRouteProps {
	children: JSX.Element;
}

export const PrivateRoute: React.FC<PrivateRouteProps> = ({ children }) => {
	const [isAuthenticated, setIsAuthenticated] = useState<boolean | null>(null);

	useEffect(() => {
		const checkAuth = async () => {
			try {
				const response = await instance.get(`/user/current`);
				// Assume successful response means authenticated
				setIsAuthenticated(response.status !== 401);
			} catch (error) {
				// Any error means unauthenticated
				setIsAuthenticated(false);
			}
		};
		checkAuth();
	}, []);

	// Show a loading state while checking authentication
	if (isAuthenticated === null) {
		return <div>Loading...</div>;
	}

	// Redirect if not authenticated
	if (!isAuthenticated) {
		return <Navigate to="/login" replace />;
	}

	// Render the children if authenticated
	return children;
};
