import { useUser } from "../../hooks/useUser"
import { useEffect, useState } from "react";
import { TV } from "../../models/tv";
import { Movie } from "../../models/movie";
import { createCurrentlyWatching } from "../../api/services/currentlyWatching.service";
import { createMedia } from "../../api/services/media.service";
import { searchMulti } from "../../api/services/search.service";
import { getWatchList } from "../../api/services/currentlyWatching.service";
import MediaList from "../../components/media-list/medialist";

export const HomePage = () => {
    const user = useUser();
    const [media, setMedia] = useState([])
    const [first, setFirst] = useState(0)
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
                if (currentlyWatching.Media.MediaType.Name === 'tv') {
                    return currentlyWatching.Media as TV;
                } else if (currentlyWatching.Media.MediaType.Name === 'movie') {
                    return currentlyWatching.Media as Movie;
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
