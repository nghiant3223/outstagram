import { requireAuthApi } from '../axios';

export function getRecentRooms() {
    return requireAuthApi.get("/rooms");
}

export function getRoom(id) {
    return requireAuthApi.get(`/rooms/${id}`)
}

export function createMessage(id, content, type) {
    return requireAuthApi.post(`/rooms/${id}/messages`, { content, type });
}

export function getMessages(id, limit, offset) {
    if (limit == 0 || offset == 0) {
        return requireAuthApi.get(`/rooms/${id}/messages`);
    }

    return requireAuthApi.get(`/rooms/${id}/messages?limit=${limit}&offset=${offset}`);
}

export function createRoom(memberIDs, message) {
    if (!message) {
        return requireAuthApi.post(`/rooms/`, { memberIDs });
    }

    return requireAuthApi.post(`/rooms/`, { memberIDs, "1stMessage": { "content": message.content, type: message.type || 0 } });
}