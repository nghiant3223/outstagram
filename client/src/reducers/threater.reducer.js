import * as actionTypes from '../constants/actionTypes';

const initialState = {
    isModalOpen: false
};

export default function storyReducer(state = initialState, action) {
    switch (action.type) {
        case actionTypes.OPEN_THEATER_MODAL: {
            return { ...state, isModalOpen: true }
        }

        case actionTypes.CLOSE_THEATER_MODAL:
            return { ...state, isModalOpen: false }

        default:
            return state
    }
}