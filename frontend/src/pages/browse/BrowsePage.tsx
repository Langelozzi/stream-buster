import { Search } from "../../components/search/Search";
import { useUser } from "../../hooks/useUser";


export const BrowsePage = () => {
    const user = useUser();
    console.log('user', user);

    const dummyId = 920;

    return (
        <>
            <Search />
        </>
    )
}
