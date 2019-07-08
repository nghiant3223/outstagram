import React, { useState } from 'react';
import { Icon, Button } from 'semantic-ui-react';

import RadioButton from "../../RadioButton/RadioButton";

import "./UploadTypeSelectionContainer.css";

function UploadTypeSelectionContainer({ closeModal, onImagesUpload }) {
    const [uploadMethod, setUploadMethod] = useState("story");

    const onUploadMethodChange = (e) => {
        setUploadMethod(e.target.value);
    }

    return (
        <div className="CreatorModal__UploadTypeContainer">
            <div className="CreatorModal__UploadTypeContainer__Radio" onChange={onUploadMethodChange}>
                <RadioButton name="upload-type" value="post" >
                    <Icon name="newspaper outline" size="large" color="blue" />Newsfeed
                </RadioButton>

                <RadioButton name="upload-type" value="story" defaultChecked>
                    <Icon className="" name="image" size="large" color="blue" />Story
                </RadioButton>
            </div>

            <div>
                <button className="ui button" type="button" onClick={closeModal}>Cancel</button>
                <Button onClick={() => onImagesUpload(uploadMethod)} primary>Upload</Button>
            </div>
        </div>
    )
}

export default UploadTypeSelectionContainer