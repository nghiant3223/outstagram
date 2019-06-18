import * as actionTypes from '../constants/actionTypes';
import { setToken, getToken } from '../localStorage';
import * as authServices from '../services/auth.service';

import Socket from '../socket';

export const loginUser = (username, password) =>
    async (dispatch) => {
        try {
            const { data: { data: token } } = await authServices.loginUser(username, password)
            setToken(token)
            const { data: { data: user } } = await authServices.getMe();
            dispatch({ type: actionTypes.AUTH_SUCCESS, payload: user });
        } catch (e) {
            alert(e);
            dispatch({ type: actionTypes.AUTH_FAIL });
        }
    }

export const logoutUser = () => {
    localStorage.clear();
    Socket.close();
    return { type: actionTypes.LOGOUT };
}

export const getMe = () =>
    async (dispatch) => {
        try {
            const { data: { data: user } } = await authServices.getMe();
            dispatch({ type: actionTypes.AUTH_SUCCESS, payload: user });
        } catch (e) {
            dispatch({ type: actionTypes.AUTH_FAIL });
        }
    }