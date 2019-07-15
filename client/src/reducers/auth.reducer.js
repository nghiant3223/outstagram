import * as actionTypes from '../constants/actionTypes';

const initialState = {
    isAuthenticated: undefined,
    user: undefined
};

export default function authReducer(state = initialState, action) {
    switch (action.type) {
        case actionTypes.AUTH_FAIL:
            return { ...state, isAuthenticated: false, user: {} };

        case actionTypes.LOGOUT:
            return { ...state, isAuthenticated: false, user: undefined };

        case actionTypes.AUTH_SUCCESS:
            return { ...state, isAuthenticated: true, user: action.payload };

        case actionTypes.UDPATE_FOLLOWING_COUNT: {
            const { user } = state;
            user.followingCount += action.payload;
            return { ...state, user };
        }

        default:
            return state;
    }
}