import apiClient from "../../services/apiClient";



type LoginPayload = {
    email: string;
    password: string;
}


type LoginResponse = {
    token: string;
    message: string;
}


export const login = async (payload: LoginPayload): Promise<LoginResponse> => {
    const response = await apiClient.post<LoginResponse>("/auth/login", payload);
    return response.data;
}