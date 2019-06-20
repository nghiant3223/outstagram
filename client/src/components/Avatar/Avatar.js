import React from 'react';

import './Avatar.css';
import defaultAvatar from '../../images/avatar.png';

export default function Avatar(props) {
    const { isActive, avatar, style } = props;

    return (
        <div className={isActive ? "Avatar Avatar--Active" : "Avatar"} style={style}>
            <img src={avatar || defaultAvatar} alt="User's avatar" />
        </div>
    )
}
