import { requireAuthApi } from '../axios';

export function createPost(files, urls, content) {
    const formData = new FormData();

    formData.append("visibility", 1);
    formData.append("content", content);
    urls.forEach(url => formData.append("imageURLs", url));
    Array.from(files).forEach(file => formData.append("images", file));
    return requireAuthApi.post("/posts", formData, { headers: { 'Content-Type': 'multipart/form-data' } });
}

export function getPosts(userID, limit, offset) {
    return requireAuthApi.get(`/users/${userID}/posts?limit=${limit}&offset=${offset}`);
}

export function getUserPosts(userID, limit, offset) {
    return requireAuthApi.get(`/users/:${userID}/posts?limit=${limit}&offset=${offset}`);
}

export function getPostImage(postImageID) {
    return requireAuthApi.get(`/posts/images/${postImageID}`);
}

export function getSpecificPost(postID) {
    return requireAuthApi.get(`/posts/specific/${postID}`);
}

export function updateSpecificPost(postID, body) {
    return requireAuthApi.put(`/posts/specific/${postID}`, body);
}

export function updatePostImage(postImageID, body) {
    return requireAuthApi.put(`/posts/images/${postImageID}`, body);
}

export function searchPost(filterText) {
    return requireAuthApi.get("/posts?filter=" + filterText);
}
