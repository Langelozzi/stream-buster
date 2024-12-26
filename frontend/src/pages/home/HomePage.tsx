import React from 'react';
import {
    Box,
    Container,
    Typography,
    Paper,
    useTheme
} from '@mui/material';

export const HomePage = () => {
    const theme = useTheme();

    return (
        <Box
            sx={{
                minHeight: '100vh',
                display: 'flex',
                alignItems: 'center',
                backgroundColor: theme.palette.grey[100]
            }}
        >
            <Container maxWidth="sm">
                <Paper
                    elevation={3}
                    sx={{
                        padding: 4,
                        textAlign: 'center',
                        borderRadius: 2
                    }}
                >
                    <Typography
                        variant="h4"
                        component="h1"
                        gutterBottom
                        color="primary"
                        sx={{ fontWeight: 'bold' }}
                    >
                        We'll Be Right Back!
                    </Typography>

                    <Typography variant="body1" gutterBottom sx={{ mb: 3 }}>
                        Our system is currently undergoing maintenance to serve you better.
                        We'll be back online in approximately 24 hours.
                    </Typography>

                    <Box
                        component="img"
                        src="https://media4.giphy.com/media/yhfTY8JL1wIAE/giphy.gif?cid=6c09b952hntb972a96dxtndh8waakt9web6zpb2navyf6err&ep=v1_gifs_search&rid=giphy.gif&ct=g" // Replace with your gif URL
                        alt="Maintenance"
                        sx={{
                            width: '100%',
                            maxWidth: 300,
                            height: 'auto',
                            margin: '20px 0'
                        }}
                    />

                    <Typography variant="body2" color="text.secondary" sx={{ mt: 3 }}>
                        Thank you for your patience!
                    </Typography>

                    <Typography variant="caption" display="block" color="text.secondary" sx={{ mt: 2 }}>
                        Expected completion: {new Date(Date.now() + 24 * 60 * 60 * 1000).toLocaleString()}
                    </Typography>
                </Paper>
            </Container>
        </Box>
    );
};

