import { useEffect } from "react";
import { getCurrentlyWatching } from "../../api/services/currentlyWatching.service";
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
	useEffect(() => {
		getCurrentlyWatching()
	}, [])
	return (<>Watch List </>)
}
