import * as actionTypes from '../constants/actionTypes';

const initialState = {
    isModalOpen: false,
    type: "NEWSFEED"
};

export default function storyReducer(state = initialState, action) {
    switch (action.type) {
        case actionTypes.OPEN_CREATOR_MODAL: {
            return { ...state, isModalOpen: true, type: action.payload }
        }

        case actionTypes.CLOSE_CREATOR_MODAL:
            return { ...state, isModalOpen: false }

        default:
            return state
    }
}