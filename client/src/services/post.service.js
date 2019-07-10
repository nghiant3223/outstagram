import { requireAuthApi } from '../axios';

export function createPost(files, urls, content) {
    const formData = new FormData();

    formData.append("visibility", 0);
    formData.append("content", content);
    urls.forEach(url => formData.append("imageURLs", url));
    Array.from(files).forEach(file => formData.append("images", file));
    return requireAuthApi.post("/posts", formData, { headers: { 'Content-Type': 'multipart/form-data' } });
}
