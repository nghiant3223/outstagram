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

export function getNewsFeed(sinceID) {
    const sinceUrl = sinceID !== undefined ? `&since_id=${sinceID}` : "";
    return requireAuthApi.get("/me/newsfeed?pagination=true" + sinceUrl);
}

export function searchUser(filterText) {
    return requireAuthApi.get("/users?filter=" + filterText);
}