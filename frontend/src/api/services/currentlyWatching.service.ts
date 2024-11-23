import { CurrentlyWatching } from "../../models/currently_watching";
import { Episode } from "../../models/episode";
import { createMedia } from "./media.service";
import instance from "../axios";
import { UserContextType } from "../../contexts/UserContext";
import { Movie } from "../../models/movie";
import { TV } from "../../models/tv";
import { getFormattedDate, compareByUpdatedAt } from "../../utils/date.helpter";

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

export const deleteCurrentlyWatching = async (id: number) => {
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
        return result.data.sort(compareByUpdatedAt);
    } catch (error) {
        console.error('Error fetching watch list:', error);
        throw error;
    }
}
export const onAddToList = async (media: Movie | TV, user: UserContextType, seasonNumber: number = 0, episodeNumber: number = 0) => {
    try {
        let mediaResponse;
        try {
            mediaResponse = await createMedia(media.Media!)
        } catch (error) {
            console.error(error);

        }

        const currentlyWatching: CurrentlyWatching = {
            MediaId: mediaResponse?.ID,
            UserID: user.user?.ID,
            SeasonNumber: seasonNumber,
            EpisodeNumber: episodeNumber,
            UpdatedAt: getFormattedDate()
        }

        await createCurrentlyWatching(currentlyWatching)

    } catch (error) {
        console.error("Error addign to list")
    }

}
