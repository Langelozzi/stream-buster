import axios from "../axios";

export const getCurrentlyWatching = async () => {
    const result = await axios.get("/currently-watching/getall");
    console.log('result', result);
}
export const getWatchList = async () => {
    const result = await axios.get("/currently-watching/watchlist");
    console.log('result', result);
}
