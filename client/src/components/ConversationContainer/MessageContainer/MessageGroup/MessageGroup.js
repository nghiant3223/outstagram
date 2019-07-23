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
                <Avatar />
            </div>}

            <div className={["MessageGroup__MessageContainer", right ? "MessageGroup__MessageContainer--Right" : "MessageGroup__MessageContainer--Left"].join(" ")}>
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