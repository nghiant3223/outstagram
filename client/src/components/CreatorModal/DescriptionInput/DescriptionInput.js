import React from 'react';

import "./DescriptionInput.css"

function DescriptionInput({ value, onChange }) {
    return (
        <textarea className="CaptionArea" value={value} onChange={onChange} placeholder="Add caption to your post, ignore if you dont like to" rows={3} ></textarea>
    )
}

export default DescriptionInput;