import React, { useState } from 'react';
import instance from '../../api/axios';
import { TextField, Button, Typography, Paper, List, ListItem, ListItemText, CircularProgress, Card } from '@mui/material';

const styles = {
	card: {
		padding: 2,
		backgroundColor: '#424242', // Slightly lighter grey for contrast
		color: '#ffffff',
		borderRadius: 2,
		boxShadow: '0px 4px 20px rgba(0, 0, 0, 0.5)',
		marginBottom: 2
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
		'& label.MuiFocused': {
			color: '#ffffff',
		},
		'& .MuiInputBaseRoot': {
			color: '#ffffff',
		},
		'& .MuiOutlinedInputRoot': {
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
		marginTop: 2,
		color: '#ffffff',
	},
}

interface AskQueryResponse {
	reply: string;
}

export const AskQuery: React.FC = () => {
	const [query, setQuery] = useState<string>('');
	const [loading, setLoading] = useState<boolean>(false);
	const [error, setError] = useState<string>('');
	const [history, setHistory] = useState<{ query: string; reply: string; }[]>([]);

	const handleQueryChange = (event: React.ChangeEvent<HTMLInputElement>) => {
		setQuery(event.target.value);
	};

	const handleSubmit = async (event: React.FormEvent) => {
		event.preventDefault();
		setLoading(true);
		setError('');

		try {
			const res = await instance.post<AskQueryResponse>('/llm/ask-query',
				[{ role: "user", content: query }],
			);
			setHistory([...history, { query, reply: res.data }]);
			setQuery('');
		} catch (err) {
			setError('An error occurred while fetching the response.');
		} finally {
			setLoading(false);
		}
	};

	return (
		<Card sx={styles.card}>
			<Typography variant="h5">Chatbuster</Typography>
			<Typography variant="body2" gutterBottom>Have a question? Ask our AI powered Chatbuster!</Typography>
			<form onSubmit={handleSubmit} style={{ ...styles.input } as any}>
				<TextField
					fullWidth
					variant="outlined"
					label="Enter your query"
					value={query}
					onChange={handleQueryChange}
					required
					disabled={loading}
					sx={styles.input}
				/>
				<Button variant="contained" color="primary" type="submit" sx={styles.button} disabled={loading}>
					{loading ? <CircularProgress size={24} /> : 'Submit'}
				</Button>
			</form>
			{error && <Typography color="error">{error}</Typography>}
			<List>
				{history.map((item, index) => (
					<Paper key={index} style={{ margin: '8px 0', padding: '16px' }} elevation={3}>
						<ListItem>
							<ListItemText primary={`Query: ${item.query}`} secondary={`Response: ${item.reply}`} />
						</ListItem>
					</Paper>
				))}
			</List>
		</Card>
	);
};
