import React, { Component } from 'react';

import './DurationIndicator.css';

class DurationIndicator extends Component {
    state = {
        percentage: 0
    }

    componentDidMount() {
        const { index, activeStoryIndex } = this.props;

        if (index === activeStoryIndex) {
            this.setEmptyIndicator();
            setTimeout(this.setFullIndicator, 0);
        }
    }


    componentDidUpdate(prevProps) {
        const { index, activeStoryIndex, onDisplayStoryBoard } = this.props;
        const { activeStoryIndex: prevActiveStoryIndex } = prevProps;

        if (activeStoryIndex === prevActiveStoryIndex && onDisplayStoryBoard === prevProps.onDisplayStoryBoard) {
            return;
        }

        if (index == activeStoryIndex) {
            this.setEmptyIndicator();
            setTimeout(this.setFullIndicator, 0);
        } else if (index < activeStoryIndex) {
            this.setFullIndicator();
        } else {
            this.setEmptyIndicator();
        }
    }

    setFullIndicator = () => {
        this.setState({ percentage: 100 });
    }

    setEmptyIndicator = () => {
        this.setState({ percentage: 0 });
    }

    render() {
        const { percentage } = this.state;
        const { duration, index, activeStoryIndex } = this.props;

        let style;
        if (index === activeStoryIndex) {
            if (percentage == 0) {
                style = { width: `0%` }
            } else if (percentage == 100) {
                style = { transition: `width ${duration}ms linear`, width: `100%` }
            }
        } else {
            style = { width: `${percentage}%` }
        }

        return (
            <div className="DurationIndicator">
                <div className="DurationIndicator__Filler" style={style} />
            </div>
        )
    }
}

export default DurationIndicator;