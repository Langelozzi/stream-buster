import axios from "../axios";

export const getCurrentlyWatching = async () => {
    const result = await axios.get("/currently-watching/getall");
    console.log('result', result);
}
