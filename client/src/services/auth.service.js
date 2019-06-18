import { noAuthApi, requireAuthApi } from '../axios';

export function loginUser(username, password) {
    return noAuthApi.post(`/auth/login`, { username, password })
}

export function getMe() {
    return requireAuthApi.get("/me")
}

export function registerUser(fullname, email, username, password) {
    return noAuthApi.post('/auth/register', { fullname, email, username, password })
}