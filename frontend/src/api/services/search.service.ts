import axios from '../axios';
import { Movie } from '../../models/movie';
import { TV } from '../../models/tv';

export const castToTvOrMovie = (item: any): (TV | Movie) => {
    if (item.Media?.MediaType.Name === 'tv') {
        return item as TV;
    } else if (item?.Media?.MediaType.Name === 'movie') {
        return item as Movie;
    } else {
        return item;
    }
}

export const searchMulti = async (query: string): Promise<(TV | Movie)[]> => {
    try {
        const response = await axios.get(`/search/multi`, {
            params: {
                query: query,
            },
        });

        const data = response.data as any[];
        if (!data || data.length === 0) {
            return [];
        }

        return data.map(item => {
            return castToTvOrMovie(item)
        })
    } catch (error) {
        console.error('Error fetching search results:', error);
        throw error;
    }
};
