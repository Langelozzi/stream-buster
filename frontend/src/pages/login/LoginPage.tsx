import { LoginForm } from "../../components/login-form/LoginForm";
import { getCurrentUser } from "../../api/services/user.service";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

export const LoginPage = () => {
    const navigate = useNavigate()

    // do this to try and get a refresh token
    // the auth middle ware will try and grab the refresh token from the request and if its valid issue a new token
    useEffect(() => {
        const getUser = async () => {
            try {
                const user = await getCurrentUser()
                if (user) {
                    navigate("/")
                }
            } catch (error) {

            }
        }
        getUser()
    }, [])


    return (
        <LoginForm />
    )
}
