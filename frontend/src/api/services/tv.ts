import axios from '../axios';
import { TV } from '../../models/tv';

export const getTVDetails = async (id: number): Promise<TV> => {
    try {
        const response = await axios.get(`/tv/${id}`);
        return response.data as TV;
    } catch (error) {
        console.error('Error fetching TV details:', error);
        throw error;
    }
}