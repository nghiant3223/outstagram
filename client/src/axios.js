import axios from 'axios';

import { getToken } from './sessionStorage';

export const noAuthApi = axios.create({
    baseURL: "/api",
    timeout: 3000
});

export const requireAuthApi = axios.create({
    baseURL: "/api",
    timeout: 3000
});

requireAuthApi.interceptors.request.use(function (config) {
    config.headers.Authorization = `Bearer ${getToken()}`;
    return config;
});