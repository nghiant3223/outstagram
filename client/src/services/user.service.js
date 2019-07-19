import { noAuthApi, requireAuthApi } from '../axios';

export function getUser(username) {
    return requireAuthApi.get(`/users/${username}`);
}

export function followUser(userID) {
    return requireAuthApi.post(`/follows/${userID}`);
}

export function unfollowUser(userID) {
    return requireAuthApi.delete(`/follows/${userID}`);
}

export function getNewsFeed() {
    return requireAuthApi.get("/me/newsfeed");
}