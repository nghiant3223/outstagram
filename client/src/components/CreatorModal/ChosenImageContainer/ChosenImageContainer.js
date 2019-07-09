import React, { Component } from 'react'

class ChosenImageContainer extends Component {
    shouldComponentUpdate(nextProps) {
        return this.props.renderImages.length !== nextProps.renderImages.length;
    }

    getImageURL(fileOrURL) {
        const isURL = typeof fileOrURL === "string"
        return <div className="CreatorModal__ChosendImageContainer__ChosenImage" style={{ backgroundImage: `url(${isURL ? fileOrURL : URL.createObjectURL(fileOrURL)})` }}></div>
    }

    render() {
        const { renderImages } = this.props;
        return (
            <div className="CreatorModal__ChosendImageContainer__Images"> {Array.from(renderImages).map((image) => this.getImageURL(image))} </div>
        )
    }
}

export default ChosenImageContainer;