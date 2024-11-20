import { useEffect, useState } from "react";
import { getWatchList } from "../../api/services/currentlyWatching.service";
import MediaList from "../../components/media-list/medialist";
import { castToTvOrMovie } from "../../api/services/search.service";
import { useUser } from "../../hooks/useUser";

export const HomePage = () => {
    const user = useUser()
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
        }
        getMediaList()

    }, [])

    return (
        <>
            {user ? (
                <>
                    <h1>Continue Watching</h1>
                    <MediaList media={media} ></MediaList>
                </>
            ) : (
                <>
                    <div>Welcome to stream buster</div>
                    <div>login or sign up</div>
                </>
            )}
        </>
    )
}
