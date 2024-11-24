import React, { useEffect, useState } from 'react';
import { Box } from '@mui/material';
import { MediaCard } from '../media-card/MediaCard';
import { castToTvOrMovie } from '../../api/services/search.service';
import { CurrentlyWatching } from '../../models/currently_watching';
import { deleteCurrentlyWatching } from '../../api/services/currentlyWatching.service';
import { Media } from '../../models/media';
interface MediaCarouselProps {
	currentlyWatchings: CurrentlyWatching[];
}

const MediaList: React.FC<MediaCarouselProps> = ({ currentlyWatchings }) => {
	const [media, setMedia] = useState<(Media)[]>();

	useEffect(() => {
		if (!currentlyWatchings) {
			return
		}
		const mediaList = currentlyWatchings.map((currentlyWatching) => {
			return currentlyWatching.Media;
		}).filter((mediaItem) => mediaItem != undefined)

		setMedia(mediaList)
	}, [currentlyWatchings])

	const onDelete = async (mediaId: number) => {
		try {
			const removeMediaItem = (mediaId: number) => {
				setMedia((prevMedia) => {
					return prevMedia!.filter((mediaItem) => {
						return mediaItem!.ID != mediaId
					})
				})
			}
			await deleteCurrentlyWatching(mediaId)
			removeMediaItem(mediaId)
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
