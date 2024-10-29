import axios, { AxiosError, AxiosResponse } from 'axios';
import bcrypt from 'bcryptjs';
import { API_BASE_URL } from '../../utils/constants';

const hashPassword = async (password: string): Promise<string> => {
	// Generate a salt (you can adjust the salt rounds for more security)
	const saltRounds = 10;
	const salt = await bcrypt.genSalt(saltRounds);

	// Hash the password with the generated salt
	const hashedPassword = await bcrypt.hash(password, salt);

	return hashedPassword;
};

export const sendForm = async (form: FormData) => {
	try {
		const res: AxiosResponse = await axios.post(API_BASE_URL + "/auth/login", form, { withCredentials: true })
		return res
	} catch (err: any) {
		if (err.status == 401) {
			alert("Wrong username or password. Try again or click Forgot password to reset it.")
		}
	}
}

export const sendTestRequest = async () => {
	const res: AxiosResponse = await axios.get(API_BASE_URL + "/auth/test", { withCredentials: true })
	return res
}
