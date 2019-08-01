import React, { useState } from 'react';
import { Button, Icon } from 'semantic-ui-react';

import "./CoverImage.css";
import { updateUser } from '../../../services/user.service';
import { noAuthStatic } from '../../../axios';

function CoverImage(props) {
    const [isLoading, setIsLoading] = useState(false);

    let coverInput;

    const onButtonClick = () => {
        coverInput.click();
    }

    const onFileSelect = async (e) => {
        e.persist();
        const cover = e.target.files[0];

        setIsLoading(true);
        try {
            await updateUser({ cover });
            window.location.reload(false);
        } catch (e) {
            alert("Cannot update user");
        } finally {
            setIsLoading(false);
        }
    }


    const { coverImageID } = props;
    return (
        <div className="ImagesContainer__Cover" style={{ backgroundImage: `url(${noAuthStatic('/images/others/' + coverImageID, { size: "big" })})` }} >
            <div className="ImagesContainer__Cover__ChangeBtn">
                <Button onClick={onButtonClick} loading={isLoading}><Icon name="photo" />Update your cover</Button>
                <input type="file" accept="image/*" ref={el => coverInput = el} onChange={onFileSelect} />
            </div>
        </div>
    )
}

export default CoverImage;