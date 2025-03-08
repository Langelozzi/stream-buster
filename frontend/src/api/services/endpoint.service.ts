import { Endpoint } from "../../models/endpoint";
import axios from "../axios";

export const getAllEndpoints = async (): Promise<Endpoint[]> => {
    try {
        const response = await axios.get(`/endpoint`);
        return response.data as Endpoint[];
    } catch (error) {
        console.error('Error fetching endpoints:', error);
        throw error;
    }
}