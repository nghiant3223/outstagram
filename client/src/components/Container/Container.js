import React from 'react';

import "./Container.css";

function Container({ children, className = "", style = {}, white = true }) {
    return (
        <div className={["Container", className].join(" ")} style={{ ...style, backgroundColor: white ? "white" : "unset" }}>
            {children}
        </div>
    )
}

export default Container;