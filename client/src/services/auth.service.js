import { noAuthApi, requireAuthApi } from '../axios';

export function loginUser(username, password) {
    return noAuthApi.post(`/auth/login`, { username, password })
}

export function getMe() {
    return requireAuthApi.get("/me")
}

export function registerUser(fullname, email, username, password, avatar) {
    const formData = new FormData();

    formData.append("avatar", avatar);
    formData.append("fullname", fullname);
    formData.append("email", email);
    formData.append("username", username);
    formData.append("password", password);

    return noAuthApi.post('/auth/register', formData)
}

export function logoutUser() {
    return requireAuthApi.post("/auth/logout");
}