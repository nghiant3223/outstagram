import React, { Component } from 'react';

import "./MessagePage.css";

import Container from '../../components/Container/Container';
import ContactContainer from '../../components/ContactContainer/ContactContainer';
import ConversationContainer from '../../components/ConversationContainer/ConversationContainer';

class MessagePage extends Component {
    // Update ContactContainer when user submits message
    updateLastMessage = (roomID, lastMessage) => {
        this.contactContainer.updateContact(roomID, lastMessage);
    }

    addFakeRoom = (room) => {
        this.contactContainer.addFakeRoom(room);
    }

    replaceFakeRoom = (roomID, newRoom) => {
        this.contactContainer.replaceFakeRoom(roomID, newRoom);
    }

    render() {
        return (
            <Container>
                <div className="MessagePage">
                    <ContactContainer ref={el => this.contactContainer = el} />
                    <ConversationContainer updateLastMessage={this.updateLastMessage} replaceFakeRoom={this.replaceFakeRoom} addFakeRoom={this.addFakeRoom} />
                </div>
            </Container>
        )
    }
}

export default MessagePage;