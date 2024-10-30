import axios from '../axios';
import { TV } from '../../models/tv';
import { Episode } from '../../models/episode';

export const getTVDetails = async (id: number): Promise<TV> => {
    try {
        const response = await axios.get(`/tv/${id}`);
        return response.data as TV;
    } catch (error) {
        console.error('Error fetching TV details:', error);
        throw error;
    }
}

export const getEpisodesForSeason = async (seriesId: number, seasonNum: number): Promise<Episode[]> => {
    try {
        const response = await axios.get(`/tv/${seriesId}/season/${seasonNum}/episodes`);
        return response.data as Episode[];
    } catch (error) {
        console.error('Error fetching episodes:', error);
        throw error;
    }
}