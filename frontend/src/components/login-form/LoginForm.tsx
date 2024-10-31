import React, { useState } from 'react';
import { TextField, Button, Box } from '@mui/material';
import { postLogin } from '../../api/services/auth.service';
import { useNavigate } from 'react-router-dom';
import { useUser } from '../../hooks/useUser';

export const LoginForm: React.FC = () => {
	// State for email and password
	const [email, setEmail] = useState<string>('');
	const [password, setPassword] = useState<string>('');

	const navigate = useNavigate();
	const { login } = useUser();
	// Handle form submission
	const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
		event.preventDefault();
		const form = new FormData();
		form.append("email", email);
		form.append("password", password);

		const res = await postLogin(form);
		if (res?.status === 200 && res.data) {
			const { user, token } = res.data;
			login(user, token);
			navigate('/dashboard');  // Redirect without reloading the page
		} else {
			alert('Login failed');
		}
	};

	return (
		<Box
			component="form"
			onSubmit={handleSubmit}
			sx={{ display: 'flex', flexDirection: 'column', width: '300px', margin: '0 auto', gap: 2 }}
		>
			{/* Email field */}
			<TextField
				label="Username"
				variant="outlined"
				type="text"
				value={email}
				onChange={(e) => setEmail(e.target.value)}
				fullWidth
				required
			/>

			{/* Password field */}
			<TextField
				label="Password"
				variant="outlined"
				type="password"
				value={password}
				onChange={(e) => setPassword(e.target.value)}
				fullWidth
				required
			/>

			{/* Submit button */}
			<Button type="submit" variant="contained" color="primary">
				Submit
			</Button>
		</Box>
	);
};
