import * as actionTypes from '../constants/actionTypes';

const initialState = {
    isModalOpen: false,
    storyBoards: []
};

export default function storyReducer(state = initialState, action) {
    switch (action.type) {
        case actionTypes.GET_STORY_FEED:
            return { ...state, storyBoards: action.payload }

        case actionTypes.OPEN_STORY_MODAL:
            return { ...state, isModalOpen: true }

        case actionTypes.CLOSE_STORY_MODAL:
            return { ...state, isModalOpen: false }
        default:
            return state
    }
}