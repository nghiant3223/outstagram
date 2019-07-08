import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Modal, Form, Icon, Segment, Grid, Divider, Header } from 'semantic-ui-react';

import * as storyServices from '../../services/story.service';
import * as creatorActions from '../../actions/creator.action';

import StoryFeedManager from '../../StoryFeedManager';
import socket from '../../Socket';

import "./CreatorModal.css";
import { isImageUrl } from '../../utils/image';

import Input from '../Input/Input';
import UploadTypeContainer from './UploadTypeContainer/UploadTypeContainer';

class CreatorModal extends Component {
    state = {
        isLoading: false,
        uploadImages: [],
        uploadUrls: [],
        renderImages: [],
        imageURL: ""
    }

    onImagesUpload = async (uploadMethod) => {
        const { uploadImages, uploadUrls } = this.state;
        const { updateStoryFeed, closeModal } = this.props;
        const storyFeedManager = StoryFeedManager.getInstance();

        try {
            this.setState({ isLoading: true });
            switch (uploadMethod) {
                case "story":
                    const { data: { data: { stories } } } = await storyServices.createStory(uploadImages, uploadUrls);

                    stories.forEach((story) => storyFeedManager.prependStory(story));
                    socket.emit("STORY.CLIENT.POST_STORY", stories);
                    updateStoryFeed();

                    break;
                case "post":
                default:
            }
        } catch (e) {
            console.log(e);
        } finally {
            this.setState({ isLoading: false });
            closeModal();
        }
    }

    onFormSubmit = () => {
        this.onImagesUpload();
    }

    onFileInputChange = (e) => {
        e.persist();
        const files = e.target.files;
        this.setState((prevState) => ({ uploadImages: [...prevState.uploadImages, ...files], renderImages: [...prevState.renderImages, ...files] }))
    }

    onUrlInputChange = (e) => {
        const url = e.target.value;
        this.setState(({ imageURL: url }));
        isImageUrl(url, (ok) => {
            if (ok) {
                this.setState((prevState) => ({ uploadUrls: [...prevState.uploadUrls, url], imageURL: "", renderImages: [...prevState.renderImages, url] }));
            } else {
                // TODO: Handle URL is not an image"
                alert("Not an image");
            }
        });
    }

    triggerFileInput = () => {
        this.fileInput.click();
    }

    getImageURL(fileOrURL) {
        const isURL = typeof fileOrURL === "string"
        return <div className="CreatorModal__ChosendImageContainer__ChosenImage" style={{ backgroundImage: `url(${isURL ? fileOrURL : URL.createObjectURL(fileOrURL)})` }}></div>
    }

    onUploadMethodChange = (e) => {
        this.setState({ uploadMethod: e.target.value });
    }

    render() {
        const { isLoading, renderImages } = this.state;
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
                        <input type="file" ref={el => this.fileInput = el} multiple onClick={e => e.target.value = null} onChange={this.onFileInputChange} style={{ display: "none" }} />
                        {
                            this.state.renderImages.length === 0 ?
                                (
                                    <Segment placeholder>
                                        <Grid columns={2} stackable textAlign='center'>
                                            <Divider vertical>Or</Divider>
                                            <Grid.Row verticalAlign='middle'>
                                                <Grid.Column>
                                                    <Header icon>
                                                        <Icon name='grid layout' />
                                                        Add new photo
                                                    </Header>
                                                    <button className="ui button primary" type="button" onClick={this.triggerFileInput}>Choose your photos</button>
                                                </Grid.Column>

                                                <Grid.Column>
                                                    <Header icon>
                                                        <Icon name='world' />
                                                        Add photo from web
                                                    </Header>

                                                    <div>
                                                        <Input width="90%" onChange={this.onUrlInputChange} placeHolder="Paste a URL" value={this.state.imageURL} />
                                                    </div>
                                                </Grid.Column>
                                            </Grid.Row>
                                        </Grid>
                                    </Segment>
                                ) :
                                (
                                    <div>
                                        <div className="CreatorModal__ChosendImageContainer__Images">
                                            {Array.from(renderImages).map((image) => this.getImageURL(image))}
                                        </div>
                                        <Segment className="CreatorModal__ChosendImageContainer__More">
                                            <Grid columns={2} stackable textAlign='center'>
                                                <Divider vertical>Or</Divider>
                                                <Grid.Row verticalAlign='middle'>
                                                    <Grid.Column>
                                                        <button className="ui button primary" type="button" onClick={this.triggerFileInput}>Choose another photo</button>
                                                    </Grid.Column>

                                                    <Grid.Column>
                                                        <div className='CreatorModal__ChosendImageContainer__More__URL'>
                                                            <Input width="90%" onChange={this.onUrlInputChange} placeHolder="Paste another URL" value={this.state.imageURL} />
                                                        </div>
                                                    </Grid.Column>
                                                </Grid.Row>
                                            </Grid>
                                        </Segment>
                                    </div>
                                )
                        }

                        <UploadTypeContainer onImagesUpload={this.onImagesUpload} closeModal={this.props.closeModal} />

                    </Form>
                </Modal>
            </div>
        );
    }
}

const mapStateToProps = ({ creatorReducer: { isModalOpen } }) => ({ isModalOpen });

const mapDispatchToProps = (dispatch) => ({
    closeModal: () => dispatch(creatorActions.closeCreatorModal())
});

export default connect(mapStateToProps, mapDispatchToProps)(CreatorModal);