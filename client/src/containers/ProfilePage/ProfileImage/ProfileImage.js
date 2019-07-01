import React from 'react';

import "./ProfileImage.css";
import defaultAvatar from "../../../images/x.png";

function ProfileImage() {
    return (
        <div className="ImagesContainer__Avatar" style={{ backgroundImage: `url(${defaultAvatar})` }}>

        </div>
    )
}

export default ProfileImage;