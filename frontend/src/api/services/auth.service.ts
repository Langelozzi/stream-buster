import { AxiosResponse } from 'axios';
import axios from '../axios';
import instance from '../axios';

export const postLogin = async (form: FormData) => {
    try {
        const res: AxiosResponse = await axios.post("/auth/login", form);
        return res;
    } catch (err: any) {
        if (err.status == 401) {
            alert("Wrong username or password. Try again or click Forgot password to reset it.");
        }
    }
}

export const postRegister = async (form: FormData) => {
    try {
        const res: AxiosResponse = await axios.post("/auth/register", form);
        return res;
    } catch (err: any) {
        if (err.status == 401) {
            alert("Registration failed");
        }
    }
}

export const sendTestRequest = async () => {
    const res: AxiosResponse = await axios.get("/auth/test");
    return res;
}

export const logoutUser = async () => {
    const res = await instance.get("/auth/logout");
    return res;
}
