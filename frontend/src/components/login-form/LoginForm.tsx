import React, { useState } from 'react';
import { TextField, Button, Typography, Box } from '@mui/material';
import { makeStyles } from '@mui/styles';
import { postLogin } from '../../api/services/auth.service';
import { useNavigate } from 'react-router-dom';
import { useUser } from '../../hooks/useUser';

const useStyles = makeStyles(() => ({
	paper: {
		padding: 16,
		maxWidth: 400,
		backgroundColor: '#424242', // Slightly lighter grey for contrast
		color: '#ffffff',
		borderRadius: 8,
		boxShadow: '0px 4px 20px rgba(0, 0, 0, 0.5)',
		margin: '0 auto',
	},
	form: {
		display: 'flex',
		flexDirection: 'column',
		gap: 16, // Vertical spacing between fields
	},
	input: {
		'& label': {
			color: '#bdbdbd', // Light grey labels
		},
		'& label.Mui-focused': {
			color: '#ffffff',
		},
		'& .MuiInputBase-root': {
			color: '#ffffff',
		},
		'& .MuiOutlinedInput-root': {
			'& fieldset': {
				borderColor: '#757575', // Grey border
			},
			'&:hover fieldset': {
				borderColor: '#ffffff',
			},
			'&.Mui-focused fieldset': {
				borderColor: '#3f51b5', // Blue when focused
			},
		},
	},
	button: {
		marginTop: 16,
		backgroundColor: '#3f51b5',
		color: '#ffffff',
		'&:hover': {
			backgroundColor: '#303f9f',
		},
	},
}));

export const LoginForm: React.FC = () => {
	const classes = useStyles();
	const [email, setEmail] = useState<string>('');
	const [password, setPassword] = useState<string>('');
	const navigate = useNavigate();
	const { login } = useUser();

	const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
		event.preventDefault();
		const form = new FormData();
		form.append("email", email);
		form.append("password", password);

		const res = await postLogin(form);
		if (res?.status === 200 && res.data) {
			const { user, token } = res.data;
			login(user, token);
			navigate('/dashboard');
		} else {
			alert('Login failed');
		}
	};

	return (
		<Box className={classes.paper}>
			<Typography variant="h5" align="center" gutterBottom>
				Login
			</Typography>
			<form onSubmit={handleSubmit} className={classes.form}>
				<TextField
					className={classes.input}
					label="Email"
					variant="outlined"
					type="text"
					value={email}
					onChange={(e) => setEmail(e.target.value)}
					fullWidth
					required
				/>
				<TextField
					className={classes.input}
					label="Password"
					variant="outlined"
					type="password"
					value={password}
					onChange={(e) => setPassword(e.target.value)}
					fullWidth
					required
				/>
				<Button type="submit" variant="contained" fullWidth className={classes.button}>
					Login
				</Button>
			</form>
		</Box>
	);
};
