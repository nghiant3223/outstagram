import React from 'react';
import PropTypes from 'prop-types';

import "./Container.css";

function Container({ children, className, style, white, center, shadow, size }) {
    const newStyle = {}

    if (white) {
        newStyle.backgroundColor = "white";
    }

    if (center) {
        newStyle.marginLeft = "auto";
        newStyle.marginRight = "auto";
    }

    if (shadow) {
        newStyle.boxShadow = "0 1px 5px rgba(0, 0, 0, 0.15)";
    }

    switch (size) {
        case "standard":
            newStyle.width = "var(--post-width)";
            break;

        case "small":
            newStyle.width = "calc(0.75 * var(--post-width)"

        default:
            break;
    }

    return (
        <div className={["Container", className].join(" ")} style={{ ...style, ...newStyle }}>
            {children}
        </div>
    )
}

Container.propTypes = {
    size: PropTypes.oneOf(["standard", "small", "unset"])
}

Container.defaultProps = {
    className: "",
    style: {},
    white: true,
    center: true,
    shadow: false,
    size: "unset"
}

export default Container;