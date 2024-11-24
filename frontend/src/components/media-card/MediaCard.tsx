import { Card, CardContent, CardMedia, IconButton } from "@mui/material";
import { Movie } from "../../models/movie";
import { TV } from "../../models/tv";
import { useState } from "react";
import MediaDetailsModal from "../media-details-modal/MediaDetailsModal";
import { CurrentlyWatching } from "../../models/currently_watching";
import { Info, PlayArrow } from "@mui/icons-material";
import DeleteIcon from '@mui/icons-material/Delete';
import { useNavigate } from "react-router-dom";
import { useSnackbar } from "../../hooks/useSnackBar";



interface MediaCardProps {
    media: TV | Movie
    currentlyWatching?: CurrentlyWatching | undefined,
    search?: boolean
    onDelete?: (mediaId: number) => Promise<void> | undefined
}

const cardContentStyles = {
    display: "flex",
    justifyContent: "space-between",

    padding: 1,
    color: '#ffffff',
    position: 'absolute', // Keeps it at the bottom
    bottom: 0,
    left: 0,
    right: 0,
    height: 0, // Starts hidden
    opacity: 0, // Fully transparent
    backgroundColor: '#121212', // Background for the sliding content
    transition: 'height 0.3s ease, opacity 0.3s ease', // Smooth animation
    '&:hover': {
        height: '20px', // Expands to reveal content
        opacity: .7, // Becomes visible
    },
}

export const MediaCard: React.FC<MediaCardProps> = ({ media, currentlyWatching, search = true, onDelete }) => {
    // State
    const [isModalOpen, setIsModalOpen] = useState(false);
    const navigate = useNavigate();
    const { showSnackbar, SnackbarComponent } = useSnackbar()

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

    const handlePlayClick = () => {
        if (currentlyWatching?.EpisodeNumber! > 0) {
            navigate(`/watch/${media.Media?.TMDBID}/${currentlyWatching?.SeasonNumber}/${currentlyWatching?.EpisodeNumber}`,
                // todo get the media
                { state: { media: media.Media } }
            );
        } else {
            navigate(`/watch/${media.Media?.TMDBID}`,
                { state: { media } }
            );
        }
    }
    const handleClickDelete = async () => {
        try {

            if (!currentlyWatching?.MediaId || !onDelete) {
                showSnackbar("Error Deleting")
                return
            }

            await onDelete(currentlyWatching?.MediaId)
            showSnackbar("Successfully deleted")
        } catch (error) {
            showSnackbar("Error Deleting")
        }

    }


    return (
        <>

            <Card
                sx={{
                    maxWidth: 150,
                    borderRadius: 2,
                    boxShadow: 5,
                    backgroundColor: '#181818',
                    transition: 'transform 0.3s ease, box-shadow 0.3s ease, height 0.3s ease',
                    overflow: 'hidden', // Ensures no content overflows the card
                    position: 'relative', // Ensures the sliding content aligns properly
                    '&:hover': {
                        transform: 'translateY(-8px)', // Moves the card up
                        boxShadow: '0px 10px 20px rgba(0, 0, 0, 0.7)', // Increases box shadow
                        height: 'auto', // Allows the card to grow
                    },
                }}
            >
                {/* Movie Poster */}
                <CardMedia
                    onClick={handleClick}
                    component="img"
                    image={media.Media?.PosterImage}
                    alt={media.Media?.Title}
                    sx={{ objectFit: 'cover' }}
                />

                {!search && (
                    <>
                        {/* Sliding Card Content */}
                        < CardContent sx={cardContentStyles}>
                            <IconButton onClick={handlePlayClick} sx={{ color: '#fff' }}>
                                <PlayArrow />
                            </IconButton>
                            <IconButton onClick={handleClick} sx={{ color: '#fff' }}>
                                <Info />
                            </IconButton>
                            <IconButton onClick={handleClickDelete} sx={{ color: '#fff' }}>
                                <DeleteIcon />
                            </IconButton>
                        </CardContent>
                    </>
                )}
            </Card >
            {media && (
                <MediaDetailsModal currentSeasonNumber={currentlyWatching?.SeasonNumber} currentEpisodeNumber={currentlyWatching?.EpisodeNumber} media={media} isOpen={isModalOpen} onClose={handleCloseModal} />
            )
            }
            {SnackbarComponent}
        </>
    )
}
