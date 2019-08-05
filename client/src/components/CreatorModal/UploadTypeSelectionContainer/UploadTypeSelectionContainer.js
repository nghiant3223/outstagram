import React, { useState } from 'react';
import { Icon, Button } from 'semantic-ui-react';

import RadioButton from "../../RadioButton/RadioButton";

import "./UploadTypeSelectionContainer.css";

function UploadTypeSelectionContainer({ closeModal, onImagesUpload, onUploadTypeChange, isLoading, type }) {
    return (
        <div className="CreatorModal__UploadTypeContainer">
            <div className="CreatorModal__UploadTypeContainer__Radio" onChange={onUploadTypeChange}>
                <RadioButton name="upload-type" value="NEWSFEED" defaultChecked={type === "NEWSFEED"}>
                    <Icon name="newspaper outline" size="large" color="blue" />Newsfeed
                </RadioButton>

                <RadioButton name="upload-type" value="STORY" defaultChecked={type === "STORY"}>
                    <Icon className="" name="image" size="large" color="blue" />Story
                </RadioButton>
            </div>

            <div>
                <button className="ui button" type="button" onClick={closeModal}>Cancel</button>
                {isLoading ?
                    <Button onClick={onImagesUpload} loading primary>Upload</Button>
                    :
                    <Button onClick={onImagesUpload} primary>Upload</Button>
                }
            </div>
        </div>
    )
}

export default UploadTypeSelectionContainer