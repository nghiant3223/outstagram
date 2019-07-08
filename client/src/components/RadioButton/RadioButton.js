import React from 'react';

import "./RadioButton.css";

function RadioButton({ className = [], name, value, checked, onChange, children, defaultChecked }) {
    return (
        <label className={["radio", ...className].join(" ")}>
            <input type="radio" name={name} value={value} checked={checked} onChange={onChange} defaultChecked={defaultChecked} />
            <span>{children}</span>
        </label>
    )
}

export default RadioButton;