import { requireAuthApi } from '../axios';

export function getRecentRooms() {
    return requireAuthApi.get("/rooms");
}

export function getRoom(id) {
    return requireAuthApi.get(`/rooms/${id}`)
}

export function createMessage(id, content, type) {
    return requireAuthApi.post(`rooms/${id}/messages`, { content, type });
}