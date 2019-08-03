import { requireAuthApi } from '../axios';

export function getStoryFeed() {
    return requireAuthApi.get("/me/storyfeed");
}

export function getUserStoryBoard(userID) {
    return requireAuthApi.get(`/users/${userID}/storyboard`);
}

export function createStoryView(storyID) {
    return requireAuthApi.post(`/stories/${storyID}/views`);
}

export function createStory(files, urls) {
    const formData = new FormData();

    urls.forEach(url => formData.append("imageURLs", url));
    Array.from(files).forEach(file => formData.append("images", file));
    return requireAuthApi.post("/stories", formData, { headers: { 'Content-Type': 'multipart/form-data' } });
}

export function reactStory(reactableID) {
    return requireAuthApi.post(`/reactables/${reactableID}`);
}

export function unreactStory(reactableID) {
    return requireAuthApi.delete(`/reactables/${reactableID}`);
}