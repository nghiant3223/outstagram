import DoublyLinkedList from './ds/DoubleLinkedList';
import { getUserStoryBoard } from './services/story.service';
import DoublyLinkedListNode from './ds/DoubleLinkedListNode';

class StoryFeedManager {
    constructor(storyBoards) {
        this.ll = new DoublyLinkedList();

        let i = 1;
        this.ll.append(storyBoards[0]);
        while (i < storyBoards.length) {
            this.ll.append(storyBoards[i++]);
        }
    }

    getFirstNode() {
        const head = this.ll.getHead();

        if (head.getValue().stories === null) {
            return head.getNext();
        }

        return head;
    }

    // Prepend story to current logged in user
    prependStory(...stories) {
        const headValue = this.ll.getHead().getValue();

        if (headValue.stories === null) {
            headValue.stories = [];
        }

        stories.forEach((story) => headValue.stories.unshift(story));
    }

    // Prepend story to specific storyboard
    async prependUserStory(userID, ...stories) {
        const node = this.ll.find({ callback: (nodeValue) => nodeValue.userID === userID });

        // In case user's storyboard does not appear in current logged in user's story feed (the former user is the user who has id equals `userID`)
        if (node === null) {
            return getUserStoryBoard(userID).then(({ data: { data: { storyBoard: userStoryBoard } } }) => {
                const storyBoard = userStoryBoard;
                storyBoard.hasNewStory = true;

                const head = this.ll.getHead();
                const afterHead = head.getNext();
                const newNode = new DoublyLinkedListNode(storyBoard);

                // If there is no storyboard other than current logged in user's
                if (afterHead === null) {
                    this.ll.setTail(newNode);
                    head.setNext(newNode);
                    newNode.setPrevious(head);
                    return;
                }

                head.setNext(newNode);
                newNode.setPrevious(head);
                newNode.setNext(afterHead);
                afterHead.setPrevious(newNode);
            });
        }

        const storyBoard = node.getValue();

        // In case user's storyboard has already appeared in current user's story feed (the former user is the user who has id equals `userID`)
        if (storyBoard.stories === null) {
            storyBoard.stories = [];
        }

        storyBoard.hasNewStory = true;
        stories.forEach((story) => storyBoard.stories.unshift(story));
        return Promise.resolve();
    }

    map(hof) {
        let sbNode = this.ll.getHead();
        let results = [];

        while (sbNode !== null) {
            results.push(hof(sbNode));
            sbNode = sbNode.getNext();
        }

        return results;
    }

    // Push storyboard node to the inactive linkedlist, in which all storyboard's `hasNewStory` = false
    inactivateNode(sbNode) {
        if (sbNode === this.ll.getHead()) {
            return;
        }

        const storyBoard = sbNode.getValue();
        if (storyBoard.stories.every((story) => story.seen)) {
            storyBoard.hasNewStory = false;
            this.ll.delete(storyBoard);
            this.ll.append(storyBoard);
        }
    }
}

export default (function () {
    let storyFeedManager;

    function initStoryFeedManager(boards) {
        if (storyFeedManager === undefined) {
            storyFeedManager = new StoryFeedManager(boards);
        } else {
            throw new Error("StoryFeedManager has already existed");
        }
    }

    function getInstance() {
        if (storyFeedManager === undefined) {
            throw new Error("StoryFeedManager hasn't been created yet");
        } else {
            return storyFeedManager;
        }
    }

    function removeInstance() {
        if (storyFeedManager === undefined) {
            throw new Error("StoryFeedManager hasn't been created yet");
        } else {
            storyFeedManager = undefined;
        }
    }

    return {
        getInstance,
        removeInstance,
        initStoryFeedManager
    }
})();