import axios from 'axios';

const BASE_URL = import.meta.env.VITE_API_BASE_URL;

export const searchMulti = async (query: string) => {
  try {
    const response = await axios.get(`${BASE_URL}/search/multi`, {
      params: {
        query: query,
      },
    });
    return response.data;
  } catch (error) {
    console.error('Error fetching search results:', error);
    throw error;
  }
};