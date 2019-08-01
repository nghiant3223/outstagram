import React from 'react';
import PropTypes from 'prop-types';

import "./Container.css";

function Container({ children, className, style, white, center }) {
    const newStyle = {}

    if (white) {
        newStyle.backgroundColor = "white";
    }

    if (center) {
        newStyle.marginLeft = "auto";
        newStyle.marginRight = "auto";
    }

    return (
        <div className={["Container", className].join(" ")} style={{ ...style, ...newStyle }}>
            {children}
        </div>
    )
}

Container.defaultProps = {
    className: "",
    style: {},
    white: true,
    center: true
}

export default Container;