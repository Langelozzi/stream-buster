import axios from 'axios';
import { Movie } from '../models/movie';
import { TV } from '../models/tv';

const BASE_URL = import.meta.env.VITE_API_BASE_URL;

export const searchMulti = async (query: string): Promise<(TV | Movie)[]> => {
    try {
        const response = await axios.get(`${BASE_URL}/search/multi`, {
            params: {
                query: query,
            },
        });

        const data = response.data as any[];

        return data.map(item => {
            if (item.Media.MediaType.Name === 'tv') {
                return {
                    mediaID: item.MediaID,
                    media: {
                        id: item.Media.ID,
                        tmdbID: item.Media.TMDBID,
                        title: item.Media.Title,
                        posterImage: item.Media.PosterImage,
                        mediaTypeId: item.Media.MediaTypeId,
                        mediaType: item.Media.MediaType,
                        deletedAt: item.Media.DeletedAt ? new Date(item.Media.DeletedAt) : undefined,
                        createdAt: item.Media.CreatedAt ? new Date(item.Media.CreatedAt) : undefined,
                    },
                    overview: item.Overview,
                    seasonCount: item.SeasonCount,
                    episodeCount: item.EpisodeCount,
                    seasons: item.Seasons,
                    firstAirDate: item.FirstAirDate ? new Date(item.FirstAirDate) : undefined
                } as TV;
            } else if (item.Media.MediaType.Name === 'movie') {
                return {
                    mediaID: item.MediaID,
                    media: {
                        id: item.Media.ID,
                        tmdbID: item.Media.TMDBID,
                        title: item.Media.Title,
                        posterImage: item.Media.PosterImage,
                        mediaTypeId: item.Media.MediaTypeId,
                        mediaType: item.Media.MediaType,
                        deletedAt: item.Media.DeletedAt ? new Date(item.Media.DeletedAt) : undefined,
                        createdAt: item.Media.CreatedAt ? new Date(item.Media.CreatedAt) : undefined,
                    },
                    overview: item.Overview,
                    posterPath: item.PosterPath,
                    genres: item.Genres,
                    releaseDate: item.ReleaseDate ? new Date(item.ReleaseDate) : undefined,
                    runtime: item.Runtime
                } as Movie;
            } else {
                return item;
            }
        });
    } catch (error) {
        console.error('Error fetching search results:', error);
        throw error;
    }
};