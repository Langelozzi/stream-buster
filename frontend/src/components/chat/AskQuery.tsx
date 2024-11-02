import React, { useState } from 'react';
import instance from '../../api/axios';
import { TextField, Button, Typography, Paper, List, ListItem, ListItemText, CircularProgress, Container } from '@mui/material';

interface AskQueryResponse {
	reply: string;
}

const AskQuery: React.FC = () => {
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
				[{ role: "user", content: query }, { role: "user", content: query }],
			);
			setHistory([...history, { query, reply: res.data.reply }]);
			setQuery('');
		} catch (err) {
			setError('An error occurred while fetching the response.');
		} finally {
			setLoading(false);
		}
	};

	return (
		<Container maxWidth="sm">
			<Typography variant="h4" component="h2" gutterBottom>Ask a Query</Typography>
			<form onSubmit={handleSubmit} style={{ marginBottom: '1rem' }}>
				<TextField
					fullWidth
					variant="outlined"
					label="Enter your query"
					value={query}
					onChange={handleQueryChange}
					required
					disabled={loading}
				/>
				<Button variant="contained" color="primary" type="submit" disabled={loading} fullWidth>
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
		</Container>
	);
};

export default AskQuery;
