import React, { Component } from 'react';
import { withRouter } from 'react-router';

import ContactInfo from './ContactInfo/ContactInfo';
import ChatboxContainer from './ChatboxContainer/ChatboxContainer';

import * as roomServices from "../../services/room.service";

class ConversationContainer extends Component {
    state = {
        room: null,
        isLoading: false
    }

    componentDidMount() {
        this.fetchRoom();
    }

    componentDidUpdate(prevProps) {
        if (this.props.match.params.roomIdOrUsername !== prevProps.match.params.roomIdOrUsername) {
            this.fetchRoom();
        }
    }

    async fetchRoom() {
        const { roomIdOrUsername } = this.props.match.params;

        if (roomIdOrUsername) {
            try {
                const { data: { data: room } } = await roomServices.getRoom(roomIdOrUsername);
                this.setState({ room });
            } catch (e) {
                console.log("Cannot fetch room", e);
            } finally {
                this.setState({ isLoading: false });
            }
        }
    }

    render() {
        const { room, isLoading } = this.state

        if (!room || isLoading) {
            return null;
        }

        const { updateLastMessage } = this.props;
        const { roomIdOrUsername } = this.props.match.params;
        const { partner, members, messages, id, type: isGroupChat } = room;

        return (
            <div className="MessageInfoContainer">
                <ContactInfo isGroupChat={isGroupChat} partner={partner} members={members} />
                <ChatboxContainer messages={messages} roomID={id} roomIdOrUsername={roomIdOrUsername} updateLastMessage={updateLastMessage} />
            </div>
        )
    }
}

export default withRouter(ConversationContainer);