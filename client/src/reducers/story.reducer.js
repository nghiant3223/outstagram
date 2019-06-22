import * as actionTypes from '../constants/actionTypes';
import StoryFeedManager from '../StoryFeedManager';

const initialState = {
    isLoading: true,
    isModalOpen: false,
    storyBoardNode: null
};

export default function storyReducer(state = initialState, action) {
    switch (action.type) {
        case actionTypes.GET_STORY_FEED: {
            const boards = action.payload;
            return { ...state, storyFeedManager: new StoryFeedManager(boards), isLoading: false }
        }

        case actionTypes.OPEN_STORY_MODAL: {
            const board = action.payload.getValue()
            return board.stories === null ? state : { ...state, isModalOpen: true, storyBoardNode: action.payload }
        }

        case actionTypes.DISPLAY_STORY_BOARD_NODE: {
            const board = action.payload.getValue()
            return board.stories === null ? state : { ...state, storyBoardNode: action.payload }
        }

        case actionTypes.CLOSE_STORY_MODAL:
            return { ...state, isModalOpen: false, storyBoardNode: null }

        default:
            return state
    }
}