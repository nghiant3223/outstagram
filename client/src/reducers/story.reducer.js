import * as actionTypes from '../constants/actionTypes';

const initialState = {
    isModalOpen: false,
    onDisplaySBNode: null
};

export default function storyReducer(state = initialState, action) {
    switch (action.type) {
        case actionTypes.OPEN_STORY_MODAL: {
            const board = action.payload.getValue()
            return board.stories === null ? state : { ...state, isModalOpen: true, onDisplaySBNode: action.payload }
        }

        case actionTypes.DISPLAY_SB_NODE: {
            const board = action.payload.getValue()
            return board.stories === null ? { ...state, onDisplaySBNode: null, isModalOpen: false } : { ...state, onDisplaySBNode: action.payload, isModalOpen: true }
        }

        case actionTypes.LOGOUT: {
            return initialState;
        }

        case actionTypes.CLOSE_STORY_MODAL:
            return { ...state, isModalOpen: false, onDisplaySBNode: null }

        default:
            return state
    }
}