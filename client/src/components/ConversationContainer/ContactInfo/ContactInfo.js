import React from 'react';
import { Icon } from 'semantic-ui-react';
import { groupChatName } from '../../../utils/api';

function ContactInfo(props) {
    const { isGroupChat, partner, members } = props;

    return (
        <div className="MessageInfoContainer__Info">
            <div className="MessageInfoContainer__Info__Detail">
                <div className="Fullname">{isGroupChat ? groupChatName(members) : partner.fullname}</div>
                <div style={{ display: "flex", alignItems: "center" }}>
                    <Icon name="circle" size="mini" color="green" />
                    <span>Online</span></div>
            </div>
        </div>
    )
}

export default ContactInfo;