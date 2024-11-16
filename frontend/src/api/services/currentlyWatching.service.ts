import { CurrentlyWatching } from "../../models/currently_watching";
import axios from "../axios";

export const getCurrentlyWatching = async () => {
    const result = await axios.get("/currently-watching/getall");
    console.log('result', result);
}

export const createCurrentlyWatching = async (data: any) => {
    const result = await axios.post("/currently-watching/", data)
    return result
}

export const getWatchList = async (): Promise<CurrentlyWatching[]> => {
    const result = await axios.get("/currently-watching/watchlist");
    return result.data;
}
