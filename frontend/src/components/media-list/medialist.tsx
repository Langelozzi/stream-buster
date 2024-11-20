import React from 'react';
import { Box } from '@mui/material';
import { MediaCard } from '../media-card/MediaCard';
import { Media } from '../../models/media';
import { castToTvOrMovie } from '../../api/services/search.service';
interface MediaCarouselProps {
	media: Media[];
}

const MediaList: React.FC<MediaCarouselProps> = ({ media }) => {
	return (
		<Box sx={{ padding: 2, display: "flex", gap: "10px", flexWrap: "wrap" }}>
			{media.map((mediaObj, index) => {
				const tvOrMovie = castToTvOrMovie({ Media: mediaObj })
				return <MediaCard key={index} media={tvOrMovie}></MediaCard>
			})}
		</Box>
	);
};

export default MediaList;
