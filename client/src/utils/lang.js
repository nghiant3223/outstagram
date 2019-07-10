export function capitalize(string) {
    return string.charAt(0).toUpperCase() + string.slice(1);
}

export function genUID() {
    return '_' + Math.random().toString(36).substr(2, 9);
}