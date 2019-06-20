import React, { Component } from 'react';
import { Dropdown, Button } from 'semantic-ui-react';

import './StoryBoard.css';
import Avatar from '../../Avatar/Avatar';
import DurationIndicator from './DurationIndicator/DurationIndicator';

class StoryBoard extends Component {
    state = {
        activeStoryIndex: 0
    }

    componentDidMount() {
        const { stories, onNextStoryBoard } = this.props;
        const { activeStoryIndex } = this.state;

        this.storyTimeout = setTimeout(() => {
            this.nextStory();
        }, stories[activeStoryIndex].duration * 1000)
    }

    componentDidUpdate(_, prevState) {
        const { stories } = this.props;
        const { activeStoryIndex } = this.state;

        if (activeStoryIndex !== prevState.activeStoryIndex) {
            clearTimeout(this.storyTimeout);
            this.storyTimeout = setTimeout(this.nextStory, stories[activeStoryIndex].duration * 1000)
        }
    }

    componentWillUnmount() {
        clearTimeout(this.storyTimeout);
    }

    nextStory = () => {
        const { activeStoryIndex } = this.state;
        const { stories } = this.props;

        if (activeStoryIndex == stories.length - 1) {
            return;
        }

        this.setState((prevState) => ({ activeStoryIndex: prevState.activeStoryIndex + 1 }));
    }

    prevStory = () => {
        const { activeStoryIndex } = this.state;

        if (activeStoryIndex == 0) {
            return;
        }

        this.setState((prevState) => ({ activeStoryIndex: prevState.activeStoryIndex - 1 }));
    }

    render() {
        const { stories } = this.props;
        const { activeStoryIndex } = this.state;

        return (
            <div className="StoryBoard" style={activeStoryIndex >= 0 ? { backgroundImage: `url(/images/${stories[activeStoryIndex].huge})` } : null} >
                <div className="StoryBoard__Progress">
                    {stories.map((story, index) => <DurationIndicator duration={story.duration} activeStoryIndex={activeStoryIndex} index={index} key={index} />)}
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
                    <Button icon='chevron left' circular color="white" onClick={this.prevStory} />
                </div>

                <div className="StoryBoard__Next">
                    <Button icon='chevron right' circular color="white" onClick={this.nextStory} />
                </div>
            </div>
        )
    }
}

export default StoryBoard;