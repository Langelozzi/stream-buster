import { Card, CardContent, CardMedia, Typography } from "@mui/material";
import { Movie } from "../../models/movie";
import { TV } from "../../models/tv";

interface MediaCardProps {
    item: TV | Movie
}

export const MediaCard: React.FC<MediaCardProps> = ({ item }) => {
    return (
        <>
            <Card sx={{ maxWidth: 300, borderRadius: 2, boxShadow: 5, backgroundColor: '#181818' }}>
                {/* Movie Poster */}
                <CardMedia
                    component="img"
                    image={item.media?.posterImage}
                    alt={item.media?.title}
                    sx={{ height: 450, objectFit: 'cover' }}
                />

                {/* Movie Info */}
                <CardContent sx={{ padding: 2, color: '#ffffff' }}>
                    {/* Title */}
                    <Typography variant="h6" fontWeight="bold" gutterBottom>
                        {item.media?.title}
                    </Typography>
                </CardContent>
            </Card>
        </>
    )
}