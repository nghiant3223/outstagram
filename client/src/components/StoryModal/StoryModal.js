import React, { Component } from 'react';
import { Modal } from 'semantic-ui-react';
import { connect } from 'react-redux';

import StoryFeedManager from '../../StoryFeedManager';

import * as uiActions from '../../actions/ui.action';

import StoryBoard from './StoryBoard/StoryBoard';

class StoryModal extends Component {
    visitedSBNodes = new Set();

    componentDidUpdate(prevProps) {
        const { isModalOpen } = this.props;

        // If modal has just closed
        if (!isModalOpen && prevProps.isModalOpen) {
            const { updateStoryFeed } = this.props;
            const storyFeedManager = StoryFeedManager.getInstance();

            // Make SBNode inactive in storyFeedManager
            this.visitedSBNodes.forEach((sbNode) => storyFeedManager.inactiveSB(sbNode));
            // Update StoryFeed in UI
            updateStoryFeed();
            // Clear visited SBNode
            this.clearVisitedSBNodes();
        }
    }

    addToVisitedSBNodes = (sbNode) => {
        if (!this.visitedSBNodes.has(sbNode)) {
            this.visitedSBNodes.add(sbNode);
        }
    }

    clearVisitedSBNodes() {
        this.visitedSBNodes.clear();
    }

    render() {
        const { closeModal, isModalOpen } = this.props;

        return (
            <Modal
                closeOnEscape
                centered={false}
                closeOnDimmerClick
                open={isModalOpen}
                onClose={closeModal}>
                <StoryBoard addToVisitedSBNodes={this.addToVisitedSBNodes} />
            </Modal>
        );
    }
}

const mapStateToProps = ({ storyReducer: { isModalOpen } }) => ({ isModalOpen });

const mapDispatchToProps = (dispatch) => ({
    closeModal: () => dispatch(uiActions.closeStoryModal())
});

export default connect(mapStateToProps, mapDispatchToProps)(StoryModal);