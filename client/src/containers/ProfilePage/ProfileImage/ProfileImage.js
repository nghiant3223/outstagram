import React from 'react';

import "./ProfileImage.css";
import Avatar from '../../../components/Avatar/Avatar';

function ProfileImage({ userID }) {
    return (
        <div className="ImagesContainer__Avatar">
            <Avatar userID={userID} size="big" width="125px"/>
        </div>
    )
}

export default ProfileImage;