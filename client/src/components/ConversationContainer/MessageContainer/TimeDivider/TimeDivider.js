import React from 'react';

import "./TimeDivider.css";

function TimeDivider(props) {
    const { time } = props;
    return (
        <div className="MessageContainer__TimeDivider">
            {time.toString()}
        </div>
    )
}

export default TimeDivider;