import { requireAuthApi } from '../axios';
import DoublyLinkedList from '../ds/DoubleLinkedList';

export function getStoryFeed() {
    return requireAuthApi.get("/me/storyfeed")
}