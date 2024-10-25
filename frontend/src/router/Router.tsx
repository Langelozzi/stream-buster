import { BrowserRouter, Route, Routes } from 'react-router-dom';
import HomePage from '../pages/home/HomePage';
import LoginPage from '../pages/login/LoginPage';
import PrivateRoute from '../components/private-route/PrivateRoute';

const Router = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route
                    path="/"
                    element={
                        <PrivateRoute>
                            <HomePage />
                        </PrivateRoute>
                    }
                />
                <Route path='/login' Component={LoginPage} />
            </Routes>
        </BrowserRouter>
    )
}

export default Router;
