import axios, { AxiosResponse } from 'axios';
import bcrypt from 'bcryptjs';

const hashPassword = async (password: string): Promise<string> => {
	// Generate a salt (you can adjust the salt rounds for more security)
	const saltRounds = 10;
	const salt = await bcrypt.genSalt(saltRounds);

	// Hash the password with the generated salt
	const hashedPassword = await bcrypt.hash(password, salt);

	return hashedPassword;
};

export const sendForm = async (form: FormData) => {
	const res: AxiosResponse = await axios.post(import.meta.env.VITE_API_URL + "/auth/login", form, { withCredentials: true })
	return res
}

export const sendTestRequest = async () => {
	const res: AxiosResponse = await axios.get(import.meta.env.VITE_API_URL + "/auth/test", { withCredentials: true })
	return res
}
