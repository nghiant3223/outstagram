import * as actionTypes from '../constants/actionTypes';
import { setToken } from '../sessionStorage';
import * as authServices from '../services/auth.service';

import Socket from '../Socket';

export const loginUser = (username, password) =>
    async (dispatch) => {
        try {
            const { data: { data: token } } = await authServices.loginUser(username, password)
            setToken(token)
            const { data: { data: user } } = await authServices.getMe();
            Socket.open({ userID: user.id });
            dispatch({ type: actionTypes.AUTH_SUCCESS, payload: user });
        } catch (e) {
            alert(e);
            dispatch({ type: actionTypes.AUTH_FAIL });
        }
    }

export const logoutUser = () => {
    return async dispatch => {
        try {
            await authServices.logoutUser();
            sessionStorage.clear();
            Socket.close();
            dispatch({ type: actionTypes.LOGOUT });
        } catch (e) {
            console.log("Cannot logout user");
        }
    }
}

export const getMe = () =>
    async (dispatch) => {
        try {
            const { data: { data: user } } = await authServices.getMe();
            Socket.open({ userID: user.id });
            dispatch({ type: actionTypes.AUTH_SUCCESS, payload: user });
        } catch (e) {
            dispatch({ type: actionTypes.AUTH_FAIL });
        }
    }

export function updateUserFollowingCount(isIncrement) {
    return { type: actionTypes.UDPATE_FOLLOWING_COUNT, payload: isIncrement ? 1 : -1 }
}