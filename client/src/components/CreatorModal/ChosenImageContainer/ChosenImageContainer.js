import React, { Component } from 'react'
import { genUID } from '../../../utils/lang';
import { Icon, Button } from 'semantic-ui-react';

import "./ChosenImageContainer.css";

class ChosenImageContainer extends Component {
    shouldComponentUpdate(nextProps) {
        return this.props.renderImages.length !== nextProps.renderImages.length;
    }

    getImageURL({ file: fileOrURL, id }, removeImage) {
        const isURL = typeof fileOrURL === "string"
        return <div key={id} className="CreatorModal__ChosendImageContainer__ChosenImage__Close">
            <Button circular className="CreatorModal__ChosendImageContainer__ChosenImage__CloseButton" onClick={() => removeImage(id)}>
                <Icon name="close" color="red" inverted />
            </Button>
            <div className="CreatorModal__ChosendImageContainer__ChosenImage" style={{ backgroundImage: `url(${isURL ? fileOrURL : URL.createObjectURL(fileOrURL)})` }} />
        </div>;
    }

    render() {
        const { renderImages, removeImage } = this.props;
        return (
            <div className="CreatorModal__ChosendImageContainer__Images"> {renderImages.map((image) => this.getImageURL(image, removeImage))} </div>
        )
    }
}

export default ChosenImageContainer;