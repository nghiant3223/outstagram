import React, { useState, useRef } from 'react';
import { TextArea, Button, Form } from 'semantic-ui-react';

import * as postServices from "../../services/post.service";

import "./PostDescription.css";

function PostDescription(props) {
    const textArea = useRef(null);
    const { content, ownerID, user, isPost, id } = props;

    if (content) {
        return <p>{content}</p>;
    }

    if (user.id !== ownerID) {
        return null;
    }

    const [showTextArea, setShowTextArea] = useState(false);
    const [description, setDescription] = useState("");

    const onFormSubmit = () => {
        if (isPost) {
            postServices.updateSpecificPost(id, { content: description });
        } else {
            postServices.updatePostImage(id, { content: description });
        }

        setShowTextArea(false);
    }

    const onAddClick = () => {
        setShowTextArea(true);
        setTimeout(() => { textArea.current.focus() }, 0);
    }


    return (description && !showTextArea) ? <p>{description}</p> :
        (
            <div>
                <div className="DescriptionContainer" style={{ display: showTextArea ? "block" : "none" }}>
                    <div>
                        <Form>
                            <input value={description} onChange={(e) => setDescription(e.target.value)} placeholder="Add description" ref={textArea} />
                        </Form>
                    </div>
                    <div style={{ marginLeft: "auto" }} className="DescriptionContainer__ButtonContainer">
                        <Button size="mini" onClick={() => setShowTextArea(false)}>Cancel</Button>
                        <Button primary size="mini" onClick={onFormSubmit}>Add</Button>
                    </div>
                </div>
                <p className="ThreaterContainer__InfoContainer__Description__Add" style={{ display: !showTextArea ? "block" : "none" }} onClick={onAddClick}>
                    Add description
                </p>
            </div>
        )
}

export default PostDescription;