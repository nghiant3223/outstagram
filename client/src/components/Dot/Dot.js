import React from 'react';

import "./Dot.css";

export default function Dot({ style = {} }) {
    return (
        <span style={{ ...style }} className="Dot">·</span>
    )
}
