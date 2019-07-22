import { requireAuthApi } from '../axios';

export function getRecentRooms() {
    return requireAuthApi.get("/rooms");
}

export function getRoom(id) {
    return requireAuthApi.get(`/rooms/${id}`)
}