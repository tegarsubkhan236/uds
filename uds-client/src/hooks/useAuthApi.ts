import {useCallback, useState} from "react";
import AuthApi from "../api/authApi.ts";
import {decodeToken, handleError, isTokenExpired, JwtPayload} from "../utils.ts";

export const useAuthApi = () => {
    const [authToken, setAuthToken] = useState(() => {
        const token = localStorage.getItem("auth_token")
        if (token && !isTokenExpired(token)) {
            return token
        }
        return null;
    });

    const [auth, setAuth] = useState<JwtPayload | null>(() => {
        if (authToken) {
            return decodeToken(authToken)
        }
        return null
    });

    const [loading, setLoading] = useState(false);

    const [error, setError] = useState<string | null>(null);

    const {login} = AuthApi();

    const loginHandler = useCallback(async (username: string, password: string) => {
        setLoading(true)
        setError(null)
        try {
            const newToken = await login(username, password);
            localStorage.setItem("auth_token", newToken);
            setAuthToken(newToken)
            setAuth(decodeToken(newToken))
        } catch (err) {
            setError(handleError(err));
        } finally {
            setLoading(false)
        }
    }, [login])

    const logoutHandler = useCallback(async () => {
        localStorage.removeItem("auth_token");
        setAuthToken(null)
        setAuth(null)
    }, [])

    return {
        isAuthenticated: !!authToken,
        auth,
        loading,
        error,
        loginHandler,
        logoutHandler,
    };
}
