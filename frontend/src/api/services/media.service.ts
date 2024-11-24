import { Media } from "../../models/media";
import { Movie } from "../../models/movie";
import { TV } from "../../models/tv";
import instance from "../axios";

export const getMedia = (id: number) => {
    try {
        const res = instance.get("/media/", { params: { id: id } })
        
        return res
    } catch (error) {
        return error
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
