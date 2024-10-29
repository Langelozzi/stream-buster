import axios from 'axios';
import { API_BASE_URL } from '../utils/constants';

// Create an instance of axios and set the default configurations
const instance = axios.create({
	baseURL: API_BASE_URL,
	withCredentials: true
});

export default instance;
