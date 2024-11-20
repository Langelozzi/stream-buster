import { CurrentlyWatching } from "../../models/currently_watching";
import { Episode } from "../../models/episode";
import { createMedia } from "./media.service";
import instance from "../axios";
import { UserContextType } from "../../contexts/UserContext";
import { Movie } from "../../models/movie";
import { TV } from "../../models/tv";

export const getCurrentlyWatching = async () => {
    try {
        const result = await instance.get("/currently-watching/getall");

        return result.data;
    } catch (error) {
        console.error('Error fetching currently watching:', error);
        throw error;
    }
}

export const createCurrentlyWatching = async (data: CurrentlyWatching) => {
    try {
        const result = await instance.post("/currently-watching/", data);
        return result.data;
    } catch (error) {
        console.error('Error creating currently watching:', error);
        throw error;
    }
}

export const updateCurrentlyWatching = async (data: CurrentlyWatching) => {
    try {
        const result = await instance.put("/currently-watching/update", data);
        return result.data;
    } catch (error) {
        console.error('Error creating currently watching:', error);
        throw error;
    }
}

export const getWatchList = async (): Promise<CurrentlyWatching[]> => {
    try {
        const result = await instance.get("/currently-watching/watchlist");
        return result.data;
    } catch (error) {
        console.error('Error fetching watch list:', error);
        throw error;
    }
}
export const onAddToList = async (media: Movie | TV, user: UserContextType, currentEpisode: Episode | null = null) => {
    try {
        let mediaResponse;
        try {
            mediaResponse = await createMedia(media.Media!)
            console.log(mediaResponse)
        } catch (error) {
            console.error(error);

        }

        const currentlyWatching: CurrentlyWatching = {
            MediaId: mediaResponse?.ID,
            UserID: user.user?.ID,
            SeasonNumber: currentEpisode ? currentEpisode?.SeasonNumber : 0,
            EpisodeNumber: currentEpisode ? currentEpisode?.EpisodeNumber : 0,
        }

        await createCurrentlyWatching(currentlyWatching)

    } catch (error) {
        console.error("Error addign to list")
    }

}
