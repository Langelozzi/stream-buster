import axios from 'axios';

// Create an instance of axios and set the default configurations
const instance = axios.create({
	baseURL: import.meta.env.VITE_SERVER_URL,
	withCredentials: true
});

export default instance;
