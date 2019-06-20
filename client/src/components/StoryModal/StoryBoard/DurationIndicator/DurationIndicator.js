import React, { Component } from 'react';

import './DurationIndicator.css';

class DurationIndicator extends Component {
    state = {
        percentage: 0
    }

    componentDidMount() {
        const { index, activeStoryIndex } = this.props;

        if (index === activeStoryIndex) {
            this.setState({ percentage: 0 });

            setTimeout(() => {
                this.setState({ percentage: 100 });
            }, 0);       
         }
    }


    componentDidUpdate(prevProps) {
        const { index, activeStoryIndex } = this.props;

        if (activeStoryIndex === prevProps.activeStoryIndex) {
            return;
        }

        if (index == activeStoryIndex) {
            this.setState({ percentage: 0 });

            setTimeout(() => {
                this.setState({ percentage: 100 });

            }, 0);

        } else if (index < activeStoryIndex) {
            this.setState({ percentage: 100 });
        } else {
            this.setState({ percentage: 0 });
        }
    }

    render() {
        const { percentage } = this.state;
        const { duration, index, activeStoryIndex } = this.props;

        let style;
        if (index === activeStoryIndex) {
            if (percentage == 0) {
                style = { width: `${percentage}%` }
            } else if (percentage == 100) {
                style = { transition: `width ${duration}s linear`, width: `${percentage}%` }
            }
        } else {
            style = { width: `${percentage}%` }
        }

        console.log(index, activeStoryIndex, style);

        return (
            <div className="DurationIndicator">
                <div className="DurationIndicator__Filler" style={style} />
            </div>
        )
    }
}

export default DurationIndicator;