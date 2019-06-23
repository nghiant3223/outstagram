import React, { Component } from 'react';

import './TimeSlicer.css';

class TimeSlicer extends Component {
    state = {
        isFull: false
    }

    componentDidMount() {
        const { index, activeStoryIndex } = this.props;

        if (index === activeStoryIndex) {
            this.setEmptyIndicator();
            setTimeout(this.setFullIndicator, 0);
        }
    }

    componentDidUpdate(prevProps) {
        const { index, activeStoryIndex, sbNode } = this.props;

        if (activeStoryIndex === prevProps.activeStoryIndex
            && sbNode === prevProps.sbNode) {
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
        this.setState({ isFull: true });
    }

    setEmptyIndicator = () => {
        this.setState({ isFull: false });
    }

    render() {
        const { isFull } = this.state;
        const { duration, index, activeStoryIndex } = this.props;

        let style;
        if (index === activeStoryIndex) {
            if (!isFull) {
                style = { width: `0%` }
            } else {
                style = { transition: `width ${duration}ms linear`, width: `100%` }
            }
        } else {
            style = { width: `${isFull ? 100 : 0}%` }
        }

        return (
            <div className="DurationIndicator">
                <div className="DurationIndicator__Filler" style={style} />
            </div>
        )
    }
}

export default TimeSlicer;