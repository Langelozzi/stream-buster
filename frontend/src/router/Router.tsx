import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { HomePage } from '../pages/home/HomePage';
import { LoginPage } from '../pages/login/LoginPage';
import { BrowsePage } from '../pages/browse/BrowsePage';
import { WatchPage } from '../pages/watch/WatchPage';

export const Router = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path='/' Component={HomePage} />
                <Route path='/login' Component={LoginPage} />
                <Route path='/browse' Component={BrowsePage} />
                <Route path='/watch' Component={WatchPage} />
            </Routes>
        </BrowserRouter>
    )
}