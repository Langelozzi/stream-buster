import { useUser } from "../../hooks/useUser"
import AskQuery from "../../components/chat/AskQuery";

export const HomePage = () => {
    const user = useUser();
    console.log('user', user);

    return (
        <>
            <div>Welcome to stream buster</div>
            <div>login or sign up</div>
            <AskQuery></AskQuery>
        </>
    )
}
