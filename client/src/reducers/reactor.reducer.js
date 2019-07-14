import * as actionTypes from '../constants/actionTypes';

const initialState = {
    isModalOpen: false,
    reactableID: undefined
};

export default function storyReducer(state = initialState, action) {
    switch (action.type) {
        case actionTypes.OPEN_REACTOR_MODAL: {
            return { ...state, isModalOpen: true, reactableID: action.payload }
        }

        case actionTypes.CLOSE_REACTOR_MODAL: {
            return { ...state, isModalOpen: false, reactableID: undefined }
        }

        default:
            return state
    }
}