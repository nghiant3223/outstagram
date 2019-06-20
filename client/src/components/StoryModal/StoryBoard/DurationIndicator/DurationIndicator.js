import React, { Component } from 'react';

import './DurationIndicator.css';

class DurationIndicator extends Component {
    state = {
        percent: 0
    }

    componentDidUpdate(prevProps) {
        if (this.props.active == true && prevProps.active == false) {
            this.setState({ percent: 100 });
        }
    }

    render() {
        const { percent } = this.state;
        const { duration } = this.props;

        return (
            <div className="DurationIndicator">
                <div className="DurationIndicator__Filler" style={{ transition: `width ${duration}s linear`, width: `${percent}%` }} />
            </div>
        )
    }
}

export default DurationIndicator;