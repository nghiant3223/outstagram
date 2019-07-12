import { requireAuthApi } from '../axios';

export function getComment(cmtableID, limit, offset) {
    return requireAuthApi.get(`/commentable/${cmtableID}/comments?limit=${limit}&offset=${offset}`);
}

export function commentPost(cmtableID, content) {
    return requireAuthApi.post(`/commentable/${cmtableID}/comments`, { content });
}