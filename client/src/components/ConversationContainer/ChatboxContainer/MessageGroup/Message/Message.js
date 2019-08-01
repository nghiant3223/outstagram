import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { Popup } from 'semantic-ui-react';
import moment from 'moment';

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
                const { roomIdOrUsername, replaceMessage, roomID, isNewRoom, partner, onNewRoomCreated, user } = this.context;

                if (!isNewRoom) {
                    const { data: { data: { message } } } = await roomServices.createMessage(roomIdOrUsername, content, type);
                    replaceMessage(id, message);
                    Socket.emit("ROOM.CLIENT.SEND_MESSAGE", { ...message, targetRoomID: roomID });
                    return;
                }

                const { data: { data: { room } } } = await roomServices.createRoom([partner.id], { content: content, type: 0 });
                onNewRoomCreated(roomID, room);
                
                const newRoom = JSON.parse(JSON.stringify(room));
                newRoom.partner = user;
                Socket.emit("ROOM.CLIENT.CREATE_ROOM", { targetRoomID: room.id, anotherMemberID: partner.id, newRoom });
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
                <Popup content={moment(new Date(createdAt)).calendar().replace(/\sat\s/, ' ').replace('Today', '')} size="mini" position="right center" style={{ padding: "0.75em" }} inverted
                    trigger={<div className="Message__Content">{content}</div>} />
            </div>
        )
    }
}

Message.propTypes = {
    content: PropTypes.string,
    createdAt: PropTypes.instanceOf(Date)
}

export default Message;