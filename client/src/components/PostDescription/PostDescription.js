import React, { useState } from 'react';
import { TextArea, Button, Form } from 'semantic-ui-react';

import * as postServices from "../../services/post.service";

import "./PostDescription.css";

function PostDescription(props) {
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

    console.log(isPost);

    return (description && !showTextArea) ? <p>{description}</p> :
        (
            showTextArea ?
                <div className="DescriptionContainer">
                    <div>
                        <Form>
                            <TextArea value={description} onChange={(e) => setDescription(e.target.value)} placeholder="Add description" />
                        </Form>
                    </div>
                    <div style={{ marginLeft: "auto" }} className="DescriptionContainer__ButtonContainer">
                        <Button size="mini" onClick={() => setShowTextArea(false)}>Cancel</Button>
                        <Button primary size="mini" onClick={onFormSubmit}>Add</Button>
                    </div>
                </div>
                :
                <p className="ThreaterContainer__InfoContainer__Description__Add" onClick={() => setShowTextArea(true)}>Add description</p>
        )
}

export default PostDescription;