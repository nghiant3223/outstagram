import { requireAuthApi } from '../axios';

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

export function localSearchUser(users, filterText) {
    const lowerCaseText = filterText.toLowerCase();
    const results = users.filter((user) => {
        const usernameMatch = user.username && user.username.toLowerCase().includes(lowerCaseText);
        const fullnameMatch = user.fullname && user.fullname.toLowerCase().includes(lowerCaseText);
        const emailMatch = user.email && user.email.toLowerCase().includes(lowerCaseText);
        return usernameMatch || fullnameMatch || emailMatch;
    });
    return results;
}

export function updateUser(body) {
    const formData = new FormData();

    Object.keys(body).forEach((key) => {
        if (key) formData.append(key, body[key]);
    })

    return requireAuthApi.patch("/me", formData);
}