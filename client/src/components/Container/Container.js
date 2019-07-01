import React from 'react';

import "./Container.css";

function Container({ children, className = "" }) {
    return (
        <div className={["Container", className].join(" ")}>
            {children}
        </div>
    )
}

export default Container;