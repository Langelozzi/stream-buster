import { useEffect } from "react";
import { getCurrentlyWatching } from "../../api/services/currentlyWatching.service";
import { useTranslation } from "react-i18next";
// interface Media {
// 	id: string,
// 	tmdb_id: string,
// 	title: string,
// 	overview: string,
// 	posterImage: string,
// 	mediaTypeId: number,
// }
// interface CurrentlyWatching {
// 	userId: number,
// 	mediaId: number,
// 	episodeNumber: number,
// 	seasonNumber: number
// }

export const WatchList = () => {
	const { t } = useTranslation();

	useEffect(() => {
		getCurrentlyWatching()
	}, [])
	return (<>{t('dictionary.watchList')}</>)
}
