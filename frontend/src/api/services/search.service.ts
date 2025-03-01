import axios from '../axios';
import { Movie } from '../../models/movie';
import { TV } from '../../models/tv';
import { SearchPage } from '../../models/search_page';

export const castToTvOrMovie = (item: any): (TV | Movie) => {
    if (item.Media?.MediaType.Name === 'tv') {
        return item as TV;
    } else if (item?.Media?.MediaType.Name === 'movie') {
        return item as Movie;
    } else {
        return item;
    }
}

export const searchMulti = async (query: string, page: number): Promise<SearchPage> => {
    try {
        const response = await axios.get(`/search/multi`, {
            params: {
                query: query,
                page: page
            },
        });

        return response.data as SearchPage;
    } catch (error) {
        console.error('Error fetching search results:', error);
        throw error;
    }
};
