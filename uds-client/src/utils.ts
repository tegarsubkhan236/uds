import {AxiosError} from "axios";
import {jwtDecode} from "jwt-decode";

export type JwtPayload = {
    user_id : number,
    username: string,
    email: string,
    status: number,
    role: {
        id: number,
        name: string,
    }[],
    permissions: string[],
    exp: number
}

export const handleError = (err: unknown): string => {
    if (err instanceof AxiosError) {
        return err.response?.data?.message || "Server error occurred";
    }
    return 'An unknown error occurred';
}

export const decodeToken = (token: string): JwtPayload | null => {
    try {
        return jwtDecode<JwtPayload>(token)
    } catch {
        return null;
    }
}

export const isTokenExpired = (token: string): boolean => {
    const decoded = decodeToken(token)
    if (!decoded?.exp) return true
    const currentTime = Math.floor(Date.now() / 1000);
    return decoded.exp < currentTime;
}
