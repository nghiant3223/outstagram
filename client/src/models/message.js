export class Message {
    constructor(id, authorID, content, isNew) {
        this.id = id;
        this.authorID = authorID;
        this.content = content;
        this.createdAt = new Date();
        this.isNew = !!isNew;
    }
}