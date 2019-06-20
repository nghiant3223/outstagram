import React, { Component } from 'react';
import { Dropdown, Button } from 'semantic-ui-react';

import './StoryBoard.css';
import Avatar from '../../Avatar/Avatar';
import DurationIndicator from './DurationIndicator/DurationIndicator';

class StoryBoard extends Component {
    state = {
        activeStoryIndex: -1
    }

    storyTimeouts = []

    componentDidMount() {
        const { stories, onNextStoryBoard } = this.props;
        let accumulatedTime = 0;

        stories.forEach((story, index) => {
            ((immediateTimeout) => {
                const storyTimeout = setTimeout(() => {
                    this.setState({ activeStoryIndex: index });
                }, immediateTimeout);

                this.storyTimeouts.push(storyTimeout);
            })(accumulatedTime * 1000);

            accumulatedTime += story.duration;
        });

        // setTimeout(() => {
        //     onNextStoryBoard();
        // }, accumulatedTime * 1000);
    }

    componentWillUnmount() {
        this.storyTimeouts.forEach((timeout) => clearTimeout(timeout));
    }


    render() {
        const { stories } = this.props;
        const { activeStoryIndex } = this.state;

        return (
            <div className="StoryBoard" style={activeStoryIndex >= 0 ? { backgroundImage: `url(/images/${stories[activeStoryIndex].huge})` } : null} >
                <div className="StoryBoard__Progress">
                    {stories.map((story, index) => <DurationIndicator duration={story.duration} active={index == activeStoryIndex} />)}
                </div>
                <div className="StoryBoard__Header" >
                    <div className="StoryBoard__Header__Left">
                        <div>
                            <Avatar />
                        </div>
                        <div className="StoryBoard__Header__Left__Info">
                            <div><b>Trọng Nghĩa</b></div>
                            <div>9hr</div>
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
                    <Button icon='chevron left' circular  color="white" />

                </div>

                <div className="StoryBoard__Next">
                    <Button icon='chevron right' circular  color="white" />

                </div>

            </div>
        )
    }
}

export default StoryBoard;