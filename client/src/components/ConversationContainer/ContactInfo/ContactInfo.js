import React from 'react';
import { Icon } from 'semantic-ui-react';

function ContactInfo(props) {
    const { header, partner } = props;

    return (
        <div className="MessageInfoContainer__Info">
            <div className="MessageInfoContainer__Info__Detail">
                <div className="Fullname">{header}</div>
                <div style={{ display: "flex", alignItems: "center" }}>
                    <Icon name="circle" size="mini" color="green" />
                    <span>Online</span></div>
            </div>
        </div>
    )
}

export default ContactInfo;