import React from 'react'

import MessageGroup from '../components/ConversationContainer/MessageContainer/MessageGroup/MessageGroup';
import TimeDivider from '../components/ConversationContainer/MessageContainer/TimeDivider/TimeDivider';

import { MESSAGE_TIME_DISTANCE } from '../configs/story.config';

export function renderMessages(messages, currentUserID) {
    let container = [];
    if (messages.length < 1) {
        return container;
    }

    let blockMessages = [];
    blockMessages.push(messages[0]);
    container.push(<TimeDivider key={"time_divider_" + 0} time={messages[0].createdAt} />);

    for (let i = 1; i < messages.length; i++) {
        if (messages[i].authorID !== messages[i - 1].authorID) {
            container.push(<MessageGroup key={"message_group_" + blockMessages[0].id} messages={blockMessages} right={currentUserID == blockMessages[0].authorID} />);
            if (messages[i].createdAt - messages[i - 1].createdAt >= MESSAGE_TIME_DISTANCE) {
                container.push(<TimeDivider key={"time_divider_" + i} time={messages[i].createdAt} />);
            }
            blockMessages = [];
        } else if (messages[i].createdAt - messages[i - 1].createdAt >= MESSAGE_TIME_DISTANCE) {
            container.push(<MessageGroup key={"message_group_" + blockMessages[0].id} messages={blockMessages} right={currentUserID == blockMessages[0].authorID} />);
            container.push(<TimeDivider key={"time_divider_" + i} time={messages[i].createdAt} />);
            blockMessages = [];
        }
        blockMessages.push(messages[i]);
    }

    container.push(<MessageGroup key={"message_group_" + blockMessages[0].id} messages={blockMessages} right={currentUserID == blockMessages[0].authorID} />);
    return container;
}