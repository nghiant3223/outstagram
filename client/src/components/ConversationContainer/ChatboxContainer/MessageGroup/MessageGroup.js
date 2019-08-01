import React from 'react';
import PropTypes from 'prop-types';
import Avatar from '../../../Avatar/Avatar';
import Message from './Message/Message';

import "./MessageGroup.css";

function MessageGroup(props) {
    const { messages, right } = props;

    return (
        <div className={["MessageGroup", right ? "MessageGroup--Right" : "MessageGroup--Left"].join(" ")}>
            {!right && <div className="MessageGroup__Avatar">
                <Avatar width="2.25em" userID={messages[0].authorID} />
            </div>}

            <div className={["MessageGroup__ChatboxContainer", right ? "MessageGroup__ChatboxContainer--Right" : "MessageGroup__ChatboxContainer--Left"].join(" ")}>
                {messages.map((message) => <Message key={message.id} {...message} />)}
            </div>
        </div>
    )
}

MessageGroup.propTypes = {
    right: PropTypes.bool,
    messages: PropTypes.array
}

MessageGroup.defaultProps = {
    right: false
}

export default MessageGroup;