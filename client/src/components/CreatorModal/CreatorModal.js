import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Modal, Form, Icon, Segment, Grid, Divider, Header } from 'semantic-ui-react';

import * as storyServices from '../../services/story.service';
import * as creatorActions from '../../actions/creator.action';
import * as postServices from '../../services/post.service';

import StoryFeedManager from '../../StoryFeedManager';
import socket from '../../Socket';

import "./CreatorModal.css";
import { isImageUrl } from '../../utils/image';

import Input from '../Input/Input';
import UploadTypeSelectionContainer from './UploadTypeSelectionContainer/UploadTypeSelectionContainer';
import UploadTypeContainer from './UploadTypeContainer/UploadTypeContainer';
import ChosenImageContainer from './ChosenImageContainer/ChosenImageContainer';
import DescriptionInput from './DescriptionInput/DescriptionInput';

const initialState = {
    isLoading: false,
    uploadImages: [],
    uploadUrls: [],
    renderImages: [],
    imageURL: "",
    caption: "",
    uploadType: "story"
}

class CreatorModal extends Component {
    state = {
        ...initialState
    }

    onImagesUpload = async () => {
        const { uploadImages, uploadUrls, uploadType, caption } = this.state;
        const { updateStoryFeed, closeModal } = this.props;
        const storyFeedManager = StoryFeedManager.getInstance();

        try {
            this.setState({ isLoading: true });
            switch (uploadType) {
                case "story":
                    const { data: { data: { stories } } } = await storyServices.createStory(uploadImages, uploadUrls);

                    stories.forEach((story) => storyFeedManager.prependStory(story));
                    socket.emit("STORY.CLIENT.POST_STORY", stories);
                    updateStoryFeed();
                    break;

                case "post":
                    await postServices.createPost(uploadImages, uploadUrls, caption);
                    break;

                default:
            }
        } catch (e) {
            console.log(e);
        } finally {
            this.setState({ isLoading: false });
            closeModal();
        }
    }

    componentWillReceiveProps(nextProps) {
        if (this.props.isModalOpen !== nextProps.isModalOpen) {
            this.setState({ ...initialState });
        }
    }

    onFormSubmit = (e) => {
        e.preventDefault();
    }

    onFileInputChange = (e) => {
        e.persist();
        const files = e.target.files;
        this.setState((prevState) => ({ uploadImages: [...prevState.uploadImages, ...files], renderImages: [...prevState.renderImages, ...files] }))
    }

    onUrlInputChange = (e) => {
        const url = e.target.value;

        this.setState(({ imageURL: url }));

        if (url === "") {
            return;
        }

        isImageUrl(url, (ok) => {
            if (!ok) {
                alert("Not an image");
                return;
            }

            this.setState((prevState) => ({ uploadUrls: [...prevState.uploadUrls, url], imageURL: "", renderImages: [...prevState.renderImages, url] }));
        });
    }

    triggerFileInput = () => {
        this.fileInput.click();
    }

    onCaptionChange = (e) => {
        this.setState({ caption: e.target.value });
    }

    onUploadTypeChange = (e) => {
        this.setState({ uploadType: e.target.value });
    }

    render() {
        const { renderImages, imageURL, caption, uploadType } = this.state;
        const { isModalOpen, closeModal } = this.props;

        return (
            <div>
                <Modal
                    closeOnEscape
                    centered={false}
                    closeOnDimmerClick
                    open={isModalOpen}
                    onClose={closeModal}>
                    <Form onSubmit={this.onFormSubmit}>
                        {
                            this.state.renderImages.length === 0 ?
                                (
                                    <UploadTypeContainer expand imageURL={imageURL} onUrlInputChange={this.onUrlInputChange} triggerFileInput={this.triggerFileInput} />
                                ) :
                                (
                                    <div>
                                        <ChosenImageContainer renderImages={renderImages} />
                                        <UploadTypeContainer expand={false} imageURL={imageURL} onUrlInputChange={this.onUrlInputChange} triggerFileInput={this.triggerFileInput} />
                                        {uploadType === "post" && <DescriptionInput value={caption} onChange={this.onCaptionChange} />}
                                    </div>
                                )
                        }

                        <UploadTypeSelectionContainer onImagesUpload={this.onImagesUpload} closeModal={this.props.closeModal} onUploadTypeChange={this.onUploadTypeChange} />

                        <input type="file" ref={el => this.fileInput = el} multiple onClick={e => e.target.value = null} onChange={this.onFileInputChange} style={{ display: "none" }} uploaType={uploadType} />
                    </Form>
                </Modal>
            </div>
        );
    }
}

const mapStateToProps = ({ creatorReducer: { isModalOpen } }) => ({ isModalOpen });

const mapDispatchToProps = (dispatch) => ({
    closeModal: () => dispatch(creatorActions.closeModal())
});

export default connect(mapStateToProps, mapDispatchToProps)(CreatorModal);