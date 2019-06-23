import { requireAuthApi } from '../axios';

export function getStoryFeed() {
    return requireAuthApi.get("/me/storyfeed")
}

export function createStoryView(storyID) {
    return requireAuthApi.post(`/stories/${storyID}/views`);
}