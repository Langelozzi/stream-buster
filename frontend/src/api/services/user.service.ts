import { Usage } from '../../models/usage';
import { User } from '../../models/user';
import axios from '../axios'

export const getCurrentUser = async (full: boolean = false): Promise<User> => {
    try {
        const response = await axios.get(`/user/current?full=${full}`);
        return response.data as User;
    } catch (error) {
        console.error('Error fetching current user details:', error);
        throw error;
    }
}

export const getAllUsers = async (full: boolean = false): Promise<User[]> => {
    try {
        const response = await axios.get(`/user/?full=${full}`);
        return response.data as User[];
    } catch (error) {
        console.error('Error fetching all users:', error);
        throw error;
    }
}

export const getUserUsage = async (userId: number): Promise<Usage> => {
    try {
        const response = await axios.get(`/user/${userId}/usage`);
        return response.data as Usage;
    } catch (error) {
        console.error('Error fetching all users:', error);
        throw error;
    }
}