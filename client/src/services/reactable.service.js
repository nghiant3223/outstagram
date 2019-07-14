import { requireAuthApi } from '../axios';

export function react(reactableID) {
    return requireAuthApi.post(`/reactables/${reactableID}`);
}

export function unreact(reactableID) {
    return requireAuthApi.delete(`/reactables/${reactableID}`);
}

export function getReactions(reactableID, limit, offset) {
    return requireAuthApi.get(`/reactables/${reactableID}?limit=${limit}&offset=${offset}`);
}