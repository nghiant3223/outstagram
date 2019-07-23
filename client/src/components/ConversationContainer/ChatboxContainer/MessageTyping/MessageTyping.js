import React, { useState } from 'react';

import { Popup } from 'semantic-ui-react';
import Avatar from "../../../Avatar/Avatar";
import Typing from '../../../Typing/Typing';

export default function MessageTyping(props) {
    const { userID, fullname } = props;
    const [isTyping, setIsTyping] = useState(true);

    if (!isTyping) return null;

    return (
        <div className={["MessageGroup", "MessageGroup--Left"].join(" ")}>
            <div className="MessageGroup__Avatar"><Avatar userID={userID} /> </div>

            <div className={["MessageGroup__ChatboxContainer", "MessageGroup__ChatboxContainer--Left"].join(" ")}>
                <div className="Message">
                    <Popup content={`${fullname} is typing`} size="mini" position="right center" inverted
                        trigger={<div className="Message__Content"><Typing /></div>} style={{ padding: "0.75em" }} />
                </div>
            </div>
        </div>
    );
}
