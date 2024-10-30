import axios from '../axios';
import { Movie } from '../../models/movie';

export const getMovieDetails = async (id: number): Promise<Movie> => {
    try {
        const response = await axios.get(`/movie/${id}`);
        return response.data as Movie;
    } catch (error) {
        console.error('Error fetching Movie details:', error);
        throw error;
    }
}