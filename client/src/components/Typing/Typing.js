import React from 'react';
import PropTypes from 'prop-types';

import "./Typing.css";

export default function Typing(props) {
    const { inverted } = props;

    if (inverted) {
        var className = "MessageTyping MessageTyping--Inverted";
    } else {
        var className = "MessageTyping";
    }

    return (
        <div className={className}>
            <div></div>
            <div></div>
            <div></div>
        </div>
    )
}

Typing.propTypes = {
    inverted: PropTypes.bool
}

Typing.defaultProps = {
    inverted: false
}