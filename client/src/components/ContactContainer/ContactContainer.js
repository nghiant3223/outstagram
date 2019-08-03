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
        Socket.on("ROOM.SERVER.CREATE_ROOM", this.onNewRoomCreated)
        this.fetchMessages();
    }

    componentWillUnmount() {
        Socket.removeListener("ROOM.SERVER.SEND_MESSAGE", this.onMessageReceive);
        Socket.removeListener("ROOM.SERVER.CREATE_ROOM", this.onNewRoomCreated)
    }

    onMessageReceive = (message) => {
        const { targetRoomID } = message.data;
        this.hoistContact(targetRoomID);
    }

    onNewRoomCreated = (message) => {
        const { newRoom } = message.data;
        this.setState((prevState) => ({ rooms: [newRoom, ...prevState.rooms] }));
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
                const match = room.id === roomID;
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

    addFakeRoom = (room) => {
        // MAGIC: I dont know why it needs to setTimeout here :?, if I dont do this, fakeRoom sometimes does not show up
        setTimeout(() => this.setState((prevState) => ({ rooms: [room, ...prevState.rooms] })), 100);
    }

    replaceFakeRoom = (roomID, newRoom) => {
        this.setState((prevState) => {
            const { rooms } = prevState;
            const restRooms = rooms.filter(room => room.id !== roomID);

            if (restRooms.length === rooms.length) {
                throw new Error("Message does not exist");
            }

            return { rooms: [newRoom, ...restRooms] };
        });
    }

    render() {
        const { rooms, isLoading } = this.state;

        if (isLoading) return <ContactContainerPlaceHolder />;

        return (
            <div className="ContactContainer" >
                <div className="ContactContainer__SearchContainer">
                    <ContactSearch users={rooms.map(room => room.partner)} />
                </div>
                <div className="ContactContainer__ContactItemContainer">
                    {rooms.map((contact) => <Contact key={contact.id} ref={(el) => this.contactRefs[contact.id] = el} {...contact} />)}
                </div>
            </div>
        )
    }
}

export default ContactContainer;