import React, { useState } from 'react';
import { Button } from 'semantic-ui-react';

import Avatar from '../../../components/Avatar/Avatar';

import "./ProfileImage.css";
import { updateUser } from '../../../services/user.service';

function ProfileImage({ userID }) {
    const [isLoading, setIsLoading] = useState(false);

    let avatarInput;

    const onButtonClick = () => {
        avatarInput.click();
    }

    const onFileSelect = async (e) => {
        e.persist();
        const avatar = e.target.files[0];

        setIsLoading(true);
        try {
            await updateUser({ avatar });
            window.location.reload(false);
        } catch (e) {
            alert("Cannot update user");
        } finally {
            setIsLoading(false);
        }
    }

    return (
        <div className="ImagesContainer__Avatar">
            <Avatar userID={userID} size="big" width="125px" />
            <div className="ImagesContainer__Avatar__ChangeBtn">
                <Button icon="photo" circular onClick={onButtonClick} loading={isLoading} />
                <input type="file" accept="image/*" ref={el => avatarInput = el} onChange={onFileSelect} />
            </div>
        </div>
    )
}

export default ProfileImage;