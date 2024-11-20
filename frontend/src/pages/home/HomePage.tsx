import { useEffect, useState } from "react";
import { getWatchList } from "../../api/services/currentlyWatching.service";
import MediaList from "../../components/media-list/medialist";
import { castToTvOrMovie } from "../../api/services/search.service";

export const HomePage = () => {
    const [media, setMedia] = useState<any[]>([])
    useEffect(() => {
        // const test = async () => {
        //     const res = await searchMulti("How to train your dragon")
        //     console.log('res', res);
        //     const himym = res[0]
        //     console.log('himym', himym);
        //     const createdMedia = await createMedia(himym.Media!)
        // }
        //
        // if (first == 0) {
        //     setFirst(1)
        //     test()
        // }


        // createCurrentlyWatching({
        //     UserID: 1,
        //     MediaId: 1,
        //     EpisodeNumber: 0,
        //     SeasonNumber: 0,
        // })
        //
        const getMediaList = async () => {
            const currentlyWatchingList = await getWatchList()
            const mediaList = currentlyWatchingList.map((currentlyWatching) => {
                if (!currentlyWatching.Media || !currentlyWatching.Media.MediaType) {
                    return
                }
                if (currentlyWatching.Media.MediaType.Name === 'tv') {
                    return castToTvOrMovie(currentlyWatching.Media);
                } else if (currentlyWatching.Media.MediaType.Name === 'movie') {
                    return castToTvOrMovie(currentlyWatching.Media);
                } else {
                    return currentlyWatching.Media;
                }
            })
            setMedia(mediaList)
        }
        getMediaList()

    }, [])

    return (
        <>
            <div>Welcome to stream buster</div>
            <div>login or sign up</div>
            <MediaList media={media} ></MediaList>
        </>
    )
}
