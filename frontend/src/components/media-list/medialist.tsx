import React, { useEffect, useState } from 'react';
import { Box } from '@mui/material';
import { MediaCard } from '../media-card/MediaCard';
import { castToTvOrMovie } from '../../api/services/search.service';
import { CurrentlyWatching } from '../../models/currently_watching';
import { Movie } from '../../models/movie';
import { TV } from '../../models/tv';
import { Media } from '../../models/media';
import _ from 'lodash';
import { deleteCurrentlyWatching } from '../../api/services/currentlyWatching.service';
interface MediaCarouselProps {
	currentlyWatchings: CurrentlyWatching[];
}

const MediaList: React.FC<MediaCarouselProps> = ({ currentlyWatchings }) => {
	const [media, setMedia] = useState<(TV | Movie | Media | undefined)[]>();

	useEffect(() => {
		if (!currentlyWatchings) {
			return
		}
		const mediaList = currentlyWatchings.map((currentlyWatching) => {
			try {
				if (!currentlyWatching.Media || !currentlyWatching.Media.MediaType) {
					return
				} else if (currentlyWatching.Media?.MediaType.Name === 'tv') {
					return castToTvOrMovie(currentlyWatching.Media);
				} else if (currentlyWatching.Media?.MediaType.Name === 'movie') {
					return castToTvOrMovie(currentlyWatching.Media);
				} else {
					return currentlyWatching.Media;
				}
			} catch (error) {
			}
		})
		setMedia(mediaList)
	}, [currentlyWatchings])

	const onDelete = async (mediaId: number) => {
		try {
			const removeMediaItem = (mediaId: number) => {
				setMedia((prevMedia) => {
					return prevMedia.filter((mediaItem) => mediaItem.ID != mediaId)
				})
			}
			removeMediaItem(mediaId)
			await deleteCurrentlyWatching(mediaId)
		} catch (error) {
			console.error("Error deleteing currently watching" + error)
			throw error
		}
	}

	return (
		<Box sx={{ display: "flex", gap: "10px", flexWrap: "wrap" }}>
			{media && media.map((mediaObj, index) => {
				const tvOrMovie = castToTvOrMovie({ Media: mediaObj })
				return <MediaCard search={false} currentlyWatching={currentlyWatchings[index]} key={index} media={tvOrMovie} onDelete={onDelete}></MediaCard>
			})}
		</Box >
	);
};

export default MediaList;
