import axios from 'axios';

import { getToken } from './sessionStorage';
import { objectToQuery } from './utils/http';

export const noAuthApi = axios.create({
    baseURL: "/api",
    timeout: 7500
});

export const requireAuthApi = axios.create({
    baseURL: "/api",
    timeout: 7500
});

export const noAuthStatic = function (url, query) {
    const queryString = query ? `?${objectToQuery(query)}` : ""
    return `/static` + url + queryString;
}

requireAuthApi.interceptors.request.use(function (config) {
    config.headers.Authorization = `Bearer ${getToken()}`;
    return config;
});