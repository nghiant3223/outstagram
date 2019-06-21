import { requireAuthApi } from '../axios';
import DoublyLinkedList from '../ds/DoubleLinkedList';

export function getStoryFeed() {
    return requireAuthApi.get("/me/storyfeed")
}

/**
 * Create two double-linked-lists, the first is for active story board, second is for inactive story board
 * @param {Array} storyBoards Array of story boards
 * @returns {Array} [activeStoryBoardHead, inactiveStoryBoardHead]
 */
export function initStoryBoardLinkedList(storyBoards) {
    let active = new DoublyLinkedList(),
        inactive = new DoublyLinkedList();

    let i = 0;
    while (i < storyBoards.length && storyBoards[i].hasNewStory) {
        active.append(storyBoards[i]);
        i++;
    }

    while (i < storyBoards.length) {
        inactive.append(storyBoards[i]);
        i++;
    }

    // active.getTail().next = inactive.getHead();
    // inactive.getHead().previous = active.getTail();

    console.log(active);

    return [active, inactive];
}