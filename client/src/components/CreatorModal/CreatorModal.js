import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Modal, Button } from 'semantic-ui-react';

import * as storyServices from '../../services/story.service';
import * as creatorActions from '../../actions/creator.action';

import StoryManager from '../../StoryFeedManager';

class CreatorModal extends Component {
    state = {
        isLoading: false,
        uploadImages: []
    }

    onImagesUpload = async () => {
        const { uploadImages } = this.state;
        const { updateStoryFeed, closeModal } = this.props;
        const storyManager = StoryManager.getInstance();

        try {
            this.setState({ isLoading: true });

            const { data: { data: { stories } } } = await storyServices.createStory(uploadImages);

            stories.forEach((story) => storyManager.prependUserStory(story));
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