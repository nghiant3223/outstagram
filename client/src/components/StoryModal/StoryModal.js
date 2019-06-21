import React, { Component } from 'react';
import { Button, Header, Image, Modal } from 'semantic-ui-react';
import { connect } from 'react-redux';

import * as storyService from '../../services/story.service';

import * as uiActions from '../../actions/ui.action';
import StoryBoard from './StoryBoard/StoryBoard';

class StoryModal extends Component {
    state = {
        storyBoards: []
    }

    componentDidMount = async () => {
        try {
            const { data: { data: { storyBoards } } } = await storyService.getStoryFeed();
            if (storyBoards !== null) {
                this.setState({ storyBoards: storyBoards });
            }
        } catch (e) {
            console.log(e);
        }
    }


    render() {
        const { closeModal, isModalOpen } = this.props;

        return (
            <Modal
                closeOnEscape
                centered={false}
                closeOnDimmerClick
                open={isModalOpen}
                onClose={closeModal} style={{padding: 0}}>
                    {this.state.storyBoards.map((board) => <StoryBoard {...board} />)}
            </Modal>
        );
    }
}

const mapStateToProps = ({ story: { isModalOpen } }) => ({ isModalOpen });

const mapDispatchToProps = (dispatch) => ({
    closeModal: () => dispatch(uiActions.closeStoryModal())
});

export default connect(mapStateToProps, mapDispatchToProps)(StoryModal);