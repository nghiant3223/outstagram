import * as actionTypes from '../constants/actionTypes';
import { setToken, getToken } from '../localStorage';
import * as authServices from '../services/auth.service';

export function loginUser(username, password) {
    return async function (dispatch) {
        try {
            const { data: { data: token } } = await authServices.loginUser(username, password)
            setToken(token)
            console.log(getToken()); 
            const { data: { data: user } } = await authServices.getMe();
            dispatch({ type: actionTypes.AUTH_SUCCESS, payload: user });
        } catch (e) {
            dispatch({ type: actionTypes.AUTH_FAIL });
        }
    }
}

export function logoutUser() {
    localStorage.clear();
    return { type: actionTypes.LOGOUT };
}

export function getMe() {
    return async function (dispatch) {
        try {
            const { data: { data: user } } = await authServices.getMe();
            dispatch({ type: actionTypes.AUTH_SUCCESS, payload: user });
        } catch (e) {
            dispatch({ type: actionTypes.AUTH_FAIL });
        }
    }
}