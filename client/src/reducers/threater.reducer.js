import * as actionTypes from '../constants/actionTypes';

const initialState = {
    isModalOpen: false,
    onDisplayPost: undefined
};

export default function storyReducer(state = initialState, action) {
    switch (action.type) {
        case actionTypes.OPEN_THEATER_MODAL: {
            return { ...state, isModalOpen: true, onDisplayPost: action.payload }
        }

        case actionTypes.CLOSE_THEATER_MODAL:
            return { ...state, isModalOpen: false, onDisplayPost: undefined }

        default:
            return state
    }
}