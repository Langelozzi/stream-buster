import { Container, Box, Typography, Button } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import { routes } from '../../router/Routes';

export const HomePage = () => {
    const navigate = useNavigate();

    const handleLogin = () => {
        navigate(routes.login)
    };

    const handleSignUp = () => {
        navigate(routes.register);
    };

    return (
        <Container maxWidth="sm">
            <Box textAlign="center" mt={5}>
                <Typography variant="h3" gutterBottom>
                    Welcome to StreamBuster
                </Typography>
                <Box mt={3}>
                    <Button variant="contained" color="primary" onClick={handleLogin} fullWidth>
                        Login
                    </Button>
                </Box>
                <Box mt={2}>
                    <Button variant="outlined" color="primary" onClick={handleSignUp} fullWidth>
                        Sign Up
                    </Button>
                </Box>
            </Box>
        </Container>
    );
};
