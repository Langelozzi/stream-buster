import { CurrentlyWatching } from "../../models/currently_watching";
import instance from "../axios";

export const getCurrentlyWatching = async () => {
    try {
        const result = await instance.get("/currently-watching/getall");
        console.log('result', result);
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

export const getWatchList = async (): Promise<CurrentlyWatching[]> => {
    try {
        const result = await instance.get("/currently-watching/watchlist");
        return result.data;
    } catch (error) {
        console.error('Error fetching watch list:', error);
        throw error;
    }
}
