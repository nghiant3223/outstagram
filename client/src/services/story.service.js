import { requireAuthApi } from '../axios';

export function getStoryFeed() {
    return requireAuthApi.get("/me/storyfeed")
}

export function createStoryView(storyID) {
    return requireAuthApi.post(`/stories/${storyID}/views`);
}

export function createStory(files) {
    const formData = new FormData();

    Object.keys(files).forEach((key) => formData.append("images", files[key]));
    return requireAuthApi.post("/stories", formData, { headers: { 'Content-Type': 'multipart/form-data' } });
}