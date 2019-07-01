import { noAuthApi, requireAuthApi } from '../axios';

export function getUser(userID) {
    return requireAuthApi.get(`/users/${userID}`);
}

export function followUser(userID) {
    return requireAuthApi.post(`/follows/${userID}`);
}

export function unfollowUser(userID) {
    return requireAuthApi.delete(`/follows/${userID}`);
}