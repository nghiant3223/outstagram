import React, { Component } from 'react'
import PropTypes from 'prop-types';
import { Popup } from 'semantic-ui-react';

import ContainerContext from '../../ContainerContext';
import * as roomServices from "../../../../../services/room.service";

import Socket from '../../../../../Socket';

import "./Message.css";

class Message extends Component {
    state = {
        error: false
    }

    static contextType = ContainerContext;

    async componentDidMount() {
        const { isNew, type, content, id } = this.props;

        if (isNew) {
            try {
                const { roomIdOrUsername, replaceMessage, roomID } = this.context;
                const { data: { data: { message } } } = await roomServices.createMessage(roomIdOrUsername, content, type);
                replaceMessage(id, message);
                Socket.emit("ROOM.CLIENT.SEND_MESSAGE", { ...message, targetRoomID: roomID });
            } catch (e) {
                this.setState({ error: true });
                console.log("Cannot create message", e);
            }
        }
    }

    render() {
        const { createdAt, content } = this.props;
        const { error } = this.state;

        if (error) {
            var className = "Message Message--Error";
        } else {
            var className = "Message"
        }

        return (
            <div className={className}>
                <Popup content={createdAt.toString()} size="mini" position="right center" style={{ padding: "0.75em" }} inverted
                    trigger={<div className="Message__Content">{content}</div>} />
            </div>
        )
    }
}

Message.propTypes = {
    content: PropTypes.string,
    createdAt: PropTypes.instanceOf(Date)
}

export default Message