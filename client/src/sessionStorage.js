export function setToken(token) {
    sessionStorage.setItem("x-access-token", token);
}

export function getToken() {
    return sessionStorage.getItem("x-access-token");
}