import { Card, CardContent, CardMedia, Typography } from "@mui/material";
import { Movie } from "../../models/movie";
import { TV } from "../../models/tv";
import { useState } from "react";
import MediaDetailsModal from "../media-details-modal/MediaDetailsModal";

interface MediaCardProps {
    media: TV | Movie
}

export const MediaCard: React.FC<MediaCardProps> = ({ media }) => {
    // State
    const [isModalOpen, setIsModalOpen] = useState(false);

    // Functions
    const handleOpenModal = () => {
        setIsModalOpen(true);
    }
    const handleCloseModal = () => {
        setIsModalOpen(false);
    }

    const handleClick = () => {
        handleOpenModal();
    }

    return (
        <>
            <Card onClick={handleClick} sx={{ maxWidth: 300, borderRadius: 2, boxShadow: 5, backgroundColor: '#181818' }}>
                {/* Movie Poster */}
                <CardMedia
                    component="img"
                    image={media.Media?.PosterImage}
                    alt={media.Media?.Title}
                    sx={{ height: 450, objectFit: 'cover' }}
                />

                {/* Movie Info */}
                <CardContent sx={{ padding: 2, color: '#ffffff' }}>
                    {/* Title */}
                    <Typography variant="h6" fontWeight="bold" gutterBottom>
                        {media.Media?.Title}
                    </Typography>
                </CardContent>
            </Card>

            {media && (
                <MediaDetailsModal media={media} isOpen={isModalOpen} onClose={handleCloseModal} />
            )}
        </>
    )
}
