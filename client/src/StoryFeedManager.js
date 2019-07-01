import DoublyLinkedList from './ds/DoubleLinkedList';
import { getUserStoryBoard } from './services/story.service';
import DoublyLinkedListNode from './ds/DoubleLinkedListNode';
import { FormInput } from 'semantic-ui-react';

class StoryFeedManager {
    constructor(storyBoards) {
        // Storyboards of people who logged in user follows
        this._ll = new DoublyLinkedList();

        // Storyboard of logged in user
        this._head = new DoublyLinkedListNode(storyBoards[0]);

        for (let i = 1; i < storyBoards.length; i++) {
            this._ll.append(storyBoards[i]);
        }

        this._makeLink();
    }

    get1stSBNode() {
        // If current user has no stories, return the next story board
        if (this._head.getValue().stories === null) {
            return this._ll.getHead();
        }

        // If current user has story
        return this._head;
    }

    // Prepend story to current logged in user
    prependStory(...stories) {
        const userBoard = this._head.getValue();

        if (userBoard.stories === null) {
            userBoard.stories = [];
        }

        stories.forEach((story) => userBoard.stories.unshift(story));
    }

    // Prepend story to specific storyboard
    async prependUserStory(userID, ...stories) {
        const sbNode = this._ll.find({ callback: (nodeValue) => nodeValue.userID === userID });

        // In case user's storyboard does not appear in current logged in user's story feed (the former user is the user who has id equals `userID`)
        if (sbNode === null) {
            const { data: { data: { storyBoard: userBoard } } } = await getUserStoryBoard(userID);
            this._ll.prepend(userBoard);
        } else {
            const storyBoard = sbNode.getValue();
            this._ll.delete(storyBoard);
            this._ll.prepend(storyBoard);
            storyBoard.hasNewStory = true;
            stories.forEach((story) => storyBoard.stories.unshift(story));
        }

        this._makeLink();
    }

    map(hof) {
        let sbNode = this._head;
        let results = [];

        while (sbNode !== null) {
            results.push(hof(sbNode));
            sbNode = sbNode.getNext();
        }

        return results;
    }

    // Push storyboard node to the inactive linkedlist, in which all storyboard's `hasNewStory` = false
    inactiveSB(sbNode) {
        const storyBoard = sbNode.getValue();

        if (sbNode === this._head) {
            storyBoard.hasNewStory = false;
            return;
        }

        if (storyBoard.stories.every((story) => story.seen)) {
            storyBoard.hasNewStory = false;

            this._ll.delete(storyBoard);
            this._ll.append(storyBoard);
            this._makeLink();

            return;
        }
    }

    // Connect first node and the linkedlist
    _makeLink() {
        const llHead = this._ll.getHead();
        this._head.setNext(llHead);

        if (llHead !== null) {
            llHead.setPrevious(this._head);
        }
    }
}

export default (function () {
    let storyFeedManager;

    function initialize(boards) {
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
        initialize
    }
})();