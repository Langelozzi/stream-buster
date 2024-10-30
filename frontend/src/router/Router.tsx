import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { HomePage } from '../pages/home/HomePage';
import { LoginPage } from '../pages/login/LoginPage';
import { BrowsePage } from '../pages/browse/BrowsePage';
import { WatchPage } from '../pages/watch/WatchPage';
import { PrivateRoute } from '../components/private-route/PrivateRoute';
import { DashboardPage } from '../pages/dashboard/DashboardPage';

export const Router = () => {
    return (
        <BrowserRouter>
            <Routes>
                {/* Public Routes */}
                <Route path='/login' Component={LoginPage} />
                <Route path='/' Component={HomePage} />

                {/* Private Routes */}
                <Route
                    path="/dashboard"
                    element={
                        <PrivateRoute>
                            <DashboardPage />
                        </PrivateRoute>
                    }
                />
                <Route
                    path='/browse'
                    element={
                        <PrivateRoute>
                            <BrowsePage />
                        </PrivateRoute>
                    }
                />
                <Route
                    path='/watch/:tmdbId/:seasonNum?/:episodeNum?'
                    element={
                        <PrivateRoute>
                            <WatchPage />
                        </PrivateRoute>
                    }
                />
            </Routes>
        </BrowserRouter>
    )
}
