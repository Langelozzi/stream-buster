import { Search } from "../../components/search/Search";
import { useUser } from "../../hooks/useUser";

const HomePage = () => {
    const user = useUser();
    console.log('user', user);

    return (
        <>
            <Search />
        </>
    )
}

export default HomePage;