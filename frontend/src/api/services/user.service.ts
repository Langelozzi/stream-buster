import { User } from '../../models/user';
import axios from '../axios'

export const getCurrentUser = async (full: boolean = false) => {
    try {
        const response = await axios.get(`/user/current?full=${full}`);
        return response.data as User;
    } catch (error) {
        console.error('Error fetching TV details:', error);
        throw error;
    }
}