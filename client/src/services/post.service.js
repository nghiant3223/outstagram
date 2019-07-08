export function createPost(files, urls) {
    const formData = new FormData();

    urls.forEach(url => formData.append("imageURLs", url));
    Array.from(files).forEach(file => formData.append("images", file));
    return requireAuthApi.post("/posts", formData, { headers: { 'Content-Type': 'multipart/form-data' } });
}
