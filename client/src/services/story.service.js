import {requireAuthApi} from '../axios';

export function getStoryFeed() {
    return requireAuthApi.get("/me/storyfeed")
}