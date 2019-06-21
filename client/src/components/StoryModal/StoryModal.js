import React, { Component } from 'react';
import { Button, Header, Image, Modal } from 'semantic-ui-react';
import { connect } from 'react-redux';

import * as storyService from '../../services/story.service';

import * as uiActions from '../../actions/ui.action';
import * as storyActions from '../../actions/ui.action';
import StoryBoard from './StoryBoard/StoryBoard';

class StoryModal extends Component {
    render() {
        const { closeModal, isModalOpen } = this.props;

        return (
            <Modal
                closeOnEscape
                centered={false}
                closeOnDimmerClick
                open={isModalOpen}
                onClose={closeModal}>
                <StoryBoard/>
            </Modal>
        );
    }
}

const mapStateToProps = ({ storyReducer: { isModalOpen } }) => ({ isModalOpen });

const mapDispatchToProps = (dispatch) => ({
    closeModal: () => dispatch(uiActions.closeStoryModal())
});

export default connect(mapStateToProps, mapDispatchToProps)(StoryModal);