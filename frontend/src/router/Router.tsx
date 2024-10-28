import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { HomePage } from '../pages/home/HomePage';
import { LoginPage } from '../pages/login/LoginPage';
import { BrowsePage } from '../pages/browse/BrowsePage';
import { WatchPage } from '../pages/watch/WatchPage';
import PrivateRoute from '../components/private-route/PrivateRoute';

export const Router = () => {
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
                <Route path='/browse' Component={BrowsePage} />
                <Route path='/watch' Component={WatchPage} />
            </Routes>
        </BrowserRouter>
    )
}
