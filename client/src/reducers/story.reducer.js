import * as actionTypes from '../constants/actionTypes';

const initialState = {
    isModalOpen: true
};

export default function storyReducer(state = initialState, action) {
    switch (action.type) {
        case actionTypes.OPEN_STORY_MODAL:
            return { ...state, isModalOpen: true }

        case actionTypes.CLOSE_STORY_MODAL:
            return { ...state, isModalOpen: false }
        default:
            return state
    }
}