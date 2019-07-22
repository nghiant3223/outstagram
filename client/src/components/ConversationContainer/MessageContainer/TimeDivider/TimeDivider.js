import React from 'react';

import "./TimeDivider.css";

function TimeDivider(props) {
    const { time } = props;
    return (
        <div className="MessageContainer__TimeDivider">
            {time}
        </div>
    )
}

export default TimeDivider;