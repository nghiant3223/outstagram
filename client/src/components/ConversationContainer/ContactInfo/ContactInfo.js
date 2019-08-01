import React from 'react';
import OnlineStatus from '../../OnlinrStatus/OnlineStatus';

import "./ContactInfo.css";

function ContactInfo(props) {
    const { header, partner } = props;

    return (
        <div className="MessageInfoContainer__Info">
            <div className="MessageInfoContainer__Info__Detail">
                <div className="Fullname">{header}</div>
                <div className="Status">
                    <OnlineStatus partnerID={partner.id} lastLogin={partner.lastLogin} lastLogout={partner.lastLogout} />
                </div>
            </div>
        </div>
    )
}

export default ContactInfo;