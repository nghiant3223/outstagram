import React from 'react';

import "./ClickableText.css";

function ClickableText(props) {
    const { children, onClick, fontSize } = props;

    return (
        <span className="ClickableText" style={{ fontSize }} onClick={onClick}>{children}</span>
    )
}

export default ClickableText;