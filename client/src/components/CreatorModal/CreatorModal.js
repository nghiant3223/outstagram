import React, { Component } from 'react';
import { connect } from 'react-redux';
import uuid from "uuid/v1";
import { Modal, Form } from 'semantic-ui-react';

import * as storyServices from '../../services/story.service';
import * as creatorActions from '../../actions/creator.action';
import * as postServices from '../../services/post.service';

import StoryFeedManager from '../../StoryFeedManager';
import Socket from '../../Socket';

import "./CreatorModal.css";
import { isImageUrl } from '../../utils/image';

import UploadTypeSelectionContainer from './UploadTypeSelectionContainer/UploadTypeSelectionContainer';
import UploadTypeContainer from './UploadTypeContainer/UploadTypeContainer';
import ChosenImageContainer from './ChosenImageContainer/ChosenImageContainer';
import DescriptionInput from './DescriptionInput/DescriptionInput';
import { validURL } from '../../utils/lang';

const initialState = {
    isLoading: false,
    uploadImages: [],
    uploadUrls: [],
    renderImages: [],
    imageURL: "",
    caption: ""
}

class CreatorModal extends Component {
    constructor(props) {
        super(props);

        this.state = {
            ...initialState,
            uploadType: props.uploadType
        }
    }

    onImagesUpload = async () => {
        const { uploadImages, uploadUrls, uploadType, caption } = this.state;
        const { updateStoryFeed, closeModal } = this.props;
        const storyFeedManager = StoryFeedManager.getInstance();

        try {
            this.setState({ isLoading: true });
            switch (uploadType) {
                case "STORY":
                    const { data: { data: { stories } } } = await storyServices.createStory(uploadImages.map(image => image.file), uploadUrls.map(url => url.file));

                    stories.forEach((story) => storyFeedManager.prependStory(story));
                    Socket.emit("STORY.CLIENT.POST_STORY", stories);
                    updateStoryFeed();
                    break;

                case "NEWSFEED":
                    await postServices.createPost(uploadImages.map(image => image.file), uploadUrls.map(url => url.file), caption);
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
            this.setState({ ...initialState, uploadType: nextProps.uploadType });
        }
    }

    onFormSubmit = (e) => {
        e.preventDefault();
    }

    onFileInputChange = (e) => {
        e.persist();

        const files = e.target.files;
        const chosenFiles = Array.from(files).map((file) => ({ file: file, id: uuid() }));

        this.setState((prevState) => ({
            uploadImages: [
                ...prevState.uploadImages,
                ...chosenFiles
            ],
            renderImages: [
                ...prevState.renderImages,
                ...chosenFiles
            ]
        }));
    }

    onUrlInputChange = (e) => {
        const url = e.target.value;

        this.setState(({ imageURL: url }));

        if (url === "") {
            return;
        }

        if (!validURL(url)) {
            alert("The string you pasted is not URL");
            return
        }

        isImageUrl(url, (ok) => {
            if (!ok) {
                alert("The url is not an image");
                return;
            }

            const id = uuid();

            this.setState((prevState) => ({
                imageURL: "",
                uploadUrls: [
                    ...prevState.uploadUrls,
                    { file: url, id }
                ],
                renderImages: [
                    ...prevState.renderImages,
                    { file: url, id }
                ]
            }));
        });
    }

    removeImage = (id) => {
        this.setState((prevState) => ({
            uploadUrls: prevState.uploadUrls.filter((file) => file.id !== id),
            uploadImages: prevState.uploadImages.filter((file) => file.id !== id),
            renderImages: prevState.renderImages.filter((file) => file.id !== id),
        }));
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
        const { renderImages, imageURL, caption, uploadType, isLoading } = this.state;
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
                                <UploadTypeContainer expand imageURL={imageURL} onUrlInputChange={this.onUrlInputChange} triggerFileInput={this.triggerFileInput} />
                                :
                                <div>
                                    <ChosenImageContainer renderImages={renderImages} removeImage={this.removeImage} />
                                    <UploadTypeContainer expand={false} imageURL={imageURL} onUrlInputChange={this.onUrlInputChange} triggerFileInput={this.triggerFileInput} />
                                    {uploadType === "NEWSFEED" && <DescriptionInput value={caption} onChange={this.onCaptionChange} />}
                                </div>
                        }

                        <UploadTypeSelectionContainer type={uploadType} isLoading={isLoading} onImagesUpload={this.onImagesUpload} closeModal={this.props.closeModal} onUploadTypeChange={this.onUploadTypeChange} />

                        <input type="file" name="images" accept="image/*" ref={el => this.fileInput = el} multiple onClick={e => e.target.value = null} onChange={this.onFileInputChange} style={{ display: "none" }} />
                    </Form>
                </Modal>
            </div>
        );
    }
}

const mapStateToProps = ({ creatorReducer: { isModalOpen, type: uploadType } }) => ({ isModalOpen, uploadType });

const mapDispatchToProps = (dispatch) => ({
    closeModal: () => dispatch(creatorActions.closeModal())
});

export default connect(mapStateToProps, mapDispatchToProps)(CreatorModal);