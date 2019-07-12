import { requireAuthApi } from '../axios';

export function react(reactableID) {
    return requireAuthApi.post(`/reactions/${reactableID}`);
}

export function unreact(reactableID) {
    return requireAuthApi.delete(`/reactions/${reactableID}`);
}