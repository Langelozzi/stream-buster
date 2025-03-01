import { Media } from "../../models/media";
import { Movie } from "../../models/movie";
import { TV } from "../../models/tv";
import instance from "../axios";

export const getMediaByTMDBId = async (tmdbId: number): Promise<Media> => {
    try {
        const res = await instance.get(`/media/${tmdbId}`);
        return res.data;
    } catch (error) {
        console.error('Error getting media by TMDB id', error);
        throw error;
    }
}

export const createMedia = async (media: Media | Movie | TV): Promise<Media> => {
    try {
        const mediaResponse = await instance.post("/media/create", media);
        return mediaResponse.data
    } catch (error) {
        console.error('Error Creating Media', error);
        throw error;
    }
}

export const getMediaAvailability = async (media: Media): Promise<number> => {
    try {
        const res = await instance.post("/media/availability", media);
        return res.data.exists;
    } catch (error) {
        console.error('Error getting media availability', error);
        throw error;
    }
}
