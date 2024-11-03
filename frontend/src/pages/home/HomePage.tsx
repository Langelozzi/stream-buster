import { useUser } from "../../hooks/useUser"

export const HomePage = () => {
    const user = useUser();
    console.log('user', user);

    return (
        <>
            <div>Welcome to stream buster</div>
            <div>login or sign up</div>
        </>
    )
}
