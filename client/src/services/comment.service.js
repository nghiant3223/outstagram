import { requireAuthApi } from '../axios';

export function getReply(cmtableID, cmtID, limit, offset) {
    return requireAuthApi.get(`/commentables/${cmtableID}/comments/${cmtID}/replies?limit=${limit}&offset=${offset}`);
}

export function createReply(cmtableID, cmtID, content) {
    return requireAuthApi.post(`/commentables/${cmtableID}/comments/${cmtID}/replies`, { content });
}