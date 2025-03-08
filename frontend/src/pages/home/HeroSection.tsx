
import React from "react";
import {
    Box,
    Typography,
    TextField,
    Button,
    Container,
} from "@mui/material";

const HeroSection: React.FC = () => {
    return (
        <Box
            sx={{
                position: "relative",
                width: "99vw",
                height: "90vh",
                display: "flex",
                alignItems: "center",
                justifyContent: "center",
                backgroundColor: "black",
                overflow: "hidden",
                backgroundImage: `url('https://assets.nflxext.com/ffe/siteui/vlv3/158a0e2a-cca4-40f5-86b8-11ea2a281b06/web_tall_panel/CA-en-20241202-TRIFECTA-perspective_5bca1e73-96aa-4036-82aa-12d7b0d9f318_large.jpg')`,
                backgroundSize: "cover",
                backgroundPosition: "center",
            }}
        >
            {/* Background gradient */}
            <Box
                sx={{
                    position: "absolute",
                    inset: 0,
                    background: "linear-gradient(to bottom, black, transparent)",
                    opacity: 0.7,
                    zIndex: 1,
                }}
            />

            {/* Content */}
            <Box
                sx={{
                    position: "relative",
                    zIndex: 2,
                    textAlign: "center",
                    color: "white",
                }}
            >
                <Typography variant="h1" sx={{ fontSize: "2.5rem", fontWeight: "bold", mb: 2 }}>
                    Unlimited Movies and TV Shows
                </Typography>
                <Typography variant="body1" sx={{ fontSize: "1.2rem", mb: 4 }}>
                    Completely Free
                </Typography>

                {/* Form */}
                <Container maxWidth="xs" sx={{ p: 0 }}>
                    <Box
                        component="form"
                        sx={{
                            display: "flex",
                            flexDirection: "column",
                            alignItems: "center",
                        }}
                    >
                        <TextField
                            type="email"
                            placeholder="Email address"
                            variant="outlined"
                            required
                            fullWidth
                            sx={{
                                mb: 2,
                                backgroundColor: "white",
                                borderRadius: 1,
                                "& .MuiOutlinedInput-root": {
                                    borderRadius: "4px",
                                },
                            }}
                        />
                        <Button
                            type="submit"
                            variant="contained"
                            sx={{
                                backgroundColor: "#e50914",
                                color: "white",
                                px: 4,
                                py: 1,
                                borderRadius: "4px",
                                "&:hover": {
                                    backgroundColor: "#d40813",
                                },
                            }}
                        >
                            Get Started
                        </Button>
                    </Box>
                </Container>
                <Typography variant="body2" sx={{ mt: 3, fontSize: "0.9rem" }}>
                    Ready to watch? Enter your email to create your account
                </Typography>
            </Box>
        </Box>
    );
};

export default HeroSection;
