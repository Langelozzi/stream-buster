import { useUser } from "../../hooks/useUser"

export const DashboardPage = () => {
    const user = useUser();
    console.log('user', user);

    return (
        <div>
            Dashboard page
        </div>
    )
}