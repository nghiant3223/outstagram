import React, { Component } from 'react';
import ContactSearch from './Search/ContactSearch';

import "./ContactContainer.css";
import Contact from './Contact/Contact';
import * as roomServices from "../../services/room.service";
import Socket from '../../Socket';
import ContactContainerPlaceHolder from './Placeholder';

class ContactContainer extends Component {
    state = {
        rooms: [],
        isLoading: false
    }

    contactRefs = {};

    componentDidMount() {
        Socket.on("ROOM.SERVER.SEND_MESSAGE", this.onMessageReceive);
        this.fetchMessages();
    }

    componentWillUnmount() {
        Socket.removeListener("ROOM.SERVER.SEND_MESSAGE", this.onMessageReceive);
    }

    onMessageReceive = (message) => {
        const { targetRoomID } = message.data;
        this.hoistContact(targetRoomID);
    }

    // Update last message of suitable contact when user submits message input
    updateContact = (roomID, message) => {
        this.contactRefs[roomID].update(message);
        this.hoistContact(roomID);
    }

    // Hoist contact to the top of the ContactConatainer
    hoistContact = (roomID) => {
        this.setState((prevState) => {
            let candidateRoom;
            const restRooms = prevState.rooms.filter(room => {
                const match = room.id == roomID;
                if (match) candidateRoom = room;
                return !match;
            });
            return { rooms: [candidateRoom, ...restRooms] };
        });
    }

    async fetchMessages() {
        this.setState({ isLoading: true });
        try {
            const { data: { data: { rooms } } } = await roomServices.getRecentRooms();
            this.setState({ rooms: rooms || [] });
        } catch (e) {
            console.warn("Error while fetching rooms", e);
        } finally {
            this.setState({ isLoading: false });
        }
    }

    render() {
        const { rooms, isLoading } = this.state;

        if (isLoading) return <ContactContainerPlaceHolder />;

        return (
            <div className="ContactContainer">
                <div className="ContactContainer__SearchContainer">
                    <ContactSearch onContactClick={this.onContactClick} />
                </div>
                <div className="ContactContainer__ContactItemContainer">
                    {rooms.map((contact) => <Contact key={contact.id} ref={(el) => this.contactRefs[contact.id] = el} {...contact} />)}
                </div>
            </div>
        )
    }
}

export default ContactContainer;