import * as actionTypes from '../constants/actionTypes';
import { initStoryBoardLinkedList } from '../services/story.service';

const initialState = {
    isLoading: true,
    isModalOpen: false,
    activeStoryBoardLL: null,
    inactiveStoryBoardLL: null,
    storyBoardLL: null,
    storyBoardNode: null
};

export default function storyReducer(state = initialState, action) {
    switch (action.type) {
        case actionTypes.GET_STORY_FEED:
            const [activeStoryBoardLL, inactiveStoryBoardLL] = initStoryBoardLinkedList(action.payload);
            console.log(activeStoryBoardLL);

            return { ...state, activeStoryBoardLL, inactiveStoryBoardLL, storyBoardLL: activeStoryBoardLL, isLoading: false }

        case actionTypes.OPEN_STORY_MODAL:
            return { ...state, isModalOpen: true, storyBoardNode: action.payload }

        case actionTypes.DISPLAY_STORY_BOARD_NODE:
            return { ...state, storyBoardNode: action.payload }

        case actionTypes.CLOSE_STORY_MODAL:
            return { ...state, isModalOpen: false, storyBoardNode: null }
        default:
            return state
    }
}