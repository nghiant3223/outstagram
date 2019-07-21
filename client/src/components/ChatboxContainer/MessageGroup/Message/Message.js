import React from 'react';
import PropTypes from 'prop-types';
import { Popup } from 'semantic-ui-react';

import "./Message.css";

function Message(props) {
    const { createdAt, content } = props;
    return (
        <div className="Message">
            <Popup content={createdAt.toString()} size="mini" position="right center" style={{ padding: "0.75em" }} inverted
                trigger={<div className="Message__Content">{content}</div>} />
        </div>
    )
}

Message.propTypes = {
    content: PropTypes.string,
    createdAt: PropTypes.instanceOf(Date)
}

export default Message