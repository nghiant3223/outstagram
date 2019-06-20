import * as actionTypes from '../constants/actionTypes';

const initialState = {
    isLoading: true,
    isStoryModalOpen: true
};

export default function uiReducer(state = initialState, action) {
    switch (action.type) {
        case actionTypes.AUTH_SUCCESS:
            return { ...state, isLoading: false }

        case actionTypes.AUTH_FAIL:
            return { ...state, isLoading: false }

        case actionTypes.OPEN_STORY_MODAL: 
        return {...state, isStoryModalOpen: true}

        case actionTypes.CLOSE_STORY_MODAL:
            return {...state, isStoryModalOpen: false}

        default:
            return state;
    }
};