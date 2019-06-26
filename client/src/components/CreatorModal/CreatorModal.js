import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Modal, Button } from 'semantic-ui-react';

import * as storyServices from '../../services/story.service';
import * as creatorActions from '../../actions/creator.action';

import StoryFeedManager from '../../StoryFeedManager';
import socket from '../../Socket';

class CreatorModal extends Component {
    state = {
        isLoading: false,
        uploadImages: []
    }

    onImagesUpload = async () => {
        const { uploadImages } = this.state;
        const { updateStoryFeed, closeModal } = this.props;
        const storyFeedManager = StoryFeedManager.getInstance();

        try {
            this.setState({ isLoading: true });

            const { data: { data: { stories } } } = await storyServices.createStory(uploadImages);

            stories.forEach((story) => storyFeedManager.prependStory(story));
            socket.emit("STORY.CLIENT.POST_STORY", stories);
            updateStoryFeed();
            closeModal();
        } catch (e) {
            console.log(e);
        } finally {
            this.setState({ isLoading: false });
        }
    }

    onFileInputChange = (e) => {
        e.persist();
        const files = e.target.files;
        this.setState({ uploadImages: files });
    }

    render() {
        const { isLoading } = this.state;
        const { isModalOpen, closeModal } = this.props;

        return (
            <div>
                <Modal
                    closeOnEscape
                    centered={false}
                    closeOnDimmerClick
                    open={isModalOpen}
                    onClose={closeModal}>
                    <input type="file" multiple onClick={e => e.target.value = null} onChange={this.onFileInputChange} />
                    <Button positive loading={isLoading} onClick={this.onImagesUpload}>Upload</Button>
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