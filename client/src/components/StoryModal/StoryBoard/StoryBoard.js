import React, { Component } from 'react';
import { Dropdown, Button, Icon } from 'semantic-ui-react';
import { connect } from 'react-redux';

import './StoryBoard.css';
import Avatar from '../../Avatar/Avatar';
import TimeSlicer from './TimeSlicer/TimeSlicer';
import { getDiffFromPast } from '../../../utils/time';

import * as storyActions from '../../../actions/story.action';
import * as uiActions from '../../../actions/ui.action';

import * as storyServices from '../../../services/story.service';

class StoryBoard extends Component {
    state = {
        activeStoryIndex: 0,
        reacted: false
    }

    componentDidMount() {
        const { sbNode, addToVisitedSBNodes } = this.props;
        const { stories, hasNewStory } = this.props.storyBoard;
        const { activeStoryIndex } = this.state;

        if (hasNewStory) {
            // HACK: Comment the following line for not creating view when view story
            storyServices.createStoryView(stories[activeStoryIndex].id);
            stories[activeStoryIndex].seen = true;
            addToVisitedSBNodes(sbNode);
        }

        this.storyTimeout = setTimeout(this.nextStory, stories[activeStoryIndex].duration)
        window.addEventListener("keydown", this.onArrowKeyDown);
    }

    onArrowKeyDown = (event) => {
        // If right arrow or down arrow is pressed
        if (event.keyCode === 39 || event.keyCode === 40) {
            this.nextStory();
            return;
        }

        // If left arrow or up arrow is pressed
        if (event.keyCode === 37 || event.keyCode === 38) {
            this.prevStory();
            return;
        }
    }

    componentWillReceiveProps(nextProps) {
        const { sbNode } = this.props;

        // If storyboard changes, restart the count
        if (sbNode !== nextProps.sbNode) {
            this.setState({ activeStoryIndex: 0 });
        }
    }

    componentDidUpdate(prevProps, prevState) {
        const { sbNode, addToVisitedSBNodes } = this.props;
        const { activeStoryIndex } = this.state;
        const { stories, hasNewStory } = this.props.storyBoard;

        if (hasNewStory) {
            // HACK: Comment the following line for not creating view when view story
            storyServices.createStoryView(stories[activeStoryIndex].id);
            stories[activeStoryIndex].seen = true;
            addToVisitedSBNodes(sbNode);
        }

        // If story in a storyboard changes or storyboard changes
        if (activeStoryIndex !== prevState.activeStoryIndex
            || sbNode !== prevProps.sbNode) {
            clearTimeout(this.storyTimeout);
            this.storyTimeout = setTimeout(this.nextStory, stories[activeStoryIndex].duration)
        }
    }

    componentWillUnmount() {
        clearTimeout(this.storyTimeout);
        window.removeEventListener("keydown", this.onArrowKeyDown);
    }

    nextStory = () => {
        const { activeStoryIndex } = this.state;
        const { displayStoryBoardNode, sbNode, closeStoryModal } = this.props;
        const { stories } = this.props.storyBoard;

        if (activeStoryIndex === stories.length - 1) {
            if (sbNode.getNext() == null) {
                closeStoryModal();
                return;
            }

            displayStoryBoardNode(sbNode.getNext());
            return;
        }

        this.setState((prevState) => ({ activeStoryIndex: prevState.activeStoryIndex + 1 }));
    }

    prevStory = () => {
        const { activeStoryIndex } = this.state;
        const { displayStoryBoardNode, sbNode, closeStoryModal } = this.props;

        if (activeStoryIndex === 0) {
            if (sbNode.getPrevious() == null) {
                closeStoryModal();
                return;
            }

            displayStoryBoardNode(sbNode.getPrevious());
            return;
        }

        this.setState((prevState) => ({ activeStoryIndex: prevState.activeStoryIndex - 1 }));
    }

    onHeartClick = () => {
        this.setState((prevState) => ({ reacted: !prevState.reacted }));
    }

    render() {
        const { sbNode } = this.props;
        const { activeStoryIndex, reacted } = this.state;
        const { stories, fullname } = this.props.storyBoard;

        return (
            <div className="StoryBoard" style={activeStoryIndex >= 0 ? { backgroundImage: `url(/images/${stories[activeStoryIndex].huge})` } : null} >
                <div className="StoryBoard__Progress">
                    {stories.map((story, index) =>
                        <TimeSlicer
                            key={story.id}
                            index={index}
                            duration={story.duration}
                            sbNode={sbNode}
                            activeStoryIndex={activeStoryIndex} />)}
                </div>
                <div className="StoryBoard__Header" >
                    <div className="StoryBoard__Header__Left">
                        <div>
                            <Avatar />
                        </div>
                        <div className="StoryBoard__Header__Left__Info">
                            <div><b>{fullname}</b></div>
                            <div>{getDiffFromPast(stories[activeStoryIndex].createdAt)}</div>
                        </div>

                    </div>
                    <div className="StoryBoard__Header__Right">
                        <Dropdown icon="ellipsis vertical" className="StoryBoard__Header__Right__Icon" direction="left">
                            <Dropdown.Menu>
                                <Dropdown.Item text='Report' />
                            </Dropdown.Menu>
                        </Dropdown>
                    </div>
                </div>

                <div className="StoryBoard__HeartContainer" onClick={this.onHeartClick}>
                    <Icon name={reacted ? "heart" : "heart outline"} color="red" size="large" />
                </div>

                <div className="StoryBoard__Prev">
                    <Button icon='chevron left' circular onClick={this.prevStory} />
                </div>

                <div className="StoryBoard__Next">
                    <Button icon='chevron right' circular onClick={this.nextStory} />
                </div>
            </div>
        )
    }
}

const mapStateToProps = ({ storyReducer: { onDisplaySBNode: sbNode } }) => ({ sbNode, storyBoard: sbNode.getValue() });

const mapDispatchToProps = (dispatch) => ({
    displayStoryBoardNode: (sbNode) => dispatch(storyActions.displayStoryBoardNode(sbNode)),
    closeStoryModal: () => dispatch(uiActions.closeStoryModal())
});

export default connect(mapStateToProps, mapDispatchToProps)(StoryBoard);