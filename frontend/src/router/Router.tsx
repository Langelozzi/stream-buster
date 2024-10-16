import { BrowserRouter, Route, Routes } from 'react-router-dom';
import HomePage from '../pages/home/HomePage';
import LoginPage from '../pages/login/LoginPage';

const Router = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path='/' Component={HomePage}/>
                <Route path='/login' Component={LoginPage} />
            </Routes>
        </BrowserRouter>
    )
}

export default Router;