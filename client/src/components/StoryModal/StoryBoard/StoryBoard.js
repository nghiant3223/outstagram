import React, { Component } from 'react';
import { Dropdown, Button } from 'semantic-ui-react';
import { connect } from 'react-redux';

import './StoryBoard.css';
import Avatar from '../../Avatar/Avatar';
import DurationIndicator from './DurationIndicator/DurationIndicator';
import { getDiffFromPast } from '../../../utils/time';

import * as storyActions from '../../../actions/story.action';
import * as uiActions from '../../../actions/ui.action';

class StoryBoard extends Component {
    state = {
        activeStoryIndex: 0
    }

    componentDidMount() {
        const { stories } = this.props.onDisplayStoryBoard.getValue();
        const { activeStoryIndex } = this.state;

        this.storyTimeout = setTimeout(this.nextStory, stories[activeStoryIndex].duration)
    }

    componentWillReceiveProps = (nextProps) => {
        const { onDisplayStoryBoard } = this.props;

        // If story board changes, restart the count
        if (onDisplayStoryBoard != nextProps.onDisplayStoryBoard) {
            this.setState({ activeStoryIndex: 0 });
        }
    }

    componentDidUpdate(prevProps, prevState) {
        const { onDisplayStoryBoard } = this.props;
        const { activeStoryIndex } = this.state;
        const { stories } = onDisplayStoryBoard.getValue();

        // If story in a story board changes or story board changes
        if (activeStoryIndex !== prevState.activeStoryIndex || onDisplayStoryBoard !== prevProps.onDisplayStoryBoard) {
            clearTimeout(this.storyTimeout);
            this.storyTimeout = setTimeout(this.nextStory, stories[activeStoryIndex].duration)
        }
    }

    componentWillUnmount() {
        clearTimeout(this.storyTimeout);
    }

    nextStory = () => {
        const { activeStoryIndex } = this.state;
        const { setOnDisplayStoryBoardNode, onDisplayStoryBoard, closeStoryModal } = this.props;
        const { stories } = this.props.onDisplayStoryBoard.getValue();

        if (activeStoryIndex == stories.length - 1) {
            if (onDisplayStoryBoard.next == null) {
                closeStoryModal();
                return;
            }
            setOnDisplayStoryBoardNode(onDisplayStoryBoard.next);
            return;
        }

        this.setState((prevState) => ({ activeStoryIndex: prevState.activeStoryIndex + 1 }));
    }

    prevStory = () => {
        const { activeStoryIndex } = this.state;
        const { setOnDisplayStoryBoardNode, onDisplayStoryBoard, closeStoryModal } = this.props;
        if (activeStoryIndex == 0) {
            if (onDisplayStoryBoard.previous == null) {
                closeStoryModal();
                return;
            }

            setOnDisplayStoryBoardNode(onDisplayStoryBoard.previous);
            return;
        }

        this.setState((prevState) => ({ activeStoryIndex: prevState.activeStoryIndex - 1 }));
    }

    render() {
        const { stories, fullname } = this.props.onDisplayStoryBoard.getValue();
        const { activeStoryIndex } = this.state;

        console.log(this.props.onDisplayStoryBoard, activeStoryIndex);

        return (
            <div className="StoryBoard" style={activeStoryIndex >= 0 ? { backgroundImage: `url(/images/${stories[activeStoryIndex].huge})` } : null} >
                <div className="StoryBoard__Progress">
                    {stories.map((story, index) => <DurationIndicator duration={story.duration} onDisplayStoryBoard={this.props.onDisplayStoryBoard} activeStoryIndex={activeStoryIndex} index={index} key={index} />)}
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
                        <Dropdown icon="ellipsis vertical" className="icon">
                            <Dropdown.Menu>
                                <Dropdown.Item text='Report' />
                            </Dropdown.Menu>
                        </Dropdown>
                    </div>
                </div>

                <div className="StoryBoard__Prev">
                    <Button icon='chevron left' circular color="white" onClick={this.prevStory} />
                </div>

                <div className="StoryBoard__Next">
                    <Button icon='chevron right' circular color="white" onClick={this.nextStory} />
                </div>
            </div>
        )
    }
}

const mapStateToProps = ({ storyReducer: { onDisplayStoryBoard } }) => ({ onDisplayStoryBoard, storyBoard: onDisplayStoryBoard.getValue() });

const mapDispatchToProps = (dispatch) => ({
    setOnDisplayStoryBoardNode: (storyBoardNode) => dispatch(storyActions.setOnDisplayStoryBoardNode(storyBoardNode)),
    closeStoryModal: () => dispatch(uiActions.closeStoryModal())
});

export default connect(mapStateToProps, mapDispatchToProps)(StoryBoard);