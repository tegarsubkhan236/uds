import {AxiosResponse} from "axios";
import axiosInstance from "../axios.ts";

type LoginResponse = {
    data: string;
}

const AuthApi = () => {
    const login = async (username : string, password : string) => {
        const response: AxiosResponse<LoginResponse> = await axiosInstance.post("/auth/login", {
            identity: username,
            password: password,
        });
        return response.data.data
    }

    return {
        login
    }
}

export default AuthApi;
