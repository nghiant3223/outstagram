import React, { Component } from 'react';
import { Button, Header, Image, Modal } from 'semantic-ui-react';
import { connect } from 'react-redux';

import * as storyService from '../../services/story.service';

import * as uiActions from '../../actions/ui.action';
import StoryBoard from './StoryBoard/StoryBoard';

class StoryModal extends Component {
    render() {
        const { closeModal, isModalOpen, storyBoards } = this.props;

        return (
            <Modal
                closeOnEscape
                centered={false}
                closeOnDimmerClick
                open={isModalOpen}
                onClose={closeModal}>
                <StoryBoard {...storyBoards[0]} />
            </Modal>
        );
    }
}

const mapStateToProps = ({ storyReducer: { isModalOpen, storyBoards } }) => ({ isModalOpen, storyBoards });

const mapDispatchToProps = (dispatch) => ({
    closeModal: () => dispatch(uiActions.closeStoryModal())
});

export default connect(mapStateToProps, mapDispatchToProps)(StoryModal);